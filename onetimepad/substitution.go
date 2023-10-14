package onetimepad

import "math/rand"

type Substitution struct {
	// An alphabet is a set of character numbers in Unicode.
	alphabet []int

	// The key is a shuffled alphabet.
	key []int
}

func NewSubstitution(alphabet []int, seed int) Substitution {
	key := make([]int, len(alphabet))
	copy(key, alphabet)

	// The seed is used for pseudo-random shuffling to generate the key.
	gen := rand.New(rand.NewSource(int64(seed)))

	gen.Shuffle(len(key), func(i, j int) {
		key[i], key[j] = key[j], key[i]
	})

	return Substitution{
		alphabet: alphabet,
		key:      key,
	}
}

func NewSubstitutionWithKey(alphabet, key []int) Substitution {
	return Substitution{
		alphabet: alphabet,
		key:      key,
	}
}

func (s Substitution) Encrypt(plaintext []int) []int {
	return translate(plaintext, s.alphabet, s.key)
}

func (s Substitution) Decrypt(ciphertext []int) []int {
	return translate(ciphertext, s.key, s.alphabet)
}

// To decrypt, swap the alphabet and key at the inputs.
func translate(message, alphabet, key []int) []int {
	keyMap := make(map[int]int, len(key))
	cipherText := make([]int, len(message))

	for i, v := range key {
		keyMap[v] = i
	}

	for i, v := range message {
		j := keyMap[v]
		cipherText[i] = alphabet[j]
	}

	return cipherText
}
