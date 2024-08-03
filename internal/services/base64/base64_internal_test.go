package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct{
	realText string
	encodeText string
}{
	{realText: "", encodeText: ""},
	{realText: "And", encodeText: "QW5k"},
	{realText: "An", encodeText: "QW4="},
	{realText: "Привет", encodeText: "0J/RgNC40LLQtdGC"},
	{realText: "😒", encodeText: "8J+Ykg=="},
	{realText: "🩬", encodeText: "8J+prA=="},
}

func TestBase64_Encoder(t *testing.T) {
	for _, cs := range testCases {
		text, _ := EncodeBase64(cs.realText)
		assert.Equal(t, text, cs.encodeText)
	}
}

func TestBase64_Decoder(t *testing.T) {
	for _, cs := range testCases {
		text, err := DecodeBase64(cs.encodeText)
		assert.NoError(t, err)
		assert.Equal(t, text, cs.realText)
	}
}
