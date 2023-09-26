package onetimepad_test

import (
	"testing"

	"github.com/Sengoku11/gocipher/onetimepad"
)

func TestNewSubstitution(t *testing.T) {
	var alphabet1 []int
	for i := 20; i <= 126; i++ {
		alphabet1 = append(alphabet1, i)
	}

	var alphabet2 []int
	for i := 20; i <= 1260; i++ {
		alphabet2 = append(alphabet2, i)
	}

	tests := []struct {
		name     string
		alphabet []int
		key      []int
		seed     int
		message  string
	}{
		{
			name:     "latin, numbers and common symbols",
			alphabet: alphabet1,
			seed:     123,
			message:  "Hello World!",
		},
		{
			name:     "wide alphabet",
			alphabet: alphabet2,
			seed:     123,
			message:  "Hello World!",
		},
		{
			name:     "different seed number",
			alphabet: alphabet1,
			seed:     12345,
			message:  "Hello World!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			subEnc := onetimepad.NewSubstitution(tt.alphabet, tt.seed)

			msgInt := make([]int, len(tt.message))
			for i, char := range tt.message {
				msgInt[i] = int(char)
			}

			ciphertextInt := subEnc.Encrypt(msgInt)
			ciphertext := ""
			for _, v := range ciphertextInt {
				ciphertext += string(rune(v))
			}

			plaintextInt := subEnc.Decrypt(ciphertextInt)
			plaintext := ""
			for _, v := range plaintextInt {
				plaintext += string(rune(v))
			}

			if tt.message != plaintext {
				t.Errorf("Decoded ciphertext (%s) doesnt match original message (%s).", plaintext, tt.message)
			}

			if ciphertext == plaintext {
				t.Errorf("Ciphertext (%v) shouldn't be equal to plaintext (%v).", ciphertext, plaintext)
			}
		})
	}
}
