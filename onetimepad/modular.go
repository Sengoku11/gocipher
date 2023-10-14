package onetimepad

// XOREncode can be used to both encode or decode a message with the same key.
// Constraint: len(key) >= len(message).
func XOREncode(key, message []byte) []byte {
	result := make([]byte, len(message))

	for i := range message {
		result[i] = message[i] ^ key[i]
	}

	return result
}
