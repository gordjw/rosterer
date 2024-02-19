package main

import "rosterer/models"

func (env *Env) setup() {
	env.Instructors = append(env.Instructors, models.NewInstructor("Lauren"))
	env.Instructors = append(env.Instructors, models.NewInstructor("Lorren"))
}
