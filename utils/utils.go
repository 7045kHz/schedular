package utils

import (
	"fmt"
	"log"
	"os/exec"
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
