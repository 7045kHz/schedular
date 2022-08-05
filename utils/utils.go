package utils

import (
	"fmt"
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
