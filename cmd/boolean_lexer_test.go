package cmd

import (
	"testing"
)

func TestBooleanLexer(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "empty string returns error",
			args:    args{""},
			want:    "",
			wantErr: true,
		},
		{
			name:    "no boolean returns error",
			args:    args{"abc"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "true returns extracted content",
			args:    args{"true"},
			want:    "true",
			wantErr: false,
		},
		{
			name:    "false returns extracted content",
			args:    args{"false"},
			want:    "false",
			wantErr: false,
		},
		{
			name:    "invalid boolean returns error",
			args:    args{"false2"},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lex_boolean(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("lex_boolean() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("lex_boolean() = %v, want %v", got, tt.want)
			}
		})
	}
}
