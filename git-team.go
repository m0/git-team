package main

import (
	"github.com/hekmekk/git-team/core/handler"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

const (
	version = "v0.1.0"
	author  = "Rea Sand <hekmek@posteo.de>"
)

func main() {
	app := kingpin.New("git-team", "Command line interface for creating git commit templates provisioned with one or more co-authors.")

	app.HelpFlag.Short('h')
	app.Version(version)
	app.Author(author)

	enable := app.Command("enable", "Provisions a git-commit template with the provided co-authors. A co-author must either be an alias or of the shape \"Name <email>\"").Default()
	coauthors := enable.Arg("coauthors", "Git co-authors").Strings()

	disable := app.Command("disable", "Use default template")
	status := app.Command("status", "Print the current status")

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case enable.FullCommand():
		handler.EnableCommand(coauthors)
	case disable.FullCommand():
		handler.DisableCommand()
	case status.FullCommand():
		handler.Status()
	}
}
