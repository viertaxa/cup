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
	"time"
)

const cfActiveTokenStatus = "active"

func CheckKeyIsActive(cf *cloudflare.API) (bool, error) {
	token, err := cf.VerifyAPIToken()
	if err != nil {
		return false, err
	}
	log.Debugf("API Token %s is %s, became valid on %s, and expires/expired on %s", token.ID, token.Status,
		token.NotBefore.Format(time.RFC1123), token.ExpiresOn.Format(time.RFC1123))

	if token.Status != cfActiveTokenStatus {
		return false, nil
	}
	return true, nil
}
