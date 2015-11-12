package quiklyutil

import "crypto/rand"

func RandStrAlphaNum(strLen int) string {
	var bytes = make([]byte, strLen)
	dictionary := "0123456789abcdefghijklmnopqrstuvwxyz"

	rand.Read(bytes)

	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}