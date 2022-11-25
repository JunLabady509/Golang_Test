package main

import (
	"strconv"
	"time"
)

func sayTime() string {

	DateAndTime := time.Now()
	mois := [12]string{
		"Janvier",
		"Février",
		"Mars",
		"Avril",
		"Mai",
		"Juin",
		"Juillet",
		"Août",
		"Septembre",
		"Octobre",
		"Novembre",
		"Décembre",
	}

	var DateString string

	Day := strconv.Itoa(DateAndTime.Day())
	Month := mois[int(DateAndTime.Month()-1)]
	Year := strconv.Itoa(DateAndTime.Year())
	Hour := strconv.Itoa(DateAndTime.Hour())
	Minute := strconv.Itoa(DateAndTime.Minute())
	Second := strconv.Itoa(DateAndTime.Second())

	DateString = Day + " " + Month + " " + Year + " " + Hour + ":" + Minute + ":" + Second
	return DateString
}
