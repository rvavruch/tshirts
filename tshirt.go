package main

import (
	"flag"
	"fmt"
	"time"
)

type ClothesState struct {
	day     time.Time
	clean   int
	dirty   int
	washing int
	drying  int
}

func (cs ClothesState) String() string {
	return format(cs.day.Format("2006-01-02"), cs.clean, cs.dirty, cs.washing, cs.drying)
}

func (cs ClothesState) Advance(doDrying, doThursdays bool) ClothesState {
	tomorrow := cs.day.AddDate(0, 0, 1)

	clean := cs.clean
	clean += cs.drying

	drying := cs.washing
	if !doDrying {
		// take all the drying directly in to clean
		clean += drying
		drying = 0
	}

	washing := 0
	dirty := cs.dirty
	if isWashingDay(tomorrow, doThursdays) {
		washing = dirty
		dirty = 0
	}

	clean--
	dirty++

	return ClothesState{
		day:     tomorrow,
		clean:   clean,
		dirty:   dirty,
		washing: washing,
		drying:  drying,
	}

}

func even(number int) bool {
	return number%2 == 0
}

func format(day, clean, dirty, washing, drying interface{}) string {
	return fmt.Sprintf("%-11v| %-2v| %-2v| %-2v| %-2v", day, clean, dirty, washing, drying)
}

func isWashingDay(today time.Time, doThursday bool) bool {

	if today.Weekday() == time.Monday {
		fmt.Println(today.Weekday().String())
		return true
	}
	
	if doThursday && today.Weekday() == time.Thursday {
		_, week := today.ISOWeek()
		if even(week) {
			fmt.Println(today.Weekday().String())
			return true
		}
	}
	
	return false
}

func main() {
	numberOftshirts := 9
	doThursdays := true
	doDrying := false
	
	flag.IntVar(&numberOftshirts, "shirts", numberOftshirts, "the number of tshirts")
	flag.BoolVar(&doThursdays, "thursday", doThursdays, "use Thursday")
	flag.BoolVar(&doDrying, "drying", doDrying, "use drying")
	flag.Parse()

	today, _ := time.Parse(time.RFC3339, "2017-02-06T00:00:00+00:00")
	
	// Set the state with 1 dirty shirt to match OP logic
	state := ClothesState{
		day:   today,
		clean: numberOftshirts - 1,
		dirty: 1,
	}
	
	minClean := state.clean
	
	fmt.Println(format("Date", "C", "D", "W", "Y"))
	fmt.Printf("%v\n", state)
	
	testDays := 365
	i := 0
	for ; i < testDays; i++ {
		state = state.Advance(doDrying, doThursdays)
		fmt.Printf("%v\n", state)

		if state.clean < minClean {
			minClean = state.clean
		}
		
		if state.clean < 0 {
			fmt.Println("Run out of clean shirts!")
	            	break
		}
	}
	
	fmt.Println("Total days:", i)
	fmt.Println("Minimum clean shirts:", minClean)
}

