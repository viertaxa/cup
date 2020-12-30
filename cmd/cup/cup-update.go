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
	"os"
)

// Define the `update` command. It's the `main` command that does the DNS update.
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the specified DNS record on Cloudflare",
	Long: "Queries several `What is my IP` type services for your public IP, and if there's a quorum, update the " +
		"record in Cloudflare.",
	Run: executeUpdate,
	// Enable flags from the root command
	TraverseChildren: true,
}

// Register the subcommand on the root command
func init() {
	rootCmd.AddCommand(updateCmd)
}

// Attempt to update the DNS record with the current public IP
func executeUpdate(_ *cobra.Command, _ []string) {
	// Lets read in the configuration
	if err := viper.ReadInConfig(); err != nil {
		// There was an error reading in the configuration
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// The error reading in the configuration was that it wasn't there.
			log.Fatalf("Config file not found! Please run 'cup configure' if you haven't already.\n%s",
				config.FilePath)
		} else {
			// There was some other error reading in the configuration file
			log.Fatalf("Error reading in the configuration file!\nError: %s", err)
		}
	}
	// Config loaded successfully
	// Unmarshal the config to our config struct, since we prefer using that over the viper interface
	unmarshalErr := viper.Unmarshal(&config.C)
	if unmarshalErr != nil {
		// Error unmarshalling the config into our struct.
		log.Fatalf("Unable to unmarshal config. Has it been hand edited?\nError: %s", unmarshalErr)
	}
	// Now that we have the full config, set the user's preferred log level.
	setLogLevel()

	// Attempt to update the DNS records specified
	if err := commands.Update(CupInfo); err != nil {
		// Error updating the records
		log.Fatalf("Error updating DNS records!\nError: %s", err)
	} else {
		// Records updated successfully. Inform the user.
		log.Info("Successfully updated all DNS records. Goodbye.")
		// Exit success
		os.Exit(0)
	}
}
