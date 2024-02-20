package models

import (
	"time"
)

type Instructor struct {
	Name         string
	Availability map[time.Weekday]map[int]bool
}

func NewInstructor(name string, availability map[time.Weekday]map[int]bool) Instructor {
	return Instructor{
		Name:         name,
		Availability: availability,
	}
}

func (i *Instructor) CanTeach(class Class) bool {
	day := class.DayOfWeek
	hour := class.TimeOfDay.Hour()

	_, ok := i.Availability[day]
	if !ok {
		return false
	}

	_, ok = i.Availability[day][hour]
	if !ok {
		return false
	}

	return i.Availability[day][hour]
}
