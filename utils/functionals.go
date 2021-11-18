package utils

import (
	"strconv"
	"time"
)

func monthToInt(month string) int {
	switch month {
	case "January":
		return 1
	case "February":
		return 2
	case "March":
		return 3
	case "April":
		return 4
	case "May":
		return 5
	case "June":
		return 6
	case "July":
		return 7
	case "August":
		return 8
	case "September":
		return 9
	case "October":
		return 10
	case "November":
		return 11
	case "December":
		return 12
	default:
		panic("Unrecognized month")
	}
}

func AgeFromDateOfBirth(dob string) (int, int) {
	layOut := "2006-01-02" //yyyy/mm/dd
	dobTime, err := time.Parse(layOut, dob)

	if err != nil {
		panic(err)
	}

	var ageYear, leapAge int
	ageYear = time.Now().Year() - dobTime.Year()

	// the trick here is to combine the day and month into an integer of string type

	dobDayMonth, _ := strconv.Atoi(strconv.Itoa(dobTime.Day()) + strconv.Itoa(monthToInt(dobTime.Month().String())))
	nowDayMonth, _ := strconv.Atoi(strconv.Itoa(time.Now().Day()) + strconv.Itoa(monthToInt(time.Now().Month().String())))

	// if the DOB's day + month is larger than today's day + month
	// then the age is still younger by 1 year

	if dobDayMonth > nowDayMonth {
		ageYear = ageYear - 1
	}

	if dobDayMonth == 292 { // dob on 29th Feb - leap year
		leapAge = ageYear / 4
	} else {
		leapAge = 0
	}

	return ageYear, leapAge
}
