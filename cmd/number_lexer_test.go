package cmd

import (
	"testing"
)

func TestLexNumber(t *testing.T) {
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
			name:    "no digits returns error",
			args:    args{"abc"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "negative number returns extracted content",
			args:    args{"-123"},
			want:    "-123",
			wantErr: false,
		},
		{
			name:    "positive number returns extracted content",
			args:    args{"123"},
			want:    "123",
			wantErr: false,
		},
		{
			name:    "number with decimal returns extracted content",
			args:    args{"123.456"},
			want:    "123.456",
			wantErr: false,
		},
		{
			name:    "number with negative decimal returns extracted content",
			args:    args{"-123.456"},
			want:    "-123.456",
			wantErr: false,
		},
		{
			name:    "number with exponent returns extracted content",
			args:    args{"123e456"},
			want:    "123e456",
			wantErr: false,
		},
		{
			name:    "number with negative exponent returns extracted content",
			args:    args{"-123e456"},
			want:    "-123e456",
			wantErr: false,
		},
		{
			name:    "number with decimal and exponent returns extracted content",
			args:    args{"123.456e789"},
			want:    "123.456e789",
			wantErr: false,
		},
		{
			name:    "number with negative decimal and exponent returns extracted content",
			args:    args{"-123.456e789"},
			want:    "-123.456e789",
			wantErr: false,
		},
		{
			name:    "number with decimal and negative exponent returns extracted content",
			args:    args{"123.456e-789"},
			want:    "123.456e-789",
			wantErr: false,
		},
		{
			name:    "number with negative decimal and negative exponent returns extracted content",
			args:    args{"-123.456e-789"},
			want:    "-123.456e-789",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lex_number(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("lex_number() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("lex_number() = %v, want %v", got, tt.want)
			}
		})
	}
}
