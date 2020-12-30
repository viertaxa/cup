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

package cfutils

import (
	"github.com/cloudflare/cloudflare-go"
	log "github.com/sirupsen/logrus"
	"github.com/viertaxa/cup/internal/config"
	"gopkg.in/errgo.v2/fmt/errors"
)

func GetZoneID(domain config.DomainConfig, cf *cloudflare.API) (string, error) {
	// Get the zone object so we can get it's ID
	zones, err := cf.ListZones(domain.Domain)
	if err != nil {
		return "", err
	}
	log.Debugf("Got %d zones matching %s", len(zones), domain.Domain)
	if len(zones) != 1 {
		return "", errors.Newf("got %d zones matching %s. Please file a bug report on GitHub if you believe "+
			"this to be incorrect", len(zones), domain.Domain)
	}
	zone := zones[0]
	log.Debugf("Got zone ID %s for domain %s", zone.ID, zone.Name)
	return zone.ID, nil
}