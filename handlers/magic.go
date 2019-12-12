package handlers

import "strings"

func Magic(card string, num string) map[string]string {
	var cardinal, number string
	for _, i := range card {

		if len(string(i)) != 0 {

			cardinal = cardinal + string(i)

		}
	}
	for _, i := range num {

		if len(string(i)) != 0 {

			number = number + string(i)

		}
	}

	seatCardinal := strings.Split(cardinal, ",")
	seatNumber := strings.Split(number, ",")

	var resMap map[string]string

	resMap = make(map[string]string)

	for i := 0; i <= len(seatNumber)-1; i++ {

		resMap[seatNumber[i]] = seatCardinal[i]

	}

	return resMap
}
