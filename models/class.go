package models

import (
	"fmt"
	"strings"
	"time"
)

type Class struct {
	Key                 string
	Instructor          Instructor
	Format              string
	DayOfWeek           time.Weekday
	TimeOfDay           time.Time
	PossibleInstructors []Instructor
}

func NewClass(instructor Instructor, format string, dayOfWeek time.Weekday, timeOfDay time.Time) Class {
	key := getClassKey(format, dayOfWeek, timeOfDay)

	return Class{
		Key:                 key,
		Instructor:          instructor,
		Format:              format,
		DayOfWeek:           dayOfWeek,
		TimeOfDay:           timeOfDay,
		PossibleInstructors: make([]Instructor, 0),
	}
}

func getClassKey(format string, dayOfWeek time.Weekday, timeOfDay time.Time) string {
	hour := timeOfDay.Hour()
	minute := timeOfDay.Minute()
	return fmt.Sprintf("%d-%02d%02d-%s",
		dayOfWeek,
		hour,
		minute,
		strings.ToLower(strings.TrimSpace(format)),
	)
}
