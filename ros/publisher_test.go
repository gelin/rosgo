package ros

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidMessageType(t *testing.T) {
	session := remoteSubscriberSession{
		typeName: "message/Test",
		md5sum:   "123456",
	}

	assert.True(t, session.isValidMessageType(map[string]string{"md5sum": "123456", "type": "message/Test"}))
	assert.True(t, session.isValidMessageType(map[string]string{"md5sum": "123456"}))
	assert.True(t, session.isValidMessageType(map[string]string{"md5sum": "*"}))
	assert.True(t, session.isValidMessageType(map[string]string{"md5sum": "*", "type": "message/Test"}))
	assert.True(t, session.isValidMessageType(map[string]string{"md5sum": "123456", "type": "*"}))
	assert.True(t, session.isValidMessageType(map[string]string{"md5sum": "*", "type": "*"}))

	assert.False(t, session.isValidMessageType(map[string]string{}))
	assert.False(t, session.isValidMessageType(map[string]string{"type": "message/Test"}))
	assert.False(t, session.isValidMessageType(map[string]string{"md5sum": "123456", "type": "message/Test2"}))
	assert.False(t, session.isValidMessageType(map[string]string{"md5sum": "12345678", "type": "message/Test"}))
	assert.False(t, session.isValidMessageType(map[string]string{"md5sum": "12345678", "type": "message/Test2"}))
	assert.False(t, session.isValidMessageType(map[string]string{"md5sum": "12345678"}))
	assert.False(t, session.isValidMessageType(map[string]string{"md5sum": "12345678", "type": "*"}))
	assert.False(t, session.isValidMessageType(map[string]string{"md5sum": "*", "type": "message/Test2"}))
}
