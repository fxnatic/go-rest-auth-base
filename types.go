package main

import (
	"sync"
	"time"
)

type ApiKey struct {
	Limit    int64         // Maximum number of requests that can be made with this key.
	Duration time.Duration // Time until the usage counter should be reset.
	Usage    int64         // Current number of requests made with this key.
	Reset    int64         // Timestamp when the usage counter should be reset.
}

type APIKeyMap struct {
	mu sync.RWMutex       // Mutex to allow multiple readers or one writer.
	m  map[string]*ApiKey // Map from string (the API key) to an ApiKey object.
}

func (sm *APIKeyMap) Get(key string) (*ApiKey, bool) {
	sm.mu.RLock()         // Lock for reading.
	defer sm.mu.RUnlock() // Unlock after this function finishes.
	val, ok := sm.m[key]  // Retrieve the ApiKey object.
	return val, ok        // Return the ApiKey object and whether it was found.
}

func (sm *APIKeyMap) Set(key string, value *ApiKey) {
	sm.mu.Lock()         // Lock for writing.
	defer sm.mu.Unlock() // Unlock after this function finishes.
	sm.m[key] = value    // Set the key to map to the ApiKey object.
}
