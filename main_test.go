// Copyright 2021 The Disco Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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
			want_stdout: "URL: https://api.foojay.io/disco/v1.0/distributions/oracle\nName: Oracle\nAPI Parameter: oracle\nNumber of versions: 79\n",
		},
		{
			args:        []string{"dist", "--name", "aoj"},
			want_stderr: "",
			want_stdout: "URL: https://api.foojay.io/disco/v1.0/distributions/aoj\nName: AOJ\nAPI Parameter: aoj\nNumber of versions: 51\n",
		},
		{
			args:        []string{"dist", "--name", "corretto"},
			want_stderr: "",
			want_stdout: "URL: https://api.foojay.io/disco/v1.0/distributions/corretto\nName: Corretto\nAPI Parameter: corretto\nNumber of versions: 30\n",
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
			args:        []string{"packages", "-v", "11", "--libc_type", "musl"},
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
