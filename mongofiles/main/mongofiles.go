// Copyright (C) MongoDB, Inc. 2014-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

// Main package for the mongofiles tool.
package main

import (
	"github.com/mongodb/mongo-tools-common/log"
	"github.com/mongodb/mongo-tools-common/signals"
	"github.com/mongodb/mongo-tools-common/util"
	"github.com/mongodb/mongo-tools/mongofiles"

	"fmt"
	"os"
)

func main() {
	opts, err := mongofiles.ParseOptions(os.Args[1:])
	if err != nil {
		log.Logv(log.Always, err.Error())
		log.Logv(log.Always, "try 'mongofiles --help' for more information")
		os.Exit(util.ExitBadOptions)
	}

	signals.Handle()

	// print help, if specified
	if opts.PrintHelp(false) {
		os.Exit(util.ExitClean)
	}

	// print version, if specified
	if opts.PrintVersion() {
		os.Exit(util.ExitClean)
	}

	mf, err := mongofiles.New(opts)
	if err != nil {
		log.Logv(log.Always, err.Error())
		if setupErr, ok := err.(util.SetupError); ok {
			if setupErr.Code == util.ExitBadOptions {
				log.Logvf(log.Always, "try mongofiles --help for more information")
			}
			os.Exit(setupErr.Code)
		}
		os.Exit(util.ExitError)
	}
	defer mf.Close()

	output, err := mf.Run(true)
	if err != nil {
		log.Logvf(log.Always, "Failed: %v", err)
		os.Exit(util.ExitError)
	}
	fmt.Printf("%s", output)
}
