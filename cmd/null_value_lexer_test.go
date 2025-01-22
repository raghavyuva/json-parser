package cmd

import (
	"testing"
)

func TestNullValueLexer(t *testing.T) {
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
			name:    "no null returns error",
			args:    args{"abc"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "null returns extracted content",
			args:    args{"null"},
			want:    "null",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lex_null_value(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("lex_null_value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("lex_null_value() = %v, want %v", got, tt.want)
			}
		})
	}
}
