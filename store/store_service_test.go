package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.google.com"
	userUUId := "r0xad870-fd4c-4988-873f-r360239t6g1"
	shortURL := "Fk7szxT4gS"

	SaveUrlMapping(shortURL, initialLink, userUUId)
	retrievedUrl := RetrieveInitialUrl(shortURL)
	assert.Equal(t, initialLink, retrievedUrl)
}
