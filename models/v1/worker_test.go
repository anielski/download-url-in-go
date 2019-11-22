package v1

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//example test
func TestAddFetcher(t *testing.T) {
	AddFetcher( 1, "http://test.test", 1 )
	assert.Equal(t, len(FetchersWork), 1)
}

//example test
func TestDeleteFetcher(t *testing.T) {
	DeleteFetcherWork( 1 )
	assert.Equal(t, len(FetchersWork), 1)
}

//example test
func TestDeleteFetcherWork(t *testing.T) {
	time.Sleep( 1 * time.Second )
	assert.Equal(t, len(FetchersWork), 0)
}

//example test
func TestEditFetcher(t *testing.T) {
	EditFetcher(1, "http://test.test", 2 )
	assert.Equal(t, len(FetchersWork), 0)
}

//example test
func TestAddEditFetcher(t *testing.T) {
	AddFetcher( 1, "http://test.test", 1 )
	EditFetcher(1, "http://test.test", 2 )
	assert.Equal(t, len(FetchersWork), 1)

	assert.Equal(t, FetchersWork[0].Interval, uint(2) )
}