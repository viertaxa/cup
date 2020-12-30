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

package config

var C CupConfig
var FilePath string

const DefaultCupConfigPath = "./cupconf.yaml"

type (
	CupConfig struct {
		LogLevel               string                   `json:"logLevel"`
		API                    ApiConfig                `json:"api"`
		Domains                []DomainConfig           `json:"domains"`
		SimpleIPReportServices []SimpleIPReporterConfig `json:"simpleIpReportServices"`
	}
	ApiConfig struct {
		ApiBaseUrl string `json:"apiBaseUrl"`
		ApiKey     string `json:"apiKey"`
	}
	DomainConfig struct {
		Domain string   `json:"domain"`
		Hosts  []string `json:"hosts"`
	}
	SimpleIPReporterConfig struct {
		IPv4ApiBase string `json:"ipv4ApiBase"`
		ServiceName string `json:"serviceName"`
	}
)

func GetDefaultCupConfig() CupConfig {
	return CupConfig{
		LogLevel: "info",
		API: ApiConfig{
			ApiBaseUrl: "https://api.cloudflare.com/client/v4",
			ApiKey:     "1234567893feefc5f0q5000bfo0c38d90bbeb",
		},
		Domains: []DomainConfig{
			{
				Domain: "example.com",
				Hosts:  []string{"dynamic1", "subdynamic.host", "@"},
			},
		},
		SimpleIPReportServices: []SimpleIPReporterConfig{
			{
				IPv4ApiBase: "https://api.ipify.org",
				ServiceName: "ipify",
			},
			{
				IPv4ApiBase: "https://ip4.seeip.org",
				ServiceName: "SeeIP",
			},
			{
				IPv4ApiBase: "https://ipv4bot.whatismyipaddress.com",
				ServiceName: "WhatIsMyIPAddress",
			},
		},
	}
}
