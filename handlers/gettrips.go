package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"../models"
)

// work on setting headers

func makeReq(url string, headername string, headervalue string) models.FullResp {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	// req.Header.Add(headername, "2833")
	req.Header.Add(headername, headervalue)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("the Request failed with error %s \n", err)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	var response models.FullResp
	egg := string(data)

	bytes := []byte(egg)
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		panic(err)
	}
	// buses := response.Buses
	return response
}

func tripsReq(url string, headername string, headervalue string, header2name string, header2value string) models.FullResp {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	// req.Header.Add(headername, "2833")
	req.Header.Add(headername, headervalue)
	req.Header.Add(header2name, header2value)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("the Request failed with error %s \n", err)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	var response models.FullResp
	egg := string(data)

	bytes := []byte(egg)
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		panic(err)
	}
	// buses := response.Buses
	return response
}

func getTripsData(routeid string, departuredate string) models.FullResp {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000/v3/buses", nil)
	req.Header.Add("Routeid", routeid)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("the Request failed with error %s \n", err)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	var response models.FullResp
	egg := string(data)

	bytes := []byte(egg)
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		panic(err)
	}
	buses := response.Buses
	var mainTrips []models.Respy

	for _, i := range buses {
		queueid := i["busqueueid"]
		classid := i["busclassid"]
		orderid := i["busorderid"]
		arrangementid := i["busseatarrangementid"]

		req, err := http.NewRequest("GET", "http://127.0.0.1:5000/v3/seatamount", nil)
		req.Header.Add("Busseatarrangementid", arrangementid)
		resp, err := client.Do(req)

		if err != nil {
			fmt.Printf("the Request failed with error %s \n", err)
		}
		defer resp.Body.Close()

		data, _ := ioutil.ReadAll(resp.Body)
		var response models.FullResp
		egg := string(data)

		bytes := []byte(egg)
		err = json.Unmarshal(bytes, &response)
		if err != nil {
			panic(err)
		}
		// fmt.Println(queueid, classid, orderid, arrangementid)
		numberOfSeat := response.BusSeat[0]["NumberOfSeat"]
		// fmt.Println(numberOfSeat)

		// get the seats

		seats := makeReq("http://127.0.0.1:5000/v3/seats", "Busqueueid", queueid)

		if len(seats.Seats) <= 0 {
			fmt.Println("no seats available")
		}
		bookedSeats := []string{}
		for _, z := range seats.Seats {
			if z["TripDate"][0:10] == departuredate {
				bookedSeats = append(bookedSeats, z["SeatNumber"])
			}
		}

		noOfSeat, err := strconv.Atoi(numberOfSeat)
		if err != nil {
			fmt.Println("an error occured")
		}
		freeSeats := noOfSeat - len(bookedSeats)
		// fmt.Println(freeSeats)
		trips := tripsReq("http://127.0.0.1:5000/v3/trips", "Routeid", routeid, "Busclassid", classid)
		// fmt.Printf("%+v", trips)
		price := trips.Prices[0]["Prizing"]

		tripT := models.Respy{
			"busQueueID":           queueid,
			"busClassID":           classid,
			"busOrderID":           orderid,
			"busSeatArrangementID": arrangementid,
			"numberOfSeat":         numberOfSeat,
			"bookedSeats":          bookedSeats,
			"freeSeats":            freeSeats,
			"price":                price,
		}
		// fmt.Println("tript", tripT)

		mainTrips = append(mainTrips, tripT)

	}

	var finRes models.FullResp
	finRes.Status = "found"
	finRes.Trips = mainTrips
	// fmt.Printf("%+v", finRes)
	return finRes
}

func GetTrips(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	routeId := r.Header["Routeid"][0]
	departuredate := r.Header["Departuredate"][0]

	data := getTripsData(routeId, departuredate)

	json.NewEncoder(w).Encode(data)

	Logger("/v3/gettrips")
}
