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

package util

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/viertaxa/cup/internal/config"
	"github.com/viertaxa/cup/internal/types"
)

const ConfigBytesStandin = "I am but a poor cup executable that was unable to marshal my config."

func GenUserAgent(cupInfo types.CupInfo, cupConfig *config.CupConfig) string {
	hasher := sha256.New()
	configBytes, err := json.Marshal(cupConfig)
	if err != nil {
		hasher.Write([]byte(ConfigBytesStandin))
	} else {
		hasher.Write(configBytes)
	}
	shaHash := hasher.Sum(nil)

	return fmt.Sprintf("cup/%s-%s (%s) instance/%x",
		cupInfo.Version,
		cupInfo.Commit,
		cupInfo.BuildType,
		shaHash)
}
