package execute

import (
	"bytes"
	"log"
	"os"
	"os/exec"
)

//ExternalCommand Execute external command as subprocess.
func ExternalCommand(cmd ...string) {
	log.Printf("Executing : %s ...\n", cmd)
	c := exec.Command(cmd[0], cmd[1:]...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Start(); err != nil {
		log.Panicln(err)
	}
	if err := c.Wait(); err != nil {
		log.Panicln(err)
	}
}

//CommandOutput represents the output of a command execution (stdout, stderr).
type CommandOutput struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

//ExternalCommandWithRedirect Execute external command as subprocess and return stdout and stderr
//plus possible error.
func ExternalCommandWithRedirect(cmd ...string) (result CommandOutput, err error) {
	c := exec.Command(cmd[0], cmd[1:]...)

	var stdoutBuffer, stderrBuffer bytes.Buffer
	c.Stdout = &stdoutBuffer
	c.Stderr = &stderrBuffer
	if err := c.Start(); err != nil {
		result = CommandOutput{Stdout: stdoutBuffer.String(), Stderr: stderrBuffer.String(), ExitCode: c.ProcessState.ExitCode()}
		return result, err
	}
	if err = c.Wait(); err != nil {
		result = CommandOutput{Stdout: stdoutBuffer.String(), Stderr: stderrBuffer.String(), ExitCode: c.ProcessState.ExitCode()}
		return result, err
	}
	return CommandOutput{Stdout: stdoutBuffer.String(), Stderr: stderrBuffer.String(), ExitCode: c.ProcessState.ExitCode()}, nil
}

//ExternalCommandInteractive Execute external command as subprocess in interactive mode.
func ExternalCommandInteractive(cmd ...string) (int, error) {
	c := exec.Command(cmd[0], cmd[1:]...)

	c.Stdout = os.Stdin
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	if err := c.Start(); err != nil {
		return c.ProcessState.ExitCode(), err
	}
	if err := c.Wait(); err != nil {
		return c.ProcessState.ExitCode(), err
	}
	return c.ProcessState.ExitCode(), nil
}
