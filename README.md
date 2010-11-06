fork.pkg
========

A simple wrapper around exec.Exec()

Installation
------------

	goinstall github.com/davecheney/fork

Usage
-----

    package main

	import "github.com/davecheney/fork"

    func main() {
		msg, err := fork.Exec("/path/to/command", []string { "arg1", "arg2", "argN" }, map[string]string { "NAME": "VALUE" })	
	}