package handlers

import (
	"database/sql"
	"fmt"
	"log"

	"../models"
	_ "github.com/denisenkom/go-mssqldb"
)

func bookData(bookingData models.Resp, seatData models.Resp) {
	// figure out how to handle the book data this night
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	conn, err := sql.Open("mssql", connString)

	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	db := conn

	tsql := fmt.Sprintf(`
	INSERT into dbo.passengerBooking(BookingCode,Title,FirstName,Surname, FullName, Gender,DateOfBirth,PhoneNumber,
	EmailAddress,NextOfKinFullName,NextOfKinPhoneNumber,
	RouteName,BusClassName,BusOrderName,NoOfSeat,SeatNumber,SeatCardinal,
	DepartureDate, DiscountedAmount,Amount,ConvenienceFee,TotalAmount,
	BookingDate,PaymentChannel,PaymentStutus,IsTickedUsed,PassengerBookingStatus )
	VALUES('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')
	`, bookingData["bookingcode"], bookingData["title"], bookingData["firstname"], bookingData["surname"],
		bookingData["fullname"], bookingData["gender"], bookingData["dob"], bookingData["number"],
		bookingData["email"], bookingData["nextofkin"], bookingData["nextnumber"], bookingData["routename"], bookingData["classname"],
		bookingData["ordername"], bookingData["noofseat"], bookingData["seatno"], bookingData["seatcardinal"], bookingData["date"],
		bookingData["discount"], bookingData["amount"], bookingData["convenience"], bookingData["total"], bookingData["bookingdate"],
		bookingData["paymentchannel"], bookingData["paymentstatus"], bookingData["isticketused"], bookingData["status"])

	_, err = db.Exec(tsql)

	if err != nil {
		fmt.Println("error executing command")
	}

	for key, value := range seatData {
		if len(key) != 0 {
			tsql := fmt.Sprintf(`
			INSERT into dbo.BookedSeat(TripDate,SeatNumber,SeatCardinal,BlockedStatus,BusesQueueID)
            VALUES ('%s','%s','%s','%s','%s')
			`, bookingData["date"], key,
				value, "0", bookingData["queueid"])

			_, err := db.Exec(tsql)

			if err != nil {
				fmt.Println("error executing command")
			}
		}
	}

	tsql = fmt.Sprintf(`
	SELECT *
	FROM dbo.CustomerDetails
	WHERE PhoneNumber = '%s'
	`, bookingData["number"])

	rows, err := db.Query(tsql)

	if err != nil {
		fmt.Println("error querying database")
	}

	defer rows.Close()

	tsql = fmt.Sprintf(`
	INSERT into dbo.CustomerDetails(PhoneNumber,Title,FirstName,Surname
	,FullName,Gender,DateOfBirth,EmailAddress,NextOfKinFullName,
	NextOfKinPhoneNumber)
	VALUES('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')
	`, bookingData["number"], bookingData["title"], bookingData["firstname"],
		bookingData["surname"], bookingData["fullname"], bookingData["gender"],
		bookingData["dob"], bookingData["email"], bookingData["nextofkin"],
		bookingData["nextnumber"])

	_, err = db.Exec(tsql)

	if err != nil {
		fmt.Println("error executing command")
	}

}

// write data for testing BookData
