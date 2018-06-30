package version

import (
	"fmt"
	"time"
)

var builtTime string

// build: $ go build -ldflags "-X main.gitRevision=$(git rev-parse HEAD)"
var gitRevision string

func init() {
	builtTime = fmt.Sprintf("%s", time.Now())
}

// Print prints your application's version
func Print() {
	if gitRevision == "" {
		gitRevision = "(not set)"
	}
	fmt.Println(fmt.Sprintf("git revision:	%s", gitRevision))
	fmt.Println(fmt.Sprintf("built at:	%s", builtTime))
}
