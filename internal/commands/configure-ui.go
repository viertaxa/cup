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
	"strings"
)

// Get the user's desired loglevel
func promptLogLevel(existingOrDefault string) string {
	// Loop until we get something satisfactory
	for {
		log.Println("Please enter the log level you would like cup to operate at.")
		log.Printf("Valid values are: %s", strings.Join(config.GetValidLogLevels(), ", "))
		response, err := util.AskUserForString("Level", existingOrDefault)
		if err != nil {
			log.Errorf("Error getting your response. Please try again.\n%s", err)
		} else if util.StringSliceContains(config.GetValidLogLevels(), strings.ToLower(response)) {
			return strings.ToLower(response)
		} else {
			log.Errorf("%s is not a valid log level. Please try again.", response)
		}
	}
}

// Get the desired baseurl
func promptApiBaseUrl(existingOrDefault string) string {
	//Loop until we get something satisfactory
	for {
		log.Println("Please enter the base URL for the CloudFlare API.")
		log.Println("Unless you really know what you're doing, please leave the default by pressing the " +
			"ENTER/RETURN key. If you provide a value, do not include a trailing `/`.")
		response, err := util.AskUserForString("API Base", existingOrDefault)
		if err != nil {
			log.Errorf("Error getting your response. Please try again.\nError: %s", err)
		}
		// Check for HTTPS. Not only is it a good idea, but Cloudflare rejects http connections.
		if strings.HasPrefix(strings.ToLower(response), "https://") {
			return strings.ToLower(response)
		} else {
			log.Warnf("%s does not begin with 'https://'. This will not work with the cloudflare API. "+
				"Continue anyways?", response)
			shouldContinue, err := util.AskUserForBool("Continue?", false)
			if err != nil {
				log.Errorf("Error getting your response.\nError:%s", err)
			} else {
				if shouldContinue {
					return strings.ToLower(response)
				}
			}
		}
	}
}

// Get the user's API key
func promptAPIKey(existingOrDefault string) string {
	// Loop until we get something satisfactory
	for {
		log.Println("Please enter the API Key you will be using")
		response, err := util.AskUserForString("API Key", existingOrDefault)
		if err != nil {
			log.Errorf("Error getting your response. Please try again.\n%s", err)
		} else {
			return response
		}
	}
}

// Get the list of hosts and their domains to update
func promptDomainsAndHosts(existingOrDefault []config.DomainConfig) []config.DomainConfig {
	var newDomainsAndHosts []config.DomainConfig

	// Loop over the existing or example configuration.
	for index, domainConfig := range existingOrDefault {
		// Loop on the domain entry until it is successfully updated or the user deletes it.
		for {
			var updatedDomain string
			var updatedHosts []string

			// If we're configuring a domain of index 1 or higher, allow the user to remove the domain from the config
			if index != 0 {
				log.Printf("Moving on to domain #%d, %s, with hosts: %s", index+1, domainConfig.Domain,
					strings.Join(domainConfig.Hosts, ","))
				keep, err := util.AskUserForBool("Should we keep this domain in the config?", true)
				if err != nil {
					log.Errorf("Error getting your response, please try again.\nError: %s", err)
					continue
				} else {
					if !keep {
						// Move on to the next domain without appending this domain to newDomainsAndHosts
						break
					}
				}
			}
			// Gather info on the domain we're iterating on
			log.Println("Please enter the domain you will be controlling records on.")
			domainResponse, err := util.AskUserForString("Domain", domainConfig.Domain)
			if err != nil {
				log.Errorf("Error getting your response. Please try again.\nError: %s", err)
				continue
			} else {
				updatedDomain = strings.ToLower(domainResponse)
			}

			// Move on to getting the list of hosts
			log.Println("Please enter a comma separated list of hosts that should be updated. Use @ for the " +
				"root domain")
			hostResponse, err := util.AskUserForString("Hosts", strings.Join(domainConfig.Hosts, ","))
			if err != nil {
				log.Errorf("Error getting your response. Please try again.\nError: %s", err)
				continue
			} else {
				updatedHosts = strings.Split(strings.ReplaceAll(hostResponse, " ", ""), ",")
			}
			// Add the information to the new list of domains to be updated
			newDomainsAndHosts = append(newDomainsAndHosts, config.DomainConfig{
				Domain: updatedDomain,
				Hosts:  updatedHosts,
			})
			// We're done with this domain, so move on.
			break
		}
	}
	// Allow user to enter additional domains
	for {
		add, err := util.AskUserForBool("Should we add an additional domain?", false)
		if err != nil {
			log.Errorf("Error getting your response, please try again.\n%s", err)
			continue
		} else {
			if !add {
				// Don't want to add another, move on.
				break
			}
		}

		// Does want to add another
		var thisDomain string
		var thisHosts []string

		var exampleDomainConfig = config.GetDefaultCupConfig().Domains[0]

		// Gather info on the new domain to be added
		log.Println("Please enter the domain you will be controlling records on.")
		domainResponse, err := util.AskUserForString("Domain", exampleDomainConfig.Domain)
		if err != nil {
			log.Errorf("Error getting your response. Please try again.\n%s", err)
			continue
		} else {
			thisDomain = strings.ToLower(domainResponse)
		}

		// Gather the info on the hosts to be added
		log.Println("Please enter a comma separated list of hosts that should be updated. Use @ for the root " +
			"domain")
		hostResponse, err := util.AskUserForString("Hosts", strings.Join(exampleDomainConfig.Hosts, ","))
		if err != nil {
			log.Errorf("Error getting your response. Please try again.\n%s", err)
			continue
		} else {
			thisHosts = strings.Split(strings.ReplaceAll(hostResponse, " ", ""), ",")
		}
		newDomainsAndHosts = append(newDomainsAndHosts, config.DomainConfig{
			Domain: thisDomain,
			Hosts:  thisHosts,
		})
	}
	return newDomainsAndHosts
}

