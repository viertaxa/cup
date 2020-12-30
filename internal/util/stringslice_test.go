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
	"reflect"
	"sort"
	"testing"
)

func TestCountOccurrences(t *testing.T) {
	type args struct {
		sl []string
		s  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "No Occurrences",
			args: args{
				sl: []string{"bar", "baz"},
				s: "foo",
			},
			want: 0,
		},
		{
			name: "1 Occurrence",
			args: args{
				sl: []string{"bar", "baz", "foo"},
				s: "foo",
			},
			want: 1,
		},
		{
			name: "3 Occurrences",
			args: args{
				sl: []string{"foo", "bar", "foo", "baz", "foo"},
				s: "foo",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountOccurrences(tt.args.sl, tt.args.s); got != tt.want {
				t.Errorf("CountOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStringSliceMode(t *testing.T) {
	type args struct {
		sl []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "All the same value",
			args: args{
				sl: []string{"foo", "foo", "foo"},
			},
			want: []string{"foo"},
		},
		{
			name: "Two different values with winner",
			args: args{
				sl: []string{"foo", "foo", "bar"},
			},
			want: []string{"foo"},
		},
		{
			name: "All different values",
			args: args{
				sl: []string{"foo", "bar", "baz"},
			},
			want: []string{"foo", "bar", "baz"},
		},
		{
			name: "Two winners with one loser",
			args: args{
				sl: []string{"foo", "foo", "bar", "bar", "baz"},
			},
			want: []string{"foo", "bar"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetStringSliceMode(tt.args.sl)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStringSliceMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSliceContains(t *testing.T) {
	type args struct {
		s   []string
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Does contain",
			args: args{
				s:   []string{"foo", "bar", "baz"},
				str: "bar",
			},
			want: true,
		},
		{
			name: "Does not contain",
			args: args{
				s:   []string{"foo", "bar", "baz"},
				str: "bang",
			},
			want: false,
		},
		{
			name: "Does not contain (wrong case)",
			args: args{
				s:   []string{"foo", "Bar", "baz"},
				str: "bar",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringSliceContains(tt.args.s, tt.args.str); got != tt.want {
				t.Errorf("StringSliceContains() = %v, want %v", got, tt.want)
			}
		})
	}
}
