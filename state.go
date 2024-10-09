package tapper

import (
	"sync"
)

type _globalstate struct {
	mu   sync.Mutex
	data map[string]interface{}
}

func initGlobalState() *_globalstate {
	g := _globalstate{}
	g.data = make(map[string]interface{})
	return &g
}

func (g *_globalstate) Get(key string) interface{} {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.data[key]
}

func (g *_globalstate) Set(key string, value interface{}) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.data[key] = value
}

func (g *_globalstate) Clear() {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.data = make(map[string]interface{})
}

func (g *_globalstate) Delete(key string) {
	g.mu.Lock()
	defer g.mu.Unlock()
	delete(g.data, key)
}
