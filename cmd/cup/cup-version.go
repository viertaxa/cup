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
)

// Define the `version` command. It's just prints some compile time variables.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Run:   executeVersion,
}

// Register the subcommand with the root command
func init() {
	rootCmd.AddCommand(versionCmd)
}

// Print some version information
func executeVersion(_ *cobra.Command, _ []string) {
	log.Printf("Version: %s\nCommit: %s\nBuild Type: %s\n", version, commit, buildType)
}
