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

func TestGetSmallerDuration(t *testing.T) {
	type args struct {
		d1 time.Duration
		d2 time.Duration
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "Different Times",
			args: args{d1: time.Second * 1, d2: time.Second * 2},
			want: time.Second * 1,
		},
		{
			name: "Same Times",
			args: args{d1: time.Second * 2, d2: time.Second * 2},
			want: time.Second * 2,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSmallerDuration(tt.args.d1, tt.args.d2); got != tt.want {
				t.Errorf("GetSmallerDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}