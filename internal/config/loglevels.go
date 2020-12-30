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

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

const traceConfigKey = "trace"
const debugConfigKey = "debug"
const infoConfigKey = "info"
const warnConfigKey = "warn"
const errorConfigKey = "error"
const fatalConfigKey = "fatal"
const panicConfigKey = "panic"

func GetValidLogLevels() []string {
	return []string{
		traceConfigKey,
		debugConfigKey,
		infoConfigKey,
		warnConfigKey,
		errorConfigKey,
		fatalConfigKey,
		panicConfigKey,
	}
}

func StringToLogLevel(str string) log.Level {
	lowerStr := strings.ToLower(str)
	switch lowerStr {
	case traceConfigKey:
		return log.TraceLevel
	case debugConfigKey:
		return log.DebugLevel
	case infoConfigKey:
		return log.InfoLevel
	case warnConfigKey:
		return log.WarnLevel
	case errorConfigKey:
		return log.ErrorLevel
	case fatalConfigKey:
		return log.FatalLevel
	case panicConfigKey:
		return log.PanicLevel
	default:
		// If we incorrectly check the input and end up here, return the program default of `info`
		return log.InfoLevel
	}
}

func LogLevelToString(level log.Level) string {
	switch level {
	case log.TraceLevel:
		return traceConfigKey
	case log.DebugLevel:
		return debugConfigKey
	case log.InfoLevel:
		return infoConfigKey
	case log.WarnLevel:
		return warnConfigKey
	case log.ErrorLevel:
		return errorConfigKey
	case log.FatalLevel:
		return fatalConfigKey
	case log.PanicLevel:
		return panicConfigKey
	default:
		// If we somehow get a different value, just return the application default log level, `info`
		return infoConfigKey
	}
}
