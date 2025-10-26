package cli

import (
	"os/exec"
	"runtime"
)

func PlaySound(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "darwin":
		cmd = "afplay"
	case "windows":
		cmd = "powershell"
		args = []string{"-c", "(New-Object Media.SoundPlayer '" + url + "').PlaySync()"}
	default: // linux, freebsd, etc.
		cmd = "play"
	}

	args = append(args, url)
	return exec.Command(cmd, args...).Run()
}
