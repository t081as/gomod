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
	fmt.Println()

	fmt.Println("Required:")
	for _, r := range mod.Require {
		if !r.Indirect {
			fmt.Printf("%s@%s\n", r.Path, r.Version)
		}
	}
	fmt.Println()

	fmt.Println("Required (indirect):")
	for _, r := range mod.Require {
		if r.Indirect {
			fmt.Printf("%s@%s\n", r.Path, r.Version)
		}
	}
	fmt.Println()

	fmt.Println("Replaced:")
	for _, r := range mod.Replace {
		if r.Old.Version != "" {
			fmt.Printf("%s@%s -> %s@%s\n", r.Old.Path, r.Old.Version, r.New.Path, r.New.Version)
		} else {
			fmt.Printf("%s -> %s\n", r.Old.Path, r.New.Path)
		}
	}
	fmt.Println()

	fmt.Println("Excluded:")
	for _, e := range mod.Exclude {
		fmt.Printf("%s@%s\n", e.Path, e.Version)
	}
	fmt.Println()

	fmt.Println("Retracted:")
	for _, r := range mod.Retract {
		if r.Low == r.High {
			fmt.Printf("%s: %s\n", r.Low, r.Rationale)
		} else {
			fmt.Printf("[%s, %s]: %s\n", r.Low, r.High, r.Rationale)
		}
	}
	fmt.Println()
}
```

## Contributing
see [CONTRIBUTING.md](CONTRIBUTING.md)

## Donating
Thanks for your interest in this project. You can show your appreciation and support further development by [donating](https://www.tk-software.de/donate).

## License
**GoMod** © 2023-2024 [Tobias Koch](https://www.tk-software.de). Released under a [BSD-style license](https://gitlab.com/tobiaskoch/gomod/-/blob/main/LICENSE).