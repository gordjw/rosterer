package main

import (
	"fmt"
	"rosterer/models"
)

type Env struct {
	Instructors      []models.Instructor
	DesiredTimetable models.Timetable
	ActualTimetable  models.Timetable
}

func main() {
	fmt.Println("Writing a roster...")

	env := &Env{
		Instructors:      make([]models.Instructor, 0),
		DesiredTimetable: models.Timetable{},
		ActualTimetable:  models.Timetable{},
	}

	env.setup()

	fmt.Printf("env: %v", env)
}
