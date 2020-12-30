/*
* Copyright 2021 Taylor Vierrether
*
* This program is free software: you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/viertaxa/cup/internal/commands"
	"github.com/viertaxa/cup/internal/config"
	"github.com/viertaxa/cup/internal/util"
	"os"
)

// Define the `configure` command. It's used to generate or update the configuration file.
var configureCmd = &cobra.Command{
	Use:   "configure [domain]",
	Short: "Configure CUP.",
	Long: "Configure CUP with Cloudflare credentials and tell it what hosts on what domains should be updated to the " +
		"current public IP.",
	Run: executeConfigure,
	// Enable flags from the root command
	TraverseChildren: true,
}

// Register the subcommand on the root command
func init() {
	rootCmd.AddCommand(configureCmd)
}

// Update or Create the configuration
func executeConfigure(_ *cobra.Command, _ []string) {
	// Check if we're interactive, exit if not. Configure is only supported interactively.
	isStdIn, err := util.CheckIsStdin()
	if err != nil {
		// There was an error checking if we're interactive, but we'll continue in case there's some weird terminal that
		// is interactive but doesn't act like it.
		log.Errorf("Determining if we're in an interactive shell failed. Assuming interactive\nError: %s", err)
	}
	if !isStdIn {
		log.Fatal("We only support running the configure command in an interactive session.")
	}
	// Attempt to read in an existing config. If one is not found at the specified path, create it, otherwise update it.
	// The config file was set earlier in execution in the initConfig() function
	if err := viper.ReadInConfig(); err != nil {
		// There was an error reading the config file, but that error might be that it wasn't there.
		// TODO: play around with passing in the config file option and testing these failure catching mechanisms
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// The error reading the config file was indeed that it wasn't found.
			// NOTE: This error only gets returned when we're using the automatic config file locating system of
			// viper. If the user passes in a specific file, a general error will be returned.
			// We'll set the loglevel now. There's no config file to read from, so we'll just use what the user passed
			// in on the command line, or the default value if that was not done.
			log.Warnf("No config file found at specified path! Creating a new one.\nError:%s", err)
			createConfig()
		} else {
			// Something went wrong when reading the specified file, or the user passed in a config file that doesn't
			// exist yet.
			// Let's check to see if the file specified exists. If it doesn't we'll create it. If it does, we'll log a
			// fatal error.
			exists, err := util.CheckFileExists(config.FilePath)
			if err != nil {
				log.Fatalf("Error reading configuration file specified.\nError:%s", err)
			} else {
				if !exists {
					log.Warnf("No config file found at specified path! Creating a new one.\nError:%s", err)
					createConfig()
				} else {
					log.Fatalf("Config file found at %s, but we were unable to read it. Has it been hand edited?"+
						"\nError: %s", config.FilePath, err)
				}
			}
		}
	} else {
		// The configuration was loaded successfully
		modifyConfig()
	}
	// Configuration struct should be up to date now, regardless of if it was there to begin with or not.

	// Actually write the configuration to disk.
	log.Infof("Writing configuration file %s to disk.", config.FilePath)
	if err := viper.WriteConfigAs(config.FilePath); err != nil {
		// Something went wrong writing the configuration file to disk. Prompt to dump the struct in JSON format to the
		// terminal in case the user went through a log of effort creating it, and would like to make the file manually.
		log.Errorf("Error writing configuration to disk!\nError: %s", err)
		dumpConfigIfDesired()
		// Wrap things up with a fatal error and exit
		log.Fatalf("Exiting.")
	} else {
		// Updating the config and config file seems to have gone well. We're done now.
		log.Infof("Successfully wrote config file to %s", config.FilePath)
		// Exit with a success exit code.
		os.Exit(0)
	}
}

func dumpConfigIfDesired() {
	if shouldDump, err := util.AskUserForBool("Should we dump the config to the shell for you?",
		false); err != nil {
		// Error prompting if we should dump the config. Defaulting to false for security purposes.
		log.Errorf("Error getting your response. We will not dump the config for security reasons."+
			"\nError:%s", err)
	} else {
		// Got a valid response, act accordingly
		if shouldDump {
			// User wants a config dump
			// Get the YAML
			if yaml, err := util.PrettyPrintSYaml(config.C); err != nil {
				// Error getting the YAML
				log.Errorf("Error when attempting to marshal config to YAML.\nError: %s", err)
			} else {
				// Got YAML back
				log.Printf("\n%s\n", yaml)
			}
		}
	}
}

func createConfig() {
	setLogLevel()

	// Config file was not found at given path, let's create it and warn the user that this is the case.
	// Generate the configuration file
	if createErr := commands.ConfigurationCreate(); createErr != nil {
		// Something went wrong when creating the new configuration struct. Log it and exit.
		log.Fatalf("Unable to create configuration.\nError:%s", createErr)
	}
	// Configuration struct was created successfully
	log.Info("Config created successfully.")
}

func modifyConfig() {
	// Config file was found, lets modify it.

	// First we copy the defaults into the config struct. This handles the case where we add new config options in the
	// future
	config.C = config.GetDefaultCupConfig()

	// Next we unmarshal the configuration loaded from the file into our config struct since we prefer to use that
	// over the viper methods to query configuration.
	unmarshalErr := viper.Unmarshal(&config.C)
	if unmarshalErr != nil {
		// There was some sort of error unmarshalling the configuration loaded from the file into our struct.
		// This was likely caused by a config file that was missing properties due to hand editing.
		log.Fatalf("Unable to unmarshal config. Has it been hand edited?\nError: %s", unmarshalErr)
	}
	// Now that we have loaded the user's preferences from file, set the log level appropriately.
	setLogLevel()

	// Update the configuration struct that we loaded previously
	log.Info("Configuration file located. Modifying...")
	// Set the actual used configuration file path in our configuration
	config.FilePath = viper.ConfigFileUsed()
	if err := commands.ConfigurationEdit(); err != nil {
		// Something went wrong updating the configuration struct.
		log.Fatalf("Error updating configuration file.\nError:%s", err)
	}
}
