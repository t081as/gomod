package gomod

import "testing"

func TestNewFromDir(t *testing.T) {
	mod, err := NewFromDir("./testdata/default")

	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}

	// Module
	if s := mod.Module.Path; s != "example.com/mymodule" {
		t.Fatalf("Expected module path example.com/mymodule, got %s", s)
	}

	// Go
	if s := mod.Go; s != "1.14" {
		t.Fatalf("Expected go version 1.14, got %s", s)
	}

	// Require
	if l := len(mod.Require); l != 4 {
		t.Fatalf("Expected 4 module requirements, got %d", l)
	}

	if !expectRequire(mod.Require, "example.com/othermodule", "v1.2.3", false) {
		t.Fatalf("Expected module requirement %s not found or incorrect", "example.com/othermodule")
	}

	if !expectRequire(mod.Require, "example.com/thismodule", "v1.6.3", false) {
		t.Fatalf("Expected module requirement %s not found or incorrect", "example.com/thismodule")
	}

	if !expectRequire(mod.Require, "example.com/thatmodule", "v1.1.3", false) {
		t.Fatalf("Expected module requirement %s not found or incorrect", "example.com/thatmodule")
	}

	if !expectRequire(mod.Require, "example.com/anothermodule", "v1.7.3", true) {
		t.Fatalf("Expected module requirement %s not found or incorrect", "example.com/anothermodule")
	}

	// Replace
	if l := len(mod.Replace); l != 2 {
		t.Fatalf("Expected 2 module replacements, got %d", l)
	}

	if !expectReplace(mod.Replace, "example.com/thatmodule", "", "../thatmodule", "") {
		t.Fatalf("Expected module replacement %s not found or incorrect", "example.com/thatmodule")
	}

	if !expectReplace(mod.Replace, "example.com/amodule", "v1.2.3", "example.com/amodule", "v1.2.4") {
		t.Fatalf("Expected module replacement %s not found or incorrect", "example.com/amodule")
	}

	// Exclude
	if l := len(mod.Exclude); l != 1 {
		t.Fatalf("Expected 1 module exclusion, got %d", l)
	}

	if !expectExclude(mod.Exclude, "example.com/thismodule", "v1.3.0") {
		t.Fatalf("Expected module exclusion %s not found or incorrect", "example.com/thismodule")
	}

	// Retract
	if l := len(mod.Retract); l != 2 {
		t.Fatalf("Expected 2 module retraction, got %d", l)
	}

	if !expectRetract(mod.Retract, "v1.1.0", "v1.1.0", "broken") {
		t.Fatalf("Expected module retraction %s not found or incorrect", "broken")
	}

	if !expectRetract(mod.Retract, "v1.1.5", "v1.1.2", "bug") {
		t.Fatalf("Expected module retraction %s not found or incorrect", "bug")
	}
}

func expectRequire(r []Require, path, version string, indirect bool) bool {
	for _, e := range r {
		if e.Path == path && e.Version == version && e.Indirect == indirect {
			return true
		}
	}

	return false
}

func expectReplace(r []Replace, opath, oversion, npath, nversion string) bool {
	for _, e := range r {
		if e.Old.Path == opath && e.Old.Version == oversion && e.New.Path == npath && e.New.Version == nversion {
			return true
		}
	}

	return false
}

func expectExclude(e []Exclude, path, version string) bool {
	for _, c := range e {
		if c.Path == path && c.Version == version {
			return true
		}
	}

	return false
}

func expectRetract(r []Rectract, high, low, rationale string) bool {
	for _, e := range r {
		if e.High == high && e.Low == low && e.Rationale == rationale {
			return true
		}
	}

	return false
}

func TestNewFromDirInvalid(t *testing.T) {
	_, err := NewFromDir("/i-do-not-exist")

	if err == nil {
		t.Error("Expected error, got none")
	}
}
