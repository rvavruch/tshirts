package main

import (
    "flag"
    "fmt"
    "time"
)

func even(number int) bool {
    return number%2 == 0
}

func isWashingDay(today time.Time, thursday bool) bool {
    if today.Format("Mon") == "Mon" {
	    fmt.Println("Monday!")
	    return true
    } else if thursday && today.Format("Mon") == "Thu" {
	    _, week := today.ISOWeek()
	    if even(week) {
		    fmt.Println("Thursday!")
		    return true
	    }
    }
    return false
}

func main() {
    numberOftshirtsPtr := flag.Int("shirts", 9, "a int number of tshirts")
    thursdayPtr := flag.Bool("thursday", true, "a bool use Thursday")
    dryingPtr := flag.Bool("drying", true, "a bool use drying")
    flag.Parse()

    var clean, minClean = *numberOftshirtsPtr, *numberOftshirtsPtr
    var dirty, washing, drying int
    today, _ := time.Parse(time.RFC3339, "2017-02-06T00:00:00+00:00")

    fmt.Println("Date       | C | D | W | Y ")

    testDays := 365
    for i := 0; i < testDays; i++ {
	    if *dryingPtr {
		    clean += drying
		    drying = 0

		    drying = washing
		    washing = 0
	    } else {
		    clean += washing
		    washing = 0
	    }

	    if clean == 0 {
		    fmt.Println("Run out of clean shirts!")
		    break
	    }

	    if isWashingDay(today, *thursdayPtr) {
		    washing = dirty
		    dirty = 0
	    }

	    // take a clean tshirt and wear it - it immediately becomes dirty
	    clean--
	    dirty++

	    if clean < minClean {
		    minClean = clean
	    }

	    fmt.Printf("%s | %d | %d | %d | %d \n", today.Format("2006-01-02"), clean, dirty, washing, drying)

	    today = today.AddDate(0, 0, 1)
    }

    fmt.Println("Total days:", testDays)
    fmt.Println("Minimum clean shirts:", minClean)
}
