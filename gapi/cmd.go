package main

import (
	"bytes"
	"context"
	"os/exec"
)

func executeCmd(ctx context.Context, command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.CommandContext(ctx, Shell, ShellArg, command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

// func executeCmd(cmdStr string) {
// 	cmd := exec.Command(cmdStr)
// 	// cmd.Shell = true
// 	output, err := cmd.Output()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(string(output))
// }
