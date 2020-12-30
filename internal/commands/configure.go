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

package commands

import (
	log "github.com/sirupsen/logrus"
	"github.com/viertaxa/cup/internal/config"
	"github.com/viertaxa/cup/internal/util"
)

// Populate the configuration struct
func ConfigurationCreate() error {
	log.Info("Creating a new configuration")

	// Start prompting the user for overall configuration values
	config.C.LogLevel = promptLogLevel(config.C.LogLevel)
	config.C.API.ApiBaseUrl = promptApiBaseUrl(config.GetDefaultCupConfig().API.ApiBaseUrl)
	config.C.API.ApiKey = promptAPIKey(config.GetDefaultCupConfig().API.ApiKey)

	log.Info("We will now begin setting up the services we will use to determine your public IP.")
	// Prompt user for what IP Reporting services we should use
	config.C.SimpleIPReportServices = promptSimpleIPServices(config.GetDefaultCupConfig().SimpleIPReportServices)

	// Start prompting for domain/host configuration
	log.Info("We will now begin setting up domains and hosts.")
	config.C.Domains = promptDomainsAndHosts(config.GetDefaultCupConfig().Domains)

	// Back-populate viper with the configuration struct so we can later save it to disk
	if err := util.LoadViperConfigFromStruct(config.C); err != nil {
		// Error populating viper, pass the error along.
		return err
	} else {
		// Config struct populated and viper populated successfully
		return nil
	}
}

// Update an existing configuration
func ConfigurationEdit() error {
	log.Infof("Modifying existing configuration at %s.", config.FilePath)

	// Start prompting the user for overall configuration values
	config.C.LogLevel = promptLogLevel(config.C.LogLevel)
	config.C.API.ApiBaseUrl = promptApiBaseUrl(config.C.API.ApiBaseUrl)
	config.C.API.ApiKey = promptAPIKey(config.C.API.ApiKey)

	log.Info("We will now begin setting up the services we will use to determine your public IP.")
	// Prompt user for what IP Reporting services we should use
	config.C.SimpleIPReportServices = promptSimpleIPServices(config.C.SimpleIPReportServices)

	// Start prompting for domain/host configuration
	log.Println("We will now begin setting up domains and hosts.")
	config.C.Domains = promptDomainsAndHosts(config.C.Domains)

	// Back-populate viper with the configuration struct so we can later save it to disk
	if err := util.LoadViperConfigFromStruct(config.C); err != nil {
		// Error populating viper, pass the error along.
		return err
	} else {
		// Config struct populated and viper populated successfully
		return nil
	}
}
