//go:build windows

package utils

import "golang.org/x/sys/windows"

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
