package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLocalIp(t *testing.T) {
	lip := getLocalIp()
	assert.Equal(t, "10.0.0.103", lip.String())
}

func TestChallengeOne(t *testing.T) {
	challengeOne()
}
