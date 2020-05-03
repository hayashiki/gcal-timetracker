package calendar

import "crypto/rand"

func randomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func randomStringURLSafe(n int) (string, error) {
	bytes, err := randomBytes(n)

	if err != nil {
		return "", err
	}

	// noinspection ALL
	const symbols = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"

	for i, b := range bytes {
		bytes[i] = symbols[b%byte(len(symbols))]
	}

	return string(bytes), nil
}
