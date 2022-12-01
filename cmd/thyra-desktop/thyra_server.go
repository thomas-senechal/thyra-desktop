package main

import (
	"log"
	"os/exec"
)

func getThyraServerVersion() string {
	cmd := exec.Command("thyra-server", "--version")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
