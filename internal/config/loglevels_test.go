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

package config

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogLevelToString(t *testing.T) {
	type args struct {
		level logrus.Level
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Trace",
			args: args{
				level: logrus.TraceLevel,
			},
			want: "trace",
		},
		{
			name: "Debug",
			args: args{
				level: logrus.DebugLevel,
			},
			want: "debug",
		},
		{
			name: "Info",
			args: args{
				level: logrus.InfoLevel,
			},
			want: "info",
		},
		{
			name: "Warn",
			args: args{
				level: logrus.WarnLevel,
			},
			want: "warn",
		},
		{
			name: "Error",
			args: args{
				level: logrus.ErrorLevel,
			},
			want: "error",
		},
		{
			name: "Fatal",
			args: args{
				level: logrus.FatalLevel,
			},
			want: "fatal",
		},
		{
			name: "Panic",
			args: args{
				level: logrus.PanicLevel,
			},
			want: "panic",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LogLevelToString(tt.args.level); got != tt.want {
				t.Errorf("LogLevelToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToLogLevel(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want logrus.Level
	}{
		{
			name: "Trace",
			args: args{
				str: "Trace",
			},
			want: logrus.TraceLevel,
		},
		{
			name: "trace",
			args: args{
				str: "trace",
			},
			want: logrus.TraceLevel,
		},
		{
			name: "Debug",
			args: args{
				str: "Debug",
			},
			want: logrus.DebugLevel,
		},
		{
			name: "debug",
			args: args{
				str: "debug",
			},
			want: logrus.DebugLevel,
		},
		{
			name: "Info",
			args: args{
				str: "Info",
			},
			want: logrus.InfoLevel,
		},
		{
			name: "info",
			args: args{
				str: "info",
			},
			want: logrus.InfoLevel,
		},
		{
			name: "Warn",
			args: args{
				str: "Warn",
			},
			want: logrus.WarnLevel,
		},
		{
			name: "warn",
			args: args{
				str: "warn",
			},
			want: logrus.WarnLevel,
		},
		{
			name: "Error",
			args: args{
				str: "Error",
			},
			want: logrus.ErrorLevel,
		},
		{
			name: "error",
			args: args{
				str: "error",
			},
			want: logrus.ErrorLevel,
		},
		{
			name: "Fatal",
			args: args{
				str: "Fatal",
			},
			want: logrus.FatalLevel,
		},
		{
			name: "fatal",
			args: args{
				str: "fatal",
			},
			want: logrus.FatalLevel,
		},
		{
			name: "Panic",
			args: args{
				str: "Panic",
			},
			want: logrus.PanicLevel,
		},
		{
			name: "panic",
			args: args{
				str: "panic",
			},
			want: logrus.PanicLevel,
		},
		{
			name: "Default",
			args: args{
				str: "foo",
			},
			want: logrus.InfoLevel,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToLogLevel(tt.args.str); got != tt.want {
				t.Errorf("StringToLogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
