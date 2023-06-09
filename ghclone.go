package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

var gitProxy = "https://ghproxy.com/"

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("git repo is empty")
		return
	}

	gitRepo := args[1]

	cmd := exec.Command("git", "clone", "--progress", gitProxy+gitRepo)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}
	}()
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}
	}()
	_ = cmd.Wait()
}
