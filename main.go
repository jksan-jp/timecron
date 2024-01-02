package main

import (
	"log"
	"os/exec"
	"time"
)

func executeScript(scriptPath string) {
	out, err := exec.Command("/bin/sh", scriptPath).Output()
	if err != nil {
		return
	}
	log.Printf("%s\n", out)
}

func main() {
	ticker := time.NewTicker(1 * time.Minute)

	for range ticker.C {
		currentTime := time.Now()
		formattedTime := currentTime.Format("200601021504")
		executeScript("./script/" + formattedTime + "/script.sh")
	}
}
