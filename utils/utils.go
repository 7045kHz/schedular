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
	weekday := time.Now().Weekday()
	hour := time.Now().Hour()
	minute := time.Now().Minute()
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
