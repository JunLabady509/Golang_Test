package main

import (
	"math/rand"
	"strconv"
	"time"
)

func MakeDiceRoll(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n) + 1
}

func PrintDice(result int) string {
	if result < 100 && result > 9 {
		str := "00"
		result2 := str + strconv.Itoa(result)
		return result2
	} else if result < 10 {
		str := "000"
		result2 := str + strconv.Itoa(result)
		return result2
	} else if result > 99 && result < 1000 {
		str := "0"
		result2 := str + strconv.Itoa(result)
		return result2
	} else {
		return strconv.Itoa(result)
	}
}

func Array_Dice(n int) []int {
	array := make([]int, n)
	var dice2 = MakeDiceRoll(2)
	var dice4 = MakeDiceRoll(4)
	var dice6 = MakeDiceRoll(6)
	var dice8 = MakeDiceRoll(8)
	var dice10 = MakeDiceRoll(10)
	var dice12 = MakeDiceRoll(12)
	var dice20 = MakeDiceRoll(20)
	var dice100 = MakeDiceRoll(100)
	for i := 0; i < n; i++ {
		array = append(array, dice2)
		array = append(array, dice4)
		array = append(array, dice6)
		array = append(array, dice8)
		array = append(array, dice10)
		array = append(array, dice12)
		array = append(array, dice20)
		array = append(array, dice100)
	}
	return array
}

func PrintMultiDice(array []int) []string {
	ArrToPrint := make([]string, len(array))
	for i := 0; i < len(array); i++ {
		if array[i] < 100 && array[i] > 9 {
			str := "00"
			result2 := str + strconv.Itoa(array[i])
			ArrToPrint = append(ArrToPrint, result2)
		} else if array[i] < 10 {
			str := "000"
			result2 := str + strconv.Itoa(array[i])
			ArrToPrint = append(ArrToPrint, result2)
		} else if array[i] > 99 && array[i] < 1000 {
			str := "0"
			result2 := str + strconv.Itoa(array[i])
			ArrToPrint = append(ArrToPrint, result2)
		} else {
			ArrToPrint = append(ArrToPrint, strconv.Itoa(array[i]))
		}
	}
	return ArrToPrint
}
