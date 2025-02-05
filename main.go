// main package
package main

import (
	"fmt"
	"os"
)

// Version is the current value injected at build time.
var Version string

// Message is the standard message that client and daemon should exchange to trigger a notify.
var Message = "org.transactionalUpdate.Notify"

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Fprintf(os.Stderr, "usage: %s transactional-update-notifier [daemon|client]\n", os.Args[0])

		return
	}

	// Version flag
	if os.Args[1] == "-v" || os.Args[1] == "version" {
		fmt.Println(Version)

		return
	}

	if os.Args[1] == "daemon" {
		NotifyDaemon()

		return
	}

	if os.Args[1] == "client" {
		if len(os.Args) > 2 && os.Args[2] == "failure" {
			NotifyDaemonClient("failure")
			return
		}
		NotifyDaemonClient("success")
		return
	}
}
