package models

type Instructor struct {
	Name string
}

func NewInstructor(name string) Instructor {
	return Instructor{
		Name: name,
	}
}
