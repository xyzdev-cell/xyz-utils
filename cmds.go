package utils

import (
	"context"
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"golang.org/x/sys/windows"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func isWindowsUIChs() bool {
	langIds, err := windows.GetUserPreferredUILanguages(windows.MUI_LANGUAGE_ID)
	if err != nil {
		return false
	}
	for _, langId := range langIds {
		if langId == "0804" {
			return true
		}
	}
	return false
}

// 执行命令行命令
func ExecCommand(command string, timeout time.Duration, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
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
