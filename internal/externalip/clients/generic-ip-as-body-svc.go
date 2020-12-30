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

package clients

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/viertaxa/cup/internal/types"
	"github.com/viertaxa/cup/internal/util"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func GetHttpIpAsBodySvcIP(url string, ua string) (*types.ExternalIP, error) {
	httpClient := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(userAgentHeaderKey, ua)

	//IPv4 request
	retryDelay := util.NewRetryTimer(3, 600*time.Second)
	for tries := 0; tries < maxRetries; tries++ {
		if tries != 0 {
			log.Warnf("Waiting %d before continuing after a failure.", int(retryDelay.GetRetryDelay().Seconds()))
		}
		time.Sleep(retryDelay.GetRetryDelay())
		resp, err := httpClient.Do(req)
		if err != nil {
			log.Errorf("Error encountered getting IP from service.\nError: %s", err)
			retryDelay.AttemptFailed()
			continue
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Error("Error decoding the response body.")
				retryDelay.AttemptFailed()
				continue
			}
			if resp.StatusCode != 200 {

				log.Errorf("Bad Response from the service.\n%d - %s", resp.StatusCode, body)
				retryDelay.AttemptFailed()
				continue
			} else {
				ip := net.ParseIP(string(body))
				if ip != nil {
					return &types.ExternalIP{IPv4: ip.String()}, nil
				} else {
					log.Errorf("Could not parse body as IP.\nBody: %s", body)
					retryDelay.AttemptFailed()
					continue
				}
			}
		}
	}
	return nil, errors.New("error getting ip from service, retries exceeded")
}
