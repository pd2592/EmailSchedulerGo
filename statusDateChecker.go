package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type bookingData struct {
	CustomerName    string
	CustomerEmail   string
	CustomerPhone   string
	NumberOfRoom    int
	CheckInDate     time.Time
	CheckOutDate    time.Time
	HotelID         int
	BookingStatusID string
}

var varBookingData []bookingData

func checkStatus() []bookingData {

	db, err := sql.Open("mysql", "root:@/user?parseTime=true")
	checkErr(err)
	//var datetime = time.Now()
	stmt, err := db.Prepare("SELECT CustomerName, CustomerEmail, CustomerPhone, NumberOfRoom, CheckInDate, CheckOutDate, HotelID, BookingStatusId FROM `booking_detail` WHERE BookingStatusId = 'CON' and CheckInDate > ?")
	checkErr(err)
	rows, err := stmt.Query("NOW()")
	checkErr(err)
	fmt.Println(">>>>", err)
	defer db.Close()
	var bData bookingData
	for rows.Next() {
		err := rows.Scan(&bData.CustomerName, &bData.CustomerEmail, &bData.CustomerPhone, &bData.NumberOfRoom, &bData.CheckInDate, &bData.CheckOutDate, &bData.HotelID, &bData.BookingStatusID)
		fmt.Println(err)
		checkErr(err)
		fmt.Println("/////")

		fmt.Println("Booking Status is : ", bData.BookingStatusID)
		fmt.Println("CheckInDate is : ", bData.CheckInDate)
		fmt.Println("CustomerEmail is : ", bData.CustomerEmail)
		BookingData := bookingData{
			CustomerName:    bData.CustomerName,
			CustomerEmail:   bData.CustomerEmail,
			CustomerPhone:   bData.CustomerPhone,
			NumberOfRoom:    bData.NumberOfRoom,
			CheckInDate:     bData.CheckInDate,
			CheckOutDate:    bData.CheckOutDate,
			HotelID:         bData.HotelID,
			BookingStatusID: bData.BookingStatusID,
		}
		varBookingData = append(varBookingData, BookingData)

	}

	return varBookingData

}

func checkErr(err error) {

	if err != nil {
		//fmt.Println(err)
		log.Fatal(err)
		os.Exit(1)
	}
}
