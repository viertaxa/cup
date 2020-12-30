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
	"math"
	"time"
)

type RetryTimer struct {
	attemptNo int
	factor    int
	max       time.Duration
}

func NewRetryTimer(backoffFactor int, maxWaitTime time.Duration) RetryTimer {
	return RetryTimer{
		attemptNo: 0,
		factor:    backoffFactor,
		max:       maxWaitTime,
	}
}

func (r *RetryTimer) GetRetryDelay() time.Duration {
	if r.attemptNo == 0 {
		return time.Duration(0)
	}
	delay := time.Duration(int64(math.Pow(float64(r.attemptNo), float64(r.factor)))) * time.Second
	return GetSmallerDuration(r.max, delay)
}

func (r *RetryTimer) AttemptFailed() {
	r.attemptNo++
}
