package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"time"
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
		log.Fatalf("cmd.Run() X rfailed with %s\n", err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	return Responce{outStr, errStr}
}
func handleError(commands string, output string, err string) {
	defer rec()
	fmt.Printf("cmd %v output %v err %v", commands, output, err)
}
func main() { //git status
	t := time.Duration(10 * time.Millisecond)
	for x := range time.Tick(t) {
		fmt.Printf("ticker %v", x)
		addres := ExecuteCmd("git", []string{"add", "."})
		handleError("add", addres.Output, addres.Err)
		//message will come from some file.
		commitRes := ExecuteCmd("git", []string{"commit", "-m", "Automated commit "})
		handleError("Commit", commitRes.Output, commitRes.Err)
		pushErr := ExecuteCmd("git", []string{"push"})
		handleError("Commit", pushErr.Output, pushErr.Err)

	}
}
func rec() {

	if r := recover(); r != nil {
		fmt.Println("Recovered in f", r)
	}
}
