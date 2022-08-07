package utils

import (
	"fmt"
	"log"
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
