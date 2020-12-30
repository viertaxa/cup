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
	"strings"
	"testing"
)

func TestPrettyPrintSYaml(t *testing.T) {
	type args struct {
		t interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Print a sample config",
			args:    args{
				t: config.CupConfig{
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
			want:    `
loglevel: trace
api:
  apibaseurl: https://foo.bar/
  apikey: dfsadsdsafdfsdfa
domains:
- domain: foo.com
  hosts:
  - bar
simpleipreportservices:
- ipv4apibase: https://ip.foo.bar
  servicename: FooBarIPReporter`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrettyPrintSYaml(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrettyPrintSYaml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if strings.TrimSpace(got) != strings.TrimSpace(tt.want) {
				t.Errorf("PrettyPrintSYaml() got = %v, want %v", got, tt.want)
			}
		})
	}
}
