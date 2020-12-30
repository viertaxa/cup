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

//StringSliceContains returns `true` if the passed in string slice contains the
//passed in string, otherwise returns false.
func StringSliceContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func GetStringSliceMode(sl []string) []string {
	if len(sl) == 0 {
		return nil
	}
	valueMap := make(map[string]int)
	for _, s := range sl {
		valueMap[s]++
	}

	var maxFrequency int
	for _, v := range valueMap {
		if v > maxFrequency {
			maxFrequency = v
		}
	}

	var maxValues []string
	for k, v := range valueMap {
		if v == maxFrequency {
			maxValues = append(maxValues, k)
		}
	}
	return maxValues
}

func CountOccurrences(sl []string, s string) int {
	if len(sl) == 0 {
		return 0
	}

	var occurrences int
	for _, v := range sl {
		if v == s {
			occurrences++
		}
	}
	return occurrences
}
