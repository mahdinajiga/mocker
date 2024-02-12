package db

import (
	"mocker/db/memory"
	"mocker/db/redis"
)

func New(dbType string) Database {
	switch dbType {
	case "redis":
		return &redis.Redis{}
	case "memory":
		return &memory.Memory{}
	default:
		return &memory.Memory{}
	}
}
