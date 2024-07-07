// main
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"

	"github.com/adrg/xdg"
	"github.com/pkg/browser"
)

// Engine - Represents a search engine with its specified command and search query
type Engine struct {
	Command      string `json:"command"`
	SearchString string `json:"searchString"`
}

// Search - Search the specified searchQuery using this engine
func (e Engine) Search(searchQuery string) {
	browser.OpenURL(strings.ReplaceAll(e.SearchString, "[QUERY]", encodeURL(searchQuery)))
}

// encodeURL()
func encodeURL(urlString string) string {
	return url.QueryEscape(urlString)
}

var engines = make([]Engine, 0)

const enginesFile = "engines.json"

func getConfigDir() string {
	return xdg.ConfigHome + "/gee/"
}

// GetEngines - Get all engines
func GetEngines() []Engine {
	if len(engines) == 0 {
		loadEngines()
	}

	return engines
}

// loadEngines - Load all engines from the config file
func loadEngines() {
	filepath := getConfigDir() + enginesFile

	file, err := os.Open(filepath)
	if err != nil {
		panic(fmt.Sprintf("Failed opening '%s' file: %v", filepath, err))
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Sprintf("Failed reading content from '%s': %v", filepath, err))
	}

	err = json.Unmarshal(content, &engines)
	if err != nil {
		panic(fmt.Sprintf("Failed deserializing content from '%s': %v", filepath, err))
	}
}
