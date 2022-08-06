package models

// Matches Table [dbo].[JOB_INVENTORY]
type Jobs struct {
	Name             string   `json:"name"`
	Job_Id           int      `json:"job_id"`
	Execution_Server string   `json:"execution_server"`
	Enabled          int      `json:"enabled"`
	Job_Definition   string   `json:"job_definition"`
	Days             []string `json:"days"`
	Hour             int      `json:"hour"`
	Minute           int      `json:"minute"`
	Created_On       string   `json:"created_on"`
	Last_Updated     string   `json:"last_updated"`
	Last_Run         string   `json:"last_run"`
}
