// Package confy is the configuration package for the liby library
package confy

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Config is the config struct with the
// settings / options for the confy library
type Config struct {
	Port        string `json:"port,omitempty"`
	Address     string `json:"address,omitempty"`
	SessionName string `json:"sessionName,omitempty"`
	Gzip        string `json:"gzip,omitempty"`
}

// createConfig creates a "config.json" file if
// the config file you specify does not exist.
func createConfig() (*os.File, error) {
	fmt.Println("Creating a default config file")
	tmpFile, err := os.Create("config.json")
	if err != nil {
		return nil, err
	}

	// Default config settings
	defaultConfig := Config{
		Port:        "8080",
		Address:     "localhost",
		SessionName: "liby",
		Gzip:        "on",
	}

	// Marshal the default config settings
	tmpDefaultConfig, err := json.Marshal(defaultConfig)
	if err != nil {
		return nil, err
	}

	// Write the default settings to the newly created config.json file
	_, err = tmpFile.Write(tmpDefaultConfig)
	if err != nil {
		return nil, err
	}

	fmt.Println("Default config file has been made (config.json)")
	return tmpFile, err
}

// Open will open and parse the config file that are specified
func Open(filepath string) error {
	file, err := os.Open(filepath)
	defer file.Close() // Close file when done with it
	// If the specified file does not exist
	// print the error and create a default config.json file
	if err != nil {
		fmt.Println(err)
		file, err = createConfig()
		if err != nil {
			return err // return error if you cant create the file
		}
	}

	// Default values / settings if empty
	config := Config{
		Port:        "8080",
		Address:     "localhost",
		SessionName: "liby",
		Gzip:        "on",
	}

	// initiate new decoder and decode config / settings
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return err // return error if cant decode
	}

	// Parse Gzip option
	switch strings.ToLower(config.Gzip) {
	case "on":
		config.Gzip = "true"
	case "off":
		config.Gzip = "false"
	default:
		// Gzip is "on" by default
		config.Gzip = "true"
	}

	// Set the environment key / values
	os.Setenv("port", config.Port)
	os.Setenv("address", config.Address)
	os.Setenv("sessionname", config.SessionName)
	os.Setenv("gzip", config.Gzip)

	return nil // return nil if no error
}
