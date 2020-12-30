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
	"github.com/viertaxa/cup/internal/config"
	"github.com/viertaxa/cup/internal/types"
	"os"
)

const envPrefix = "CUP"

// To be set by the compiler/linker at build time.
var version = "0.0.0"
var commit = "x"
var buildType string

// To be set during Init
var CupInfo types.CupInfo

// Create the "root command" for cobra. It doesn't do anything but print usage info.
var rootCmd = &cobra.Command{
	Use:   "cup [command]",
	Short: "CUP - Cloudflare (DNS) Updater Program",
	Long: "CUP is a Cloudflare (DNS) Updater Program that uses the Cloudflare API and a quorum of public " +
		"'What is my IP' services to create what amounts to a dynamic DNS service.",
}

func init() {
	// Generate a new CupInfo for us to inject
	CupInfo = types.CupInfo{
		Version:   version,
		Commit:    commit,
		BuildType: buildType,
	}

	// logrus configuration
	// If we're building a development build, include the calling function in logging for debugging purposes.
	if buildType == "development" {
		log.SetReportCaller(true)
	}
	// Log to stdout rather than stderr as is the default.
	log.SetOutput(os.Stdout)
	// Use detailed timestamps rather than execution time.
	formatter := &log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02-15-04-05.000-0700",
	}
	log.SetFormatter(formatter)

	// Cobra configuration
	// Add root level flags to the root command
	rootCmd.PersistentFlags().StringVar(&config.FilePath, "config", "", "Configuration "+
		"file path (default \"./cupconf.{yaml, json, toml}\")")
	rootCmd.PersistentFlags().StringVar(&config.C.LogLevel, "loglevel", "", "Set the maximum log "+
		"level.")

	// Initialize Cobra/Viper
	cobra.OnInitialize(initConfig)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	// Enable environment variable input
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()

	// if the config file path was not passed in, set it to the default config file path
	if config.FilePath == "" {
		viper.SetConfigFile(config.DefaultCupConfigPath)
		config.FilePath = config.DefaultCupConfigPath
	} else {
		// If the config file path was passed in, tell viper look for it there.
		viper.SetConfigFile(config.FilePath)
	}
}
