package factory

import (
	"bookstore/store"
	"fmt"
	"sync"
)

var (
	providerMux sync.RWMutex
	providers   = make(map[string]store.Store)
)

func Register(name string, p store.Store) {
	providerMux.Lock()
	defer providerMux.Unlock()

	if p == nil {
		panic("store: Register provider is nil")
	}
	// 如果已经存在
	if _, dup := providers[name]; dup {
		panic("store: Register called twice for provider " + name)
	}
	providers[name] = p
}

func New(name string) (store.Store, error) {
	providerMux.RLock()
	p, ok := providers[name]
	providerMux.RUnlock()

	if !ok {
		return nil, fmt.Errorf("store: unknown provider %q", name)
	}

	return p, nil
}
