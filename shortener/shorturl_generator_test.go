package shortener

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const UserId = "f1dbc740-th4x-8544-p89v-c24y852f2e1a"

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://en.wikipedia.org/wiki/Wikipedia"
	shortLink_1 := GenerateShortLink(initialLink_1, UserId)

	initialLink_2 := "https://redis.io/docs/latest/get-started/"
	shortLink_2 := GenerateShortLink(initialLink_2, UserId)

	initialLink_3 := "https://pokemonletsgo.pokemon.com/en-us/"
	shortLink_3 := GenerateShortLink(initialLink_3, UserId)

	assert.Equal(t, shortLink_1, "B7Co8cFG")
	assert.Equal(t, shortLink_2, "Avaxt3pj")
	assert.Equal(t, shortLink_3, "AC58s3Ck")
}
