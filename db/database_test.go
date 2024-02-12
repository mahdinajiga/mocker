package db

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDatabase(t *testing.T) {
	mockDb := NewMockDatabase(t)
	mockDb.EXPECT().Connect().Return(nil)
	mockDb.EXPECT().Set("TestKey", "TestVal").Return(nil)
	mockDb.EXPECT().Get("TestKey").Return("TestVal", nil)
	mockDb.EXPECT().Get("testKey").Return("", fmt.Errorf("key not found in mock"))

	assert.Equal(t, nil, mockDb.Connect())
	assert.Equal(t, nil, mockDb.Set("TestKey", "TestVal"))

	val, err := mockDb.Get("TestKey")
	assert.Equal(t, "TestVal", val)
	assert.Equal(t, nil, err)

	val, err = mockDb.Get("testKey")
	assert.Equal(t, val, "")
	assert.Equal(t, err, fmt.Errorf("key not found in mock"))
}
