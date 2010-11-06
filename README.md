fork.pkg
========

A simple wrapper around exec.Exec()

Usage
-----

    package main

    import "fork"

    func main() {
		msg, err := fork.Exec("/path/to/command", []string { "arg1", "arg2", "argN" }, map[string]string { "NAME": "VALUE" })
	}