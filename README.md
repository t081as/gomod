# GoMod
Package gomod provides a parser for go.mod files according to the official specification.
Since this package heavily relies on the `go` command, a go installation is required for this module to work.

The documentation is available here: https://pkg.tk-software.de/gomod

```go
package main

import (
	"fmt"

	"pkg.tk-software.de/gomod"
)

func main() {
	mod, err := gomod.NewFromDir("./testdata/default/")
	if err != nil {
		return
	}

	fmt.Println("Module path:", mod.Module.Path)
	fmt.Println("Go version:", mod.Go)
}

```

## Contributing
see [CONTRIBUTING.md](CONTRIBUTING.md)

## Donating
Thanks for your interest in this project. You can show your appreciation and support further development by [donating](https://www.tk-software.de/donate).

## License
**GoMod** © 2023 [Tobias Koch](https://www.tk-software.de). Released under a [BSD-style license](https://gitlab.com/tobiaskoch/gomod/-/blob/main/LICENSE).