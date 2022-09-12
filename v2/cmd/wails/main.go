package main

import (
	"fmt"
	"os"

	"github.com/frankfang/wails/v2/cmd/wails/internal"
	"github.com/frankfang/wails/v2/cmd/wails/internal/commands/show"

	"github.com/frankfang/wails/v2/internal/colour"

	"github.com/frankfang/wails/v2/cmd/wails/internal/commands/update"

	"github.com/frankfang/wails/v2/cmd/wails/internal/commands/build"
	"github.com/frankfang/wails/v2/cmd/wails/internal/commands/dev"
	"github.com/frankfang/wails/v2/cmd/wails/internal/commands/doctor"
	"github.com/frankfang/wails/v2/cmd/wails/internal/commands/generate"
	"github.com/frankfang/wails/v2/cmd/wails/internal/commands/initialise"
	"github.com/leaanthony/clir"
)

func fatal(message string) {
	println(message)
	os.Exit(1)
}

func banner(_ *clir.Cli) string {
	return fmt.Sprintf("%s %s",
		colour.Green("Wails CLI"),
		colour.DarkRed(internal.Version))
}

func printFooter() {
	println(colour.Green("\nIf Wails is useful to you or your company, please consider sponsoring the project:\nhttps://github.com/sponsors/leaanthony\n"))
}

func main() {

	var err error

	app := clir.NewCli("Wails", "Go/HTML Appkit", internal.Version)

	app.SetBannerFunction(banner)
	defer printFooter()

	build.AddBuildSubcommand(app, os.Stdout)
	err = initialise.AddSubcommand(app, os.Stdout)
	if err != nil {
		fatal(err.Error())
	}

	err = doctor.AddSubcommand(app, os.Stdout)
	if err != nil {
		fatal(err.Error())
	}

	err = dev.AddSubcommand(app, os.Stdout)
	if err != nil {
		fatal(err.Error())
	}

	err = generate.AddSubcommand(app, os.Stdout)
	if err != nil {
		fatal(err.Error())
	}

	show.AddSubcommand(app, os.Stdout)

	err = update.AddSubcommand(app, os.Stdout, internal.Version)
	if err != nil {
		fatal(err.Error())
	}

	command := app.NewSubCommand("version", "The Wails CLI version")
	command.Action(func() error {
		println(internal.Version)
		return nil
	})

	err = app.Run()
	if err != nil {
		println("\n\nERROR: " + err.Error())
		printFooter()
		os.Exit(1)
	}
}
