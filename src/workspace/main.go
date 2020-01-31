package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

type Responce struct {
	Output string
	Err    string
}

func ExecuteCmd(commands string, list []string) Responce {
	cmd := exec.Command(commands, list...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() X failed with %s\n", err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	return Responce{outStr, errStr}
}
func main() { //git status
	_ = ExecuteCmd("git", []string{"add", "."})
	//message will come from some file.
	_ = ExecuteCmd("git", []string{"commit", "-m", "Automated commit "})
	_ = ExecuteCmd("git", []string{"push"})

}
