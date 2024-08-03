package services

import "encoding/base64"

func EncodeBase64(text string) (string, error) {
	encodeArr := make([]byte, base64.StdEncoding.EncodedLen(len(text)))
	base64.StdEncoding.Encode(encodeArr, []byte(text))

	return string(encodeArr), nil
}

func DecodeBase64(text string) (string, error) {
	encodeText, err := base64.StdEncoding.DecodeString(text)

	if err != nil {
		return "", err
	}

	return string(encodeText), nil
}
