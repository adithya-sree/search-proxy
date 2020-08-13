package cache

import (
	"fmt"
	"search/src/app/service/deezer"
	"sync"
)

type InMemory struct {
	Cache map[string]*deezer.Response
	sync.Mutex
}

func NewCache() *InMemory {
	return &InMemory{
		Cache: make(map[string]*deezer.Response),
		Mutex: sync.Mutex{},
	}
}

func (im *InMemory) Find(query string) (*deezer.Response, error) {
	im.Lock()
	resp := im.Cache[query]
	im.Unlock()
	if resp != nil && resp.Data != nil {
		return resp, nil
	} else {
		return nil, fmt.Errorf("unable to find query [%s] in cache", query)
	}
}

func (im *InMemory) Add(query string, resp *deezer.Response) {
	im.Lock()
	im.Cache[query] = resp
	im.Unlock()
}

func (im *InMemory) Clear() {
	im.Lock()
	im.Cache = make(map[string]*deezer.Response)
	im.Unlock()
}