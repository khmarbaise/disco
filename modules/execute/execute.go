// Copyright 2021 The Disco Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
