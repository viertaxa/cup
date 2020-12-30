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
	"testing"
	"time"
)

func TestNewRetryTimer(t *testing.T) {
	timer := NewRetryTimer(3, 600*time.Second)

	if timer.attemptNo != 0 {
		t.Errorf("New timer instantiated with intial counter other than 0. Actual value: %d", timer.attemptNo)
	}
}

func TestGetRetryDelay(t *testing.T) {
	timer := NewRetryTimer(3, 600*time.Second)

	// RetryDelay should be 0 to begin with.
	if timer.GetRetryDelay() != time.Duration(0) {
		t.Errorf("Timer initial retry delay should be 0. Actual value in seconds: %d",
			timer.GetRetryDelay()/time.Second)
	}

	// with a factor of 3 and a max time of 600s, we should see retry delays of  1, 8, 27, 64, 125, 216, 343, 512, 600

	expected := map[int]time.Duration{
		0: 0 * time.Second,
		1: 1 * time.Second,
		2: 8 * time.Second,
		3: 27 * time.Second,
		4: 64 * time.Second,
		5: 125 * time.Second,
		6: 216 * time.Second,
		7: 343 * time.Second,
		8: 512 * time.Second,
		9: 600 * time.Second,
		10: 600 * time.Second,
		11: 600 * time.Second,
		12: 600 * time.Second,
		13: 600 * time.Second,
		14: 600 * time.Second,
	}

	for tr := 0; tr < 15; tr++ {
		if timer.GetRetryDelay() != expected[tr] {
			t.Errorf("Timer retry delay should be %d seconds for attempt %d. Actual value in seconds: %d",
				expected[tr]/time.Second, tr+1, timer.GetRetryDelay()/time.Second)
		}
		timer.AttemptFailed()
	}
}

func TestAttemptFailed(t *testing.T) {
	timer := NewRetryTimer(3, 600*time.Second)
	timer.AttemptFailed()
	if timer.attemptNo != 1 {
		t.Errorf("Timer incrimented once did not have counter of 1. Actual value: %d", timer.attemptNo)
	}
}