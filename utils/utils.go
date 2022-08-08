package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func GetPath(fileName string) (path string, err error) {
	path, err = exec.LookPath(fileName)
	if err != nil {
		fmt.Printf("GetPath() error: %v\n", err)
		return "", err
	}
	return path, nil
}

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Now() string {

	loc, _ := time.LoadLocation("America/New_York")
	weekday := time.Now().In(loc).Weekday()
	hour := time.Now().In(loc).Hour()
	minute := time.Now().In(loc).Minute()
	now := fmt.Sprintf("%d:%d:%d", weekday, hour, minute)
	return now
}

func GetHostname() (hostname string, err error) {
	hostname, err = os.Hostname()
	if err != nil {
		return "", err
	}

	return hostname, err
}
