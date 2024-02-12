package main

import (
	"mocker/db"
	"testing"
)

func TestRun(t *testing.T) {
	mockDB := db.NewMockDatabase(t)
	mockDB.EXPECT().Set("MyKey", "A value").Return(nil)
	mockDB.EXPECT().Get("MyKey").Return("A value", nil)
	//mockDB.EXPECT().Get("").Return("", fmt.Errorf("key cannot be empty"))

	Run(mockDB)
}
