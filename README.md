# version
Small package for helping to print your build version

## Usage
```
package main

import (
	"log"
	"os"

	"github.com/nguyendangminh/version"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "version" {
		version.Print()
		return
	}
	log.Println("error: invalid command")
	log.Printf("usage: %s version", os.Args[0])
}

```
- Build
```
go build -ldflags "-X github.com/nguyendangminh/version.gitRevision=$(git rev-parse HEAD)"
```

## Example

See example in `example` directory
```
$ go install -ldflags "-X github.com/nguyendangminh/version.gitRevision=$(git rev-parse HEAD)" github.com/nguyendangminh/version/example

$ example version
```

Note: There's a bug in (Go)[https://github.com/golang/go/issues/18369] that causes `go install` with `-ldflags` does not rebuild the binary. The bug is going to be fixed in version Go 1.10. Workaround is to delete the binary before compiling.