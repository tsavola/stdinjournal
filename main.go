// Copyright (c) 2021 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/coreos/go-systemd/v22/journal"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s syslog-identifier\n", os.Args[0])
		os.Exit(2)
	}

	vars := map[string]string{
		"SYSLOG_IDENTIFIER": os.Args[1],
	}

	if !journal.Enabled() {
		fmt.Fprintf(os.Stderr, "%s: journal is not available\n", os.Args[0])
		os.Exit(1)
	}

	r := bufio.NewReader(os.Stdin)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}

			msg := fmt.Sprintf("stdin read error: %v", err)
			journal.Send(msg, journal.PriCrit, nil)
			fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], msg)
			os.Exit(1)
		}

		if err := journal.Send(line, journal.PriErr, vars); err != nil {
			fmt.Fprintf(os.Stderr, "%s: journal error: %v\n", os.Args[0], err)
		}
	}
}
