package models

type Job struct {
	Engine string
	Exec   string
	Args   []string
	Env    []string
}
