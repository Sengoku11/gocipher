package onetimepad_test

import (
	"reflect"
	"testing"

	"github.com/Sengoku11/gocipher/onetimepad"
)

func TestXOREncode(t *testing.T) {
	type args struct {
		key     []byte
		message []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "equal length",
			args: args{
				key:     []byte("123456789ab"),
				message: []byte("Hello World"),
			},
		},
		{
			name: "len(key) >= len(message)",
			args: args{
				key:     []byte("SAn91jsad&*SDhklAjsml!2"),
				message: []byte("Hello World"),
			},
		},
		{
			name: "symbols (uses more bytes)",
			args: args{
				key:     []byte("ййй"),
				message: []byte("Hello"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ciphertext := onetimepad.XOREncode(tt.args.key, tt.args.message)
			plaintext := onetimepad.XOREncode(tt.args.key, ciphertext)

			// Check if the decoded message matches what it was before encoding.
			if !reflect.DeepEqual(plaintext, tt.args.message) {
				t.Errorf("XOREncode() = %v, want %v", plaintext, tt.args.message)
			}

			// Check if encoding was really happened and plaintext != ciphertext.
			if reflect.DeepEqual(plaintext, ciphertext) {
				t.Errorf("Ciphertext (%v) shouldn't be equal to plaintext (%v).", ciphertext, plaintext)
			}
		})
	}
}
