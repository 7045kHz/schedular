package models

type Job struct {
	Name    string   `json:"name"`
	Engine  string   `json:"engine"`
	Exec    string   `json:"exec"`
	Verbose int      `json:"verbose"`
	Args    []string `json:"args"`
	Env     []string `json:"env"`
}
