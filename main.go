// main
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	dirFlag := flag.Bool("dir", false, "Output the directory the configuration is stored in")

	// Get all engines from the file
	engines := GetEngines()

	// Store engine flags in a map
	engineFlags := make(map[string]*bool)

	// Loop through engines and check if flag was supplied, store in engineFlags
	for _, engine := range engines {
		engineFlags[engine.Command] = flag.Bool(engine.Command, false, fmt.Sprintf("Use %s search engine", engine.Command))
	}

	argv := os.Args
	// Assume the last argument is the search query
	// Make this better in future, so the "" is not not required for multiple words
	searchQuery := argv[len(argv)-1]

	flag.Parse()

	if *dirFlag {
		fmt.Println(GetConfigDir())
	}

	if searchQuery == "" {
		fmt.Println("No search query was supplied.")
		os.Exit(1)
	}

	// Loop through engines and perform search if query was supplied
	for _, engine := range engines {
		if *engineFlags[engine.Command] {
			engine.Search(searchQuery)
		}
	}
}
