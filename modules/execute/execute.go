// Copyright 2021 The disco Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package execute

import (
	"bytes"
	"os/exec"
)

//CommandOutput represents the output of a command execution (stdout, stderr).
type CommandOutput struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

//ExternalCommandWithRedirect Execute external command as subprocess and return stdout and stderr
//plus possible error.
func ExternalCommandWithRedirect(name string, cmd ...string) (result CommandOutput, err error) {
	c := exec.Command(name, cmd...)

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
