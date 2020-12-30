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
	"github.com/viertaxa/cup/internal/config"
	"github.com/viertaxa/cup/internal/types"
	"testing"
)

func TestGenUserAgent(t *testing.T) {
	type args struct {
		cupInfo   types.CupInfo
		cupConfig *config.CupConfig
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Unable to marshal config",
			args: args{
				cupInfo:   types.CupInfo{
					Version:   "1.0.0",
					Commit:    "deadbeef",
					BuildType: "testing",
				},
				cupConfig: nil,
			},
			want: "cup/1.0.0-deadbeef (testing) instance/74234e98afe7498fb5daf1f36ac2d78acc339464f950703b8c019892f982b90b",
		},
		{
			name: "Valid Config",
			args: args{
				cupInfo:   types.CupInfo{
					Version:   "1.0.0",
					Commit:    "deadbeef",
					BuildType: "testing",
				},
				cupConfig: &config.CupConfig{
					LogLevel:               "trace",
					API:                    config.ApiConfig{
						ApiBaseUrl: "https://foo.bar/",
						ApiKey:     "dfsadsdsafdfsdfa",
					},
					Domains:                []config.DomainConfig{
						{
							Domain: "foo.com",
							Hosts:  []string{
								"bar",
							},
						},
					},
					SimpleIPReportServices: []config.SimpleIPReporterConfig{
						{
							IPv4ApiBase: "https://ip.foo.bar",
							ServiceName: "FooBarIPReporter",
						},
					},
				},
			},
			want: "cup/1.0.0-deadbeef (testing) instance/7b80fdc9cb0a664c2333cc41bb8875c5821fe1d7b4257a26c13661c363368f99",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenUserAgent(tt.args.cupInfo, tt.args.cupConfig); got != tt.want {
				t.Errorf("GenUserAgent() = %v, want %v", got, tt.want)
			}
		})
	}
}