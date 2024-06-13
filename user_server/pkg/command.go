package pkg

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type ClosableCommand struct {
	Command *exec.Cmd
}

func (c *ClosableCommand) Close() error {
	return c.Command.Process.Kill()
}

func (c *ClosableCommand) Start() error {
	return c.Command.Start()
}
func (c *ClosableCommand) Wait() error {
	return c.Command.Wait()
}
func (c *ClosableCommand) Output() ([]byte, error) {
	return c.Command.Output()
}

func NewCMD(name string, arg ...string) ClosableCommand {
	outFile, err := os.OpenFile("api/assets/logs/server/cmd_out.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	outFileErr, err := os.OpenFile("api/assets/logs/server/cmd_err.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}

	cmd := &exec.Cmd{
		Path:   name,
		Args:   append([]string{name}, arg...),
		Stdout: outFile,
		Stderr: outFileErr,
	}

	log.Println(cmd.String())

	if filepath.Base(name) == name {
		if lp, err := exec.LookPath(name); err != nil {
			log.Fatal(err)
		} else {
			cmd.Path = lp
		}
	}
	return ClosableCommand{Command: cmd}
}

func ReadConsole(pipe io.ReadCloser) {
	i := true
	for {
		tmp := make([]byte, 1024)
		_, err := pipe.Read(tmp)
		fmt.Print(string(tmp))
		if len(string(tmp)) > 0 && i {
			i = false
		}
		if err != nil {
			break
		}
	}
}
