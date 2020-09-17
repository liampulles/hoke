package app

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/liampulles/hoke/pkg/http"
)

type programArgs struct {
	Sleep int
	URL   string
}

// Run is the entrypoint for hoke. We delegate logic here
// so that it can be tested.
func Run(args []string) int {
	progArgs, res := extractArgs(args)
	if res != 0 {
		return res
	}

	fmt.Fprintf(os.Stderr, "Checking %s every %d seconds\n", progArgs.URL, progArgs.Sleep)
	pollUntilHealthy(progArgs)
	return 0
}

func pollUntilHealthy(progArgs *programArgs) {
	for !http.GetReturns2xx(progArgs.URL) {
		fmt.Fprintf(os.Stderr, "Failed to call %s, waiting %d seconds...\n", progArgs.URL, progArgs.Sleep)
		time.Sleep(time.Duration(progArgs.Sleep) * time.Second)
	}
}

func extractArgs(args []string) (*programArgs, int) {
	flagSet := flag.NewFlagSet("hoke", flag.ContinueOnError)
	sleepPtr := flagSet.Int("sleep", 5, "Sleep duration")
	if err := flagSet.Parse(args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "flag parse error: %v\n", err)
		return nil, 1
	}
	if flagSet.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "positional arg error: expecting 1 positional argument for the url, but got %d\n", flagSet.NArg())
		return nil, 2
	}
	url := flagSet.Arg(0)
	if *sleepPtr < 0 {
		fmt.Fprintf(os.Stderr, "sleep flag error: sleep should be larger than 0, but was %d\n", *sleepPtr)
		return nil, 3
	}

	return &programArgs{
		Sleep: *sleepPtr,
		URL:   url,
	}, 0
}
