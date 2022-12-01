package main

import (
	"log"
	"os/exec"
	"strings"
)

func getThyraServerVersion() string {
	cmd := exec.Command("thyra-server", "--version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	version := strings.Split(string(out), "Version: ")[1]
	return version
}
