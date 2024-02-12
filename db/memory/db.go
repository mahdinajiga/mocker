package memory

import "fmt"

type Memory struct {
	Data map[string]string
}

func (m *Memory) Connect() error {
	m.Data = map[string]string{}
	return nil
}

func (m *Memory) Close() error {
	return nil
}

func (m *Memory) Set(key, value string) error {
	m.Data[key] = value
	return nil
}

func (m *Memory) Get(key string) (val string, err error) {
	if val, exists := m.Data[key]; exists {
		return val, nil
	} else {
		return "", fmt.Errorf("key not found")
	}
}
