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
	"github.com/viertaxa/cup/internal/config"
	"github.com/viertaxa/cup/internal/util"
	"strings"
)

// Set up the log levels. Called by sub-commands
func setLogLevel() {
	// Check that the user passed in a valid log level and act accordingly.
	if config.C.LogLevel == "" {
		// If no log level was set on the command line, and there was no config file specifying this value, go with
		// something sane by default.
		log.SetLevel(log.InfoLevel)
		config.C.LogLevel = config.LogLevelToString(log.InfoLevel)
	} else if util.StringSliceContains(config.GetValidLogLevels(), strings.ToLower(config.C.LogLevel)) {
		// User's passed in log level was valid, so we use it.
		// Note: We avoid logging anything other than an Error or Fatal log message prior to this call
		// to avoid giving the user unwanted logs.
		log.SetLevel(config.StringToLogLevel(config.C.LogLevel))
		log.Infof("Setting log level to %s", config.C.LogLevel)
	} else {
		// User passed in an invalid log level identifier. We don't know how to log in this case!
		log.Fatalf("Error, you specified an invalid log level (%s)! Exiting.", config.C.LogLevel)
	}
}
