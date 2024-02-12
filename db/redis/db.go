package redis

import "fmt"

type Redis struct {
	Data map[string]string
}

func (m *Redis) Connect() error {
	m.Data = map[string]string{}
	return nil
}

func (m *Redis) Close() error {
	return nil
}

func (m *Redis) Set(key, value string) error {
	m.Data[key] = value
	return nil
}

func (m *Redis) Get(key string) (val string, err error) {
	if val, exists := m.Data[key]; exists {
		return val, nil
	} else {
		return "", fmt.Errorf("key not found in redis")
	}
}
