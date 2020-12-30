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

package externalip

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/viertaxa/cup/internal/config"
	"github.com/viertaxa/cup/internal/externalip/clients"
	"github.com/viertaxa/cup/internal/types"
	"github.com/viertaxa/cup/internal/util"
)

func GetExternalIP(cupInfo types.CupInfo) (*types.ExternalIP, error) {
	// Generate a user agent specific to this configuration, but identifiable as CUP
	ua := util.GenUserAgent(cupInfo, &config.C)

	var externalIPs []types.ExternalIP
	// Loop on provided Simple IP Reporting Services
	for _, svc := range config.C.SimpleIPReportServices {
		ip, err := clients.GetHttpIpAsBodySvcIP(svc.IPv4ApiBase, ua)
		if err != nil {
			log.Errorf("Error getting IP from %s.\nError: %s",svc.ServiceName, err)
		}
		if ip != nil {
			externalIPs = append(externalIPs, *ip)
			log.Debugf("Got public IP %s from ipify.", ip.IPv4)
		}
	}

	// Get what the consensus is for the public IP
	ipv4, err := getIPConsensus(externalIPs)
	if err != nil {
		log.Errorf("Error getting public IP\nError: %s", err)
		return nil, err
	}

	return &types.ExternalIP{IPv4: ipv4}, nil
}

// Get the consensus for what the public IP is
func getIPConsensus(externalIps []types.ExternalIP) (string, error) {
	var extractedIPv4Addresses []string

	// Pull the IPv4 addresses out.
	// Note: this is a bit odd for now because my IPv6 setup is broken, and I can't develop/test that part of the tool
	for _, externalIp := range externalIps {
		extractedIPv4Addresses = append(extractedIPv4Addresses, externalIp.IPv4)
	}

	ipConsensus := util.GetStringSliceMode(extractedIPv4Addresses)
	if len(ipConsensus) == 1 {
		log.Infof("It's agreed, your IP is %s", ipConsensus[0])
		if util.CountOccurrences(extractedIPv4Addresses, ipConsensus[0]) != len(extractedIPv4Addresses) {
			log.Warnf("Warning, not everyone agreed, though. only %d of %d thought that was the case.",
				util.CountOccurrences(extractedIPv4Addresses, ipConsensus[0]), len(extractedIPv4Addresses))
		}
		return ipConsensus[0], nil
	} else {
		return "", errors.New("could not come to a consensus on what the public IP address is")
	}
}
