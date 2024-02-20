package main

import (
	"fmt"
	"log"
	"rosterer/models"
	"time"

	"github.com/joho/godotenv"
)

type Env struct {
	Instructors      map[string]models.Instructor
	DesiredTimetable models.Timetable
}

func main() {
	// Load .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	fmt.Println("Writing a roster...")

	env := &Env{
		Instructors:      make(map[string]models.Instructor, 0),
		DesiredTimetable: models.NewTimetable(time.April, time.June),
	}

	env.setup()

	env.DesiredTimetable.SortSchedule()
	env.DesiredTimetable.PrintSchedule()
}
