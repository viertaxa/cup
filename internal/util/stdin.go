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

func CheckIsStdin() (bool,error){
	// Grab stdin and check the properties of it
	if fileInfo, err := os.Stdin.Stat(); err != nil {
		// We weren't able to grab stdin for some reason.
		return false, err
	} else {
		// Since we were able to grab stdin, do the actual check.
		// Do a binary AND on the ModeCharDevice bit. If it's 0/false, return false.
		if (fileInfo.Mode() & os.ModeCharDevice) == 0 {
			return false, nil
		} else {
			return true, nil
		}
	}
}
