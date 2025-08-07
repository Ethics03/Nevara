package shell

import (
	"os"
	"os/exec"

	"github.com/creack/pty"
	
)

type Shell struct {
	Cmd *exec.Cmd
	Pty *os.File
}

func StartShell() (*Shell, error) {
	shell := "/bin/bash"
	

	cmd := exec.Command(shell,"-l")
	ptmx, err := pty.Start(cmd)
	if err != nil {
		return nil, err
	}

	return &Shell{
		Cmd: cmd,
		Pty: ptmx,
	},nil 
}

