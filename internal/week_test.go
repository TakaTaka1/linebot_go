package week

import (
	"testing"
	"time"
)

const (
	TestSunday    = 0
	TestMonday    = 1
	TestTuesday   = 2
	TestWednesday = 3
	TestThursday  = 4
	TestFriday    = 5
	TestSaturday  = 6
	TestKanen     = "🔥可燃ゴミ🔥"
	TestFunen     = "👀不燃ゴミ👀"
	TestShigen    = "♻️資源ゴミ️️️️️️♻️"
)

func setTestDate() []time.Weekday {

	weekDay := []time.Weekday{
		time.Weekday(TestSunday),
		time.Weekday(TestMonday),
		time.Weekday(TestTuesday),
		time.Weekday(TestWednesday),
		time.Weekday(TestThursday),
		time.Weekday(TestFriday),
		time.Weekday(TestSaturday),
	}

	return weekDay
}

func TestSelectDayBefore(t *testing.T) {
	weekDay := setTestDate()

	testStruct := []struct {
		message string
		day     string
	}{
		{
			message: "",
			day:     "Sunday",
		},
		{
			message: "",
			day:     "Monday",
		},
		{
			message: TestKanen,
			day:     "Tueseday",
		},
		{
			message: TestFunen,
			day:     "Wednesday",
		},
		{
			message: TestShigen,
			day:     "Thursday",
		},
		{
			message: TestKanen,
			day:     "Friday",
		},
		{
			message: "",
			day:     "Saturday",
		},
	}

	for i, v := range testStruct {
		t.Run("SelectDayBefore", func(t *testing.T) {
			t.Log("SelectDayBefore Test : " + v.day)
			result := SelectDayBefore(weekDay[i])
			expect := v.message

			if result != expect {
				t.Error("\nActual： ", result, "\nExpectation： ", expect)
			}
		})
	}
}

func TestSelectToday(t *testing.T) {
	testStruct := []struct {
		message string
		day     string
	}{
		{
			message: "",
			day:     "Sunday",
		},
		{
			message: "",
			day:     "Monday",
		},
		{
			message: "",
			day:     "Tueseday",
		},
		{
			message: TestKanen,
			day:     "Wednesday",
		},
		{
			message: TestFunen,
			day:     "Thursday",
		},
		{
			message: TestShigen,
			day:     "Friday",
		},
		{
			message: TestKanen,
			day:     "Saturday",
		},
	}
	// weekDay := setTestDate()
	for i, v := range testStruct {
		t.Run("SelectToday", func(t *testing.T) {
			t.Log("SelectToday test day : " + v.day)
			result := SelectToday(setTestDate()[i])
			expect := v.message
			if result != expect {
				t.Error("\nActual： ", result, "\nExpectation： ", expect)
			}
		})
	}
}

func TestCreateMessageForDate(t *testing.T) {
	testStruct := []struct {
		input1 string
		input2 string
	}{
		{
			input1: "",
			input2: "",
		},
		{
			input1: TestKanen,
			input2: "",
		},
		{
			input1: TestFunen,
			input2: TestKanen,
		},
	}
	for _, v := range testStruct {
		t.Run("CreateMessageForDate", func(t *testing.T) {
			result1, result2 := CreateMessageForDate(v.input1, v.input2)
			expect1 := "明日は" + v.input1
			expect2 := "今日は" + v.input2
			if v.input1 == "" {
				expect1 = ""
			}
			if v.input2 == "" {
				expect2 = ""
			}

			if result1 != expect1 {
				t.Error("\nActual： ", result1, "\nExpectation： ", expect1)
			}
			if result2 != expect2 {
				t.Error("\nActual： ", result2, "\nExpectation： ", expect2)
			}
		})
	}

}

func TestMergeMessage(t *testing.T) {
	t.Run("MergeMessage", func(t *testing.T) {
		t.Log("MergeMessage test")
		result := MergeMessage("", "")
		expect := ""
		if result != expect {
			t.Error("\nActual： ", result, "\nExpectation： ", expect)
		}
	})
	t.Run("MergeMessage", func(t *testing.T) {
		t.Log("MergeMessage test")
		result := MergeMessage("test", "")
		expect := "test"
		if result != expect {
			t.Error("\nActual： ", result, "\nExpectation： ", expect)
		}
	})
	t.Run("MergeMessage", func(t *testing.T) {
		t.Log("MergeMessage test")
		result := MergeMessage("test1", "test2")
		expect := "test1\ntest2"
		if result != expect {
			t.Error("\nActual： ", result, "\nExpectation： ", expect)
		}
	})
}
