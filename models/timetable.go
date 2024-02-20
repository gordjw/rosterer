package models

import (
	"fmt"
	"time"
)

type Timetable struct {
	StartMonth       time.Month
	EndMonth         time.Month
	Schedule         map[string]Class
	OrderedClassKeys []string
}

func NewTimetable(startMonth, endMonth time.Month) Timetable {
	schedule := make(map[string]Class, 0)

	return Timetable{
		StartMonth: startMonth,
		EndMonth:   endMonth,
		Schedule:   schedule,
	}
}

func (t *Timetable) AddClass(class Class) {
	t.Schedule[class.Key] = class
}

func (t *Timetable) PrintSchedule() {
	fmt.Println("========================")
	fmt.Printf("Start: %v\nEnd: %v\n\n", t.StartMonth, t.EndMonth)

	for _, key := range t.OrderedClassKeys {
		class := t.Schedule[key]
		fmt.Printf("--- %v ---\n", key)

		fmt.Printf("%s (%v)\n",
			class.Format,
			class.TimeOfDay.Format("03:04 PM"),
		)
		for _, possibleInstructor := range class.PossibleInstructors {
			fmt.Printf("  - %s\n", possibleInstructor.Name)
		}

		fmt.Printf("\n")
	}
}

// Sorts ScheduleWithKey by the key value
func (t *Timetable) SortSchedule() {
	keys := make([]string, 0)
	for k := range t.Schedule {
		keys = append(keys, k)
	}

	t.OrderedClassKeys = quicksort(&keys, 0, len(keys)-1)
}

func quicksort(keys *[]string, low, high int) []string {
	if low < 0 || low >= high {
		return *keys
	}
	if len(*keys) <= 1 {
		return *keys
	}

	pivot := partition(*keys, low, high)

	quicksort(keys, low, pivot-1)
	quicksort(keys, pivot+1, high)

	return *keys
}

func partition(keys []string, low, high int) int {
	pivotValue := keys[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if keys[j] <= pivotValue {
			i += 1
			t := keys[i]
			keys[i] = keys[j]
			keys[j] = t
		}
	}
	i += 1
	keys[high] = keys[i]
	keys[i] = pivotValue

	return i
}
