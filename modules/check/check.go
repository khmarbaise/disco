package check

import (
	"fmt"
	"os"
	"strings"
)

// IfError should be used to naively panics if an error is not nil.
func IfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// IfErrorWithOutput should be used to naively panics if an error is not nil.
func IfErrorWithOutput(err error, stdout string, stderr string) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))

	fmt.Println("---- stdout  ----")
	fmt.Println(stdout)
	fmt.Println("---- stderr  ----")
	fmt.Println(stderr)

	os.Exit(1)
}

//IsMainBranch check for "main" or "master" and return true, false otherwise.
func IsMainBranch(branch string) bool {
	branchWithoutSpaces := strings.TrimSpace(branch)
	if branchWithoutSpaces == "main" || branchWithoutSpaces == "master" {
		return true
	}
	return false
}
