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

import "os"

func CheckFileExists(path string) (bool, error) {
	// Attempt to get file stats on the file.
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// There was an error getting the file stats, but that error was that the file isn't there.
		return false, nil
	} else if err != nil {
		// There was some other error getting the file stats, but the file may be there.
		return false, err
	} else {
		// The file was there and we were able to get the stats on it
		return true, nil
	}
}
