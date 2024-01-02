// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gomod provides a parser for go.mod files according to the official specification.
// Since this package heavily relies on the `go` command, a go installation is required for this module to work.
//
//	mod, err := gomod.NewFromDir("./testdata/default/")
//	if err != nil {
//		return
//	}
//
//	fmt.Println("Module path:", mod.Module.Path)
//	fmt.Println("Go version:", mod.Go)
package gomod

import (
	"encoding/json"
	"os/exec"
	"path/filepath"
)

// A module is defined by a file named go.mod in its root directory.
// See https://go.dev/ref/mod#go-mod-file
type GoMod struct {

	// A module directive defines the main module’s path.
	Module Module `json:"Module"`

	// A go directive indicates that a module was written assuming the semantics of a given version of Go.
	Go string `json:"Go"`

	// A toolchain directive declares a suggested Go toolchain to use with a module.
	Toolchain string `json:"Toolchain"`

	// A require directive declares a minimum required version of a given module dependency.
	Require []Require `json:"Require"`

	// An exclude directive prevents a module version from being loaded by the go command.
	Exclude []Exclude `json:"Exclude"`

	// A replace directive replaces the contents of a specific version of a module, or all versions of a module, with contents found elsewhere.
	Replace []Replace `json:"Replace"`

	// A retract directive indicates that a version or range of versions of the module should not be depended upon.
	Retract []Rectract `json:"Retract"`
}

// A module directive defines the main module’s path.
// See https://go.dev/ref/mod#go-mod-file-module
type Module struct {

	// The main module’s path.
	Path string `json:"Path"`
}

// An exclude directive prevents a module version from being loaded by the go command.
// See https://go.dev/ref/mod#go-mod-file-exclude
type Exclude struct {

	// The path of the module to be excluded.
	Path string `json:"Path"`

	// The version of the module to be excluded.
	Version string `json:"Version"`
}

// A require directive declares a minimum required version of a given module dependency.
// See https://go.dev/ref/mod#go-mod-file-require
type Require struct {

	// The path of the module dependency.
	Path string `json:"Path"`

	// The minimum required version of the module dependency.
	Version string `json:"Version"`

	// Indicates that no package from the required module is directly imported.
	Indirect bool `json:"Indirect"`
}

// A replace directive replaces the contents of a specific version of a module, or all versions of a module, with contents found elsewhere.
// See https://go.dev/ref/mod#go-mod-file-replace
type Replace struct {

	// The module to be replaced.
	Old Old `json:"Old"`

	// The module replacement.
	New New `json:"New"`
}

// The module to be replaced.
type Old struct {

	// The path of the module.
	Path string `json:"Path"`

	// If a version is present only that specific version of the module is replaced.
	Version string `json:"Version"`
}

// The module replacement.
type New struct {

	// The path of the module.
	Path string `json:"Path"`

	// If the path is not a local path, it must be a valid module path. In this case, a version is required.
	Version string `json:"Version"`
}

// A retract directive indicates that a version or range of versions of the module should not be depended upon.
// See https://go.dev/ref/mod#go-mod-file-retract
type Rectract struct {

	// A directive may be written with either a single version or with a closed interval of versions with an upper and lower bound.
	Low string `json:"Low"`

	// A directive may be written with either a single version or with a closed interval of versions with an upper and lower bound.
	High string `json:"High"`

	// A directive should have a comment explaining the rationale for the retraction.
	Rationale string `json:"Rationale"`
}

// NewFromDir parses and returns the go.mod file located in the directory d.
func NewFromDir(d string) (*GoMod, error) {
	b, err := executeCmd(d, "go", "mod", "edit", "-json")
	if err != nil {
		return nil, err
	}

	var m GoMod

	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	return &m, nil
}

func executeCmd(dir, name string, arg ...string) ([]byte, error) {
	d, err := filepath.Abs(dir)
	if err != nil {
		return make([]byte, 0), err
	}

	cmd := exec.Command(name, arg...)
	cmd.Dir = d
	outp, err := cmd.Output()

	if err != nil {
		return make([]byte, 0), err
	}

	return outp, nil
}
