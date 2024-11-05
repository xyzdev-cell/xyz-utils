package cmd

import (
	"context"
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
)

// 执行命令行命令
func ExecCommandWithTimeout(command string, timeout time.Duration, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd := exec.CommandContext(
		ctx,
		command,
		args...,
	)

	output, err := cmd.CombinedOutput()
	if ctx.Err() == context.DeadlineExceeded {
		return string(output), fmt.Errorf("进程超时, %w", ctx.Err())
	}
	if runtime.GOOS == "windows" && isWindowsUIChs() {
		s, err := simplifiedchinese.GB18030.NewDecoder().String(string(output))
		if err != nil {
			return string(output), err
		}
		return s, nil
	}
	return string(output), err
}
