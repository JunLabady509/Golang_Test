package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getTime(w http.ResponseWriter, r *http.Request) {
	var time = sayTime()
	fmt.Printf("New Request: %s", r.URL.Path)
	io.WriteString(w, time)
}

func startServer() {
	http.HandleFunc("/", getTime)
	http.HandleFunc("/dice", getDice)
	http.HandleFunc("/dices", getMultiDice)
	err := http.ListenAndServe(":4567", nil)
	fmt.Println("Server started on port 4567 ... ")
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getDice(w http.ResponseWriter, r *http.Request) {
	var dice = MakeDiceRoll(1000)
	var toPrint = PrintDice(dice)
	fmt.Printf("New Request: %s", r.URL.Path)
	io.WriteString(w, toPrint)
}

func getMultiDice(w http.ResponseWriter, r *http.Request) {
	var arr = Array_Dice(15)
	var toPrint = PrintMultiDice(arr)
	fmt.Printf("New Request: %s", r.URL.Path)
	for i := 0; i < len(toPrint); i++ {
		io.WriteString(w, toPrint[i])
	}
}
