package cmd

import (
	"testing"
)

func TestLexString(t *testing.T) {
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
			name:    "no quotes returns error",
			args:    args{"abc"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "no closing quotes returns error",
			args:    args{"\"abc"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "no opening quotes returns error",
			args:    args{"abc\""},
			want:    "",
			wantErr: true,
		},
		{
			name:    "quotes present in string returns error",
			args:    args{"\"ab\"c\""},
			want:    "",
			wantErr: true,
		},
		{
			name: "valid string returns extracted content",
			args: args{
				s: "\"abc\"",
			},
			want:    "abc",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lex_string(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("lex_string() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("lex_string() = %v, want %v", got, tt.want)
			}
		})
	}
}
