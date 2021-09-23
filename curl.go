package main

import (
	"fmt"
	"os/exec"
)

func Curl(curl string, c chan <- RequestResult) {
	cmd := exec.Command("sh", "-c", curl)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		c <- RequestResult{
			Command:      cmd.String(),
			ResultHeader: "-",
			ResultBody:   string(output),
		}
	}
}