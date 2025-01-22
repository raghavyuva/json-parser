package cmd

import (
	"reflect"
	"testing"
)

func TestLexer(t *testing.T) {
	type args struct {
		s string
	}
 
	tests := []struct {
		name    string
		args    args 
		want    []string
		wantErr bool
	}{
		{
			name: "empty string",
			args: args{s: ""},
			want: []string{},
			wantErr: false,
		},
		{
			name: "simple object",
			args: args{s: `{"foo": "bar"}`},
			want: []string{"{", "foo", ":", "bar", "}"},
			wantErr: false,
		},
		{
			name: "nested object",
			args: args{s: `{"foo": {"bar": 42}}`},
			want: []string{"{", "foo", ":", "{", "bar", ":", "42", "}", "}"},
			wantErr: false,
		},
		{
			name: "array with mixed types",
			args: args{s: `[1, "two", true, null]`},
			want: []string{"[", "1", ",", "two", ",", "true", ",", "null", "]"},
			wantErr: false,
		},
		{
			name: "unterminated string",
			args: args{s: `{"foo": "bar`},
			want: nil,
			wantErr: true,
		},
		{
			name: "whitespace handling",
			args: args{s: ` { "foo" : 42 } `},
			want: []string{"{", "foo", ":", "42", "}"},
			wantErr: false,
		},
	}
 
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lexer(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("lexer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexer() = %v, want %v", got, tt.want)
			}
		})
	}
 }