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