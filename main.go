package main

import (
	"fmt"
	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
	"github.com/hayashiki/gcal-timetracker/handler"
	"log"
	"os"
	"unicode"
)

const repo = "hayashiki/gcal-timetracker"

func main() {
	wf := aw.New(update.GitHub(repo))
	wf.Run(func() {
		run(wf)
	})
}

func run(wf *aw.Workflow) {
	args := wf.Args()
	if len(args) == 0 {
		exitWithError("please provide some input 👀")
	}

	handlers := map[string]func(*aw.Workflow, []string) (string, error) {
		"authorize": handler.DoAuthorize,
		"install": handler.DoInstall,
		"setup": handler.DoSetup,
		"track": handler.DoTrack,
		"start":       handler.DoStart,
		"stop":        handler.DoStop,
	}

	h, found := handlers[args[0]]

	if !found {
		exitWithError(args[0])
	}

	msg, err := h(wf, args[1:])
	if err != nil {
		exitWithError(err.Error())
		os.Exit(1)
	}

	if msg != "" {
		fmt.Print(msg)
	}
}

func exitWithError(msg string) {
	fmt.Print(capitalize(msg))
	log.Print(msg)
	os.Exit(1)
}

func capitalize(msg string) string {
	return string(unicode.ToUpper(rune(msg[0]))) + msg[1:]
}
