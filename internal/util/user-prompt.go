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
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskUserForString(question string, defaultValue string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s [%s]:", question, defaultValue)

	response, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	if strings.TrimSpace(response) == "" {
		return defaultValue, nil
	}
	return strings.TrimSpace(response), nil
}

func AskUserForBool(question string, defaultValue bool) (bool, error) {
	var defaultWord string
	var trueResponses = []string{"y", "yes", "t", "true"}
	switch defaultValue {
	case true:
		defaultWord = "Y"
	case false:
		defaultWord = "N"
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s [%s]:", question, defaultWord)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	if strings.TrimSpace(response) == "" {
		return defaultValue, nil
	}
	if StringSliceContains(trueResponses, strings.ToLower(strings.TrimSpace(response))) {
		return true, nil
	}
	return false, nil
}
