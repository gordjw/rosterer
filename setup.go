package main

import (
	"fmt"
	"log"
	"os"
	"rosterer/models"
	"strings"
	"time"

	"github.com/gordjw/go-momence"
)

func (env *Env) setup() {
	momence := momence.NewMomence(
		os.Getenv("MOMENCE_HOST_ID"),
		os.Getenv("MOMENCE_TOKEN"),
		"https://api.momence.com/api/v1",
	)

	var tzLocation = os.Getenv("LOCAL_TIMEZONE")

	var availability = map[string]map[time.Weekday]map[int]bool{
		"lauren": {
			1: {
				6: true,
				7: true,
				8: true,
			},
		},
		"paige": {
			1: {
				17: true,
				18: true,
			},
		},
	}

	instructors, err := momence.GetTeachers()
	if err != nil {
		fmt.Printf("Couldn't get instructors from Momence: %s", err.Error())
	}
	for _, instructor := range instructors {
		var key = strings.ToLower(strings.TrimSpace(instructor.FirstName))

		if key == "lauren" || key == "paige" {
			awt := availability[key]
			env.Instructors[key] = models.NewInstructor(instructor.FirstName, awt)
		}
	}

	// Get a list of classes from Momence
	classList, err := momence.GetEvents()
	if err != nil {
		fmt.Printf("Couldn't get classes from Momence: %s", err.Error())
	}

	loc, err := time.LoadLocation(tzLocation)
	if err != nil {
		log.Fatalf("Couldn't load time location: %s\n", err.Error())
	}

	for _, class := range classList {
		var teacherKey = strings.ToLower(strings.TrimSpace(strings.Split(class.OriginalTeacher, " ")[0]))

		// Momence returns dateTime strings in RFC3339 and UTC, so convert to local timezone
		t, err := time.Parse(time.RFC3339, class.DateTime)
		if err != nil {
			log.Fatalf("Couldn't parse date: %v", err.Error())
		}
		t = t.In(loc)

		c := models.NewClass(env.Instructors[teacherKey], class.Title, t.Weekday(), getTimeFromHoursMinutes(t.Hour(), t.Minute()))

		possibleInstructors := make([]models.Instructor, 0)
		for _, instructor := range env.Instructors {
			if instructor.CanTeach(c) {
				possibleInstructors = append(possibleInstructors, instructor)
			}
		}
		c.PossibleInstructors = possibleInstructors

		env.DesiredTimetable.AddClass(c)
	}
}

// TODO - Is this function necessary? Could it just be a Sprintf in the caller?
func getTimeFromHoursMinutes(hours, minutes int) time.Time {
	timeString := fmt.Sprintf("%02d:%02d", hours, minutes)
	t, err := time.Parse("15:04", timeString)
	if err != nil {
		fmt.Printf("Error parsing time: %s, %v\n", timeString, err.Error())
		return time.Time{}
	}
	return t
}
