package integrations

import (
	"LineMongo/internals/implements"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMongoDBConnected(t *testing.T) {
	assert.Equal(t, true, implements.MongoDBManager.IsConnected())
}

func TestLineBotConnected(t *testing.T) {
	assert.Equal(t, true, implements.LineBotManager.IsConnected())
}
