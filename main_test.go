// Copyright 2021, 2022 The Disco Authors
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

package main_test

import (
	"github.com/khmarbaise/disco/modules/execute"
	"strings"
	"testing"
)

func Test_Main_DifferentCommands(t *testing.T) {
	tests := []struct {
		args        []string
		want_stderr string
		want_stdout string
	}{
		{
			args:        []string{"dist", "--help"},
			want_stderr: "",
			want_stdout: "", //can not give sequence which must be equal. Need to reconsider.
		},
		{
			args:        []string{"dist", "--name", "oraclxe"},
			want_stderr: "Status Code: 500 Internal Server Error",
			want_stdout: "",
		},
		{
			args:        []string{"dist", "--name", "oracle"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"dist", "--name", "aoj"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"dist", "--name", "corretto"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"majorversions"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"majorversions", "--major-version", "11"},
			want_stderr: "Error: either --ea or --ga must be given",
			want_stdout: "",
		},
		{
			args:        []string{"majorversions", "-v", "11", "--ea"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"majorversions", "--major-version", "11", "--ea"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"majorversions", "--major-version", "11", "--ga"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"majorversions", "--major-version", "11", "--latest", "ga"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"majorversions", "-v", "11", "-l", "ea"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"majorversions", "-v", "11", "-l", "ga"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"majorversions", "-v", "11", "-l", "sts"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"majorversions", "-v", "11", "-l", "mts"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"majorversions", "-v", "11", "-l", "lts"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"packages"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"packages", "-v", "11", "--architecture", "mips"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"packages", "-v", "11", "--distro", "redhat"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"packages", "-v", "11", "--archive_type", "pkg"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"packages", "-v", "11", "--package_type", "jre"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"packages", "-v", "11", "--operating_system", "macos"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"packages", "-v", "11", "--os", "macos"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"packages", "-v", "11", "--libc_type", "musl"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"packages", "-v", "11", "--lt", "musl"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"packages", "-v", "11", "--release_status", "ga"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"packages", "-v", "11", "--rs", "ga"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"packages", "-v", "11", "--term_of_support", "lts"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"packages", "-v", "11", "--tos", "lts"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"packages", "-v", "11", "--bitness", "32"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"packages", "-v", "11", "-b", "32"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"packages", "-v", "11", "--javafx_bundled"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"packages", "-v", "11", "--fx"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"packages", "-v", "11", "--directly_downloadable"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"packages", "-v", "11", "--dd"},
			want_stderr: "",
			want_stdout: "",
		},
		{
			args:        []string{"packages", "-v", "11", "--latest", "overall"},
			want_stderr: "",
			want_stdout: "",
		},
	}
	for _, tt := range tests {
		description := strings.Join(tt.args, " ")
		t.Run(description, func(t *testing.T) {
			redirect, err := execute.ExternalCommandWithRedirect("./disco", tt.args...)
			if err != nil {
				stderr := strings.TrimSpace(redirect.Stderr) //remove leading line feed.
				if tt.want_stderr != stderr {
					t.Errorf("\n\nerrorous execution\n%v\n stdout: ---------------------\n%v\nstderr:--------------------\n%v", err, redirect.Stdout, redirect.Stderr)
				}
			}
			if len(tt.want_stdout) > 0 && strings.TrimSpace(redirect.Stdout) != strings.TrimSpace(tt.want_stdout) {
				t.Errorf("\n stdout: ---------------------\n%v\nstderr:--------------------\n%v", redirect.Stdout, redirect.Stderr)
			}
		})
	}

}
