package week

import (
	"fmt"
	"time"
)

const (
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
	Kanen     = "π₯ε―ηγ΄γπ₯"
	Funen     = "πδΈηγ΄γπ"
	Shigen    = "β»οΈθ³ζΊγ΄γοΈοΈοΈοΈοΈοΈβ»οΈ"
)

func SelectDayBefore(day time.Weekday) string {
	dArr := []string{
		"", "", Kanen, Funen, Shigen, Kanen, "",
	}

	return dArr[day]
}

func SelectToday(day time.Weekday) string {
	tArr := []string{
		"", "", "", Kanen, Funen, Shigen, Kanen,
	}
	return tArr[day]
}

func CreateMessageForDate(d string, t string) (string, string) {
	var dMessage string = ""
	var tMessage string = ""

	if d != "" {
		dMessage = "ζζ₯γ―" + d
	}
	if t != "" {
		tMessage = "δ»ζ₯γ―" + t
	}

	return dMessage, tMessage
}

func MergeMessage(dMessage string, tMessage string) string {
	var retMessage string = ""
	if dMessage != "" {
		retMessage = dMessage
	}
	if tMessage != "" && dMessage == "" {
		retMessage = tMessage
	}

	if tMessage != "" && dMessage != "" {
		retMessage += "\n" + tMessage
	}
	fmt.Println(retMessage)
	return retMessage
}