// Get the list of hosts and their domains to update
func promptSimpleIPServices(existingOrDefault []config.SimpleIPReporterConfig) []config.SimpleIPReporterConfig {
	var newSimpleIPReporterServices []config.SimpleIPReporterConfig

	// Loop over the existing or example configuration.
	for index, svc := range existingOrDefault {
		// Loop on the domain entry until it is successfully updated or the user deletes it.
		for {
			var updatedApiUrl string
			var updatedSvcName string

			// If we're configuring a service of index 1 or higher, allow the user to remove the service from the config
			if index != 0 {
				log.Printf("Moving on to service #%d, %s", index+1, svc.ServiceName)
				keep, err := util.AskUserForBool("Should we keep this service in the config?", true)
				if err != nil {
					log.Errorf("Error getting your response, please try again.\nError: %s", err)
					continue
				} else {
					if !keep {
						// Move on to the next service without appending this one to newSimpleIPReporterServices
						break
					}
				}
			}
			// Gather info on the domain we're iterating on
			log.Println("Please enter the name of the IP Reporting Service.")
			serviceNameResponse, err := util.AskUserForString("Service Name", svc.ServiceName)
			if err != nil {
				log.Errorf("Error getting your response. Please try again.\nError: %s", err)
				continue
			} else {
				updatedSvcName = serviceNameResponse
			}

			// Get the API Base
			log.Println("Please enter the API Base URL for this service.")
			hostResponse, err := util.AskUserForString("API Base", svc.IPv4ApiBase)
			if err != nil {
				log.Errorf("Error getting your response. Please try again.\nError: %s", err)
				continue
			} else {
				// Check for HTTPS. Not only is it a good idea, but Cloudflare rejects http connections.
				if strings.HasPrefix(strings.ToLower(hostResponse), "https://") {
					updatedApiUrl = strings.ToLower(hostResponse)
				} else {
					log.Warnf("%s does not begin with 'https://'. This is highly insecure. "+
						"Continue anyways?", hostResponse)
					shouldContinue, err := util.AskUserForBool("Continue?", false)
					if err != nil {
						log.Errorf("Error getting your response.\nError:%s", err)
					} else {
						if shouldContinue {
							updatedApiUrl = strings.ToLower(hostResponse)
						}
					}
				}
			}
			// Add the information to the new list of domains to be updated
			newSimpleIPReporterServices = append(newSimpleIPReporterServices, config.SimpleIPReporterConfig{
				ServiceName: updatedSvcName,
				IPv4ApiBase: updatedApiUrl,
			})
			// We're done with this service, so move on.
			break
		}
	}
	// Allow user to enter additional services
	for {
		add, err := util.AskUserForBool("Should we add an additional service?", false)
		if err != nil {
			log.Errorf("Error getting your response, please try again.\n%s", err)
			continue
		} else {
			if !add {
				// Don't want to add another, move on.
				break
			}
		}

		// Does want to add another
		var updatedApiUrl string
		var updatedSvcName string

		// Gather info on the domain we're iterating on
		log.Println("Please enter the name of the IP Reporting Service you wish to use.")
		serviceNameResponse, err := util.AskUserForString("Service Name", config.GetDefaultCupConfig().
			SimpleIPReportServices[0].ServiceName)
		if err != nil {
			log.Errorf("Error getting your response. Please try again.\nError: %s", err)
			continue
		} else {
			updatedSvcName = serviceNameResponse
		}

		// Get the API Base
		log.Println("Please enter the API Base URL for this service.")
		hostResponse, err := util.AskUserForString("API Base", config.GetDefaultCupConfig().
			SimpleIPReportServices[0].IPv4ApiBase)
		if err != nil {
			log.Errorf("Error getting your response. Please try again.\nError: %s", err)
			continue
		} else {
			// Check for HTTPS.
			if strings.HasPrefix(strings.ToLower(hostResponse), "https://") {
				updatedApiUrl = strings.ToLower(hostResponse)
			} else {
				log.Warnf("%s does not begin with 'https://'. This is highly insecure. "+
					"Continue anyways?", hostResponse)
				shouldContinue, err := util.AskUserForBool("Continue?", false)
				if err != nil {
					log.Errorf("Error getting your response.\nError:%s", err)
				} else {
					if shouldContinue {
						updatedApiUrl = strings.ToLower(hostResponse)
					}
				}
			}
		}
		// Add the information to the new list of domains to be updated
		newSimpleIPReporterServices = append(newSimpleIPReporterServices, config.SimpleIPReporterConfig{
			ServiceName: updatedSvcName,
			IPv4ApiBase: updatedApiUrl,
		})
		// We're done with this service, so move on.
	}
	return newSimpleIPReporterServices
}