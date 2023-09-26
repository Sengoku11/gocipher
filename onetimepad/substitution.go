package onetimepad

import "math/rand"

// An extended version of a Caesar cipher.
type Substitution struct {
	alphabet []int
	key      []int
}

// Alphabet is a set of chararacter numbers in Unicode.
// The seed is used for pseudo-random shuffling to generate the key.
func NewSubstitution(alphabet []int, seed int) Substitution {
	key := make([]int, len(alphabet))
	copy(key, alphabet)

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
	return translate(s.alphabet, s.key, plaintext)
}

func (s Substitution) Decrypt(ciphertext []int) []int {
	return translate(s.key, s.alphabet, ciphertext)
}

// To decrypt, swap the alphabet and key.
func translate(alphabet, key, message []int) []int {
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
