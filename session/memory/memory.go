package memory

import (
	"sync"
	"time"
)

type MemoryStore struct {
	data map[string]interface{}
	mu   sync.Mutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{map[string]interface{}{}, sync.Mutex{}}
}

func (m *MemoryStore) GetData(sessionId string) interface{} {
	m.mu.Lock()
	defer m.mu.Unlock()

	data, ok := m.data[sessionId]
	if !ok {
		return nil
	}

	return data
}

func (m *MemoryStore) SetData(sessionId string, data map[string]interface{}, expires time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data[sessionId] = data
}
