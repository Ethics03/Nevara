package shell

import (
	"os"
	"os/exec"

	"github.com/creack/pty"
	"golang.org/x/term"
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

oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}

	go func() {
		_, _ = os.Stdout.ReadFrom(ptmx)
	}()
	go func() {
		_, _ = ptmx.ReadFrom(os.Stdin)
	}()

	// Optional: restore terminal state on exit
	go func() {
		cmd.Wait()
		_ = term.Restore(int(os.Stdin.Fd()), oldState)
	}()

	return &Shell{
		Cmd: cmd,
		Pty: ptmx,
	},nil 
}

