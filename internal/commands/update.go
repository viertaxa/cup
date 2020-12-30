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
	"errors"
	"github.com/cloudflare/cloudflare-go"
	log "github.com/sirupsen/logrus"
	"github.com/viertaxa/cup/internal/cfutils"
	"github.com/viertaxa/cup/internal/config"
	"github.com/viertaxa/cup/internal/externalip"
	"github.com/viertaxa/cup/internal/types"
)

func Update(cupInfo types.CupInfo) error {
	// Get our external IP
	ip, err := externalip.GetExternalIP(cupInfo)
	if err != nil || ip == nil {
		log.Error("Error getting external IP.")
		return err
	}
	log.Debugf("Successfully got public IP: %s", ip.IPv4)

	// Create a CloudFlare API Client
	cf, err := cloudflare.NewWithAPIToken(config.C.API.ApiKey,
		func(api *cloudflare.API) error {
			api.BaseURL = config.C.API.ApiBaseUrl
			return nil
		})
	if err != nil {
		return err
	}

	// Check that the provided API Key is valid
	valid, err := cfutils.CheckKeyIsActive(cf)
	if err != nil {
		log.Error("Unable to check API key validity")
		return err
	}
	if !valid {
		log.Error("API Key is not active. Can't continue.")
		return errors.New("api key is not active")
	}

	// For each configured Domain, updated the specified hosts
	// We'll not bubble up and errors in an attempt to updates as many hosts as possible in the event of a failure
	for _, domain := range config.C.Domains {
		// Get the zone ID for the domain
		zoneID, err := cfutils.GetZoneID(domain, cf)
		if err != nil {
			log.Errorf("Error listing zone for domain `%s\nError:%s`", domain.Domain, err)
			continue
		}

		// Update each host record
		for _, host := range domain.Hosts {
			// Get the host ID for each host

			// The CF API has the whole domain as the name value for hosts. Generate that value.
			var apiNameValue string
			if host == "@" {
				apiNameValue = domain.Domain
			} else {
				apiNameValue = host + "." + domain.Domain
			}
			log.Debugf("API Name Value for %s is %s", host, apiNameValue)

			// Generate a DNSRecord for us to use as filter parameters
			filterRecord := cloudflare.DNSRecord{
				// Get only A records
				Type: "A",
				Name: apiNameValue,
			}

			// Request the records matching the host
			// Note that multiple are able to be returned if the user has defined multiple IPs (i.e. for round robin)
			// within Cloudflare.
			records, err := cf.DNSRecords(zoneID, filterRecord)
			if err != nil {
				log.Errorf("Error getting DNS records matching %s\nError: %s", host, err)
				continue
			}

			// Make sure we only got back one record.
			if len(records) != 1 {
				log.Errorf("There were %d records found for %s. We only support hosts with a single record in "+
					"this tool.", len(records), host)
				continue
			}
			record := records[0]
			log.Debugf("Got host record with ID %s for host %s, current value %s",
				record.ID, record.Name, record.Content)

			// If the value that's currently in Cloudflare does not match the public IP that we determined, update it
			if record.Content != ip.IPv4 {
				// Generate a partial record to update with
				newRecord := cloudflare.DNSRecord{
					Content: ip.IPv4,
				}

				// Send the record to Cloudflare
				err = cf.UpdateDNSRecord(zoneID, record.ID, newRecord)
				if err != nil {
					log.Errorf("Error updating IP for host\nError: %s", err)
					continue
				} else {
					log.Infof("Succsfully updated record for host %s with ip %s", host, ip.IPv4)
				}
			} else {
				// Don't need to update, won't try to not be abusive to the API
				log.Info("DNS Record already matches the public IP. Not updating.")
			}
		}
	}
	return nil
}
