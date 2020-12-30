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

package externalip

import (
	"github.com/viertaxa/cup/internal/types"
	"testing"
)

func Test_getIPConsensus(t *testing.T) {
	type args struct {
		externalIps []types.ExternalIP
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Is Consensus",
			args:    args{
				externalIps: []types.ExternalIP{
					{
						IPv4: "1.1.1.1",
					},
					{
						IPv4: "1.1.1.1",
					},
					{
						IPv4: "2.2.2.2",
					},
				},
			},
			want:    "1.1.1.1",
			wantErr: false,
		},
		{
			name:    "Isn't Consensus",
			args:    args{
				externalIps: []types.ExternalIP{
					{
						IPv4: "1.1.1.1",
					},
					{
						IPv4: "3.3.3.3",
					},
					{
						IPv4: "2.2.2.2",
					},
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name:    "No values",
			args:    args{
				externalIps: []types.ExternalIP{},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getIPConsensus(tt.args.externalIps)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIPConsensus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getIPConsensus() got = %v, want %v", got, tt.want)
			}
		})
	}
}
