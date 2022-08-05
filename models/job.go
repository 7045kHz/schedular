package models
type Job struct {
	Engine string
	Script string
	Args []string
	Env []string
}