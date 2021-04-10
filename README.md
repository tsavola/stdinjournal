Send stdin to systemd journal.  Each line is sent separately.  Error priority
is used.  Syslog identifier must be specified as the command-line argument.

Build instructions:

	go install github.com/tsavola/stdinjournal@latest

Usage instructions:

	ls -l | stdinjournal ls

