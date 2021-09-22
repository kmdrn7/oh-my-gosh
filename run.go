package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	Curlfile, err := os.ReadFile("Curlfile")
	if err != nil {
		fmt.Println("Error reading Curlfile:", err)
	}

	curls := strings.Split(string(Curlfile), "\n")

	for _, curl := range(curls) {
		cmd := exec.Command("sh", "-c", curl)
		fmt.Println(cmd)
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(string(output))
		fmt.Println()
	}
}
