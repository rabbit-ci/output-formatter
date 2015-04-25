package main

import (
	"fmt"
	"github.com/rabbit-ci/logstreamer"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	logger := log.New(os.Stdout, "", 0)
	logStreamerOut := logstreamer.NewLogstreamer(logger, "stdout", false)
	logStreamerErr := logstreamer.NewLogstreamer(logger, "stderr", true)
	command := strings.Join(os.Args[1:], " ")
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = logStreamerOut
	cmd.Stderr = logStreamerErr
	logStreamerErr.FlushRecord()
	err := cmd.Start()

	if err != nil {
		fmt.Printf("ERROR! Could not spawn command. %v\n", err.Error())
	}

	err = cmd.Wait()
	if err != nil {
		os.Exit(1)
	}
}
