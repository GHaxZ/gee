// main.go
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
)

func main() {
	// Initialize CLI app
	app := cli.NewApp()
	app.Name = "gee"
	app.Usage = "A CLI tool for performing searches using different engines"
	app.UsageText = "gee [flags] [search query]"

	dirCommand := &cli.Command{
		Name:  "dir",
		Usage: "Output the directory the configuration is stored in",
		Action: func(_ *cli.Context) error {
			fmt.Println(GetConfigDir())
			return nil
		},
	}

	app.Commands = []*cli.Command{dirCommand}

	// Get all engines from the file
	engines := GetEngines()

	// Define a flag to track if any engine-specific flag is set
	engineFlags := make(map[string]bool)

	// Engine-specific commands and flags
	for _, engine := range engines {
		command := engine.Command
		app.Flags = append(app.Flags, &cli.BoolFlag{
			Name:        command,
			Usage:       fmt.Sprintf("Enable %s search engine", command),
			DefaultText: strconv.FormatBool(engine.Default),
			Category:    "AVAILABLE SEARCH ENGINES:",
		})
		engineFlags[command] = false
	}

	// Set the action for the app
	app.Action = func(c *cli.Context) error {
		// Get the search query (last argument)
		args := c.Args()
		if args.Len() < 1 {
			return fmt.Errorf("no search query was supplied")
		}
		searchQuery := args.Get(args.Len() - 1)

		// Check if any engine-specific flag is set
		engineProvided := false
		for command := range engineFlags {
			if c.Bool(command) {
				engineFlags[command] = true
				engineProvided = true
			}
		}

		// Loop through engines and perform search if flag is set
		for _, engine := range engines {
			if (engineProvided && engineFlags[engine.Command]) || (!engineProvided && engine.Default) {
				engine.Search(searchQuery)
			}
		}

		return nil
	}

	// Run the CLI app
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
