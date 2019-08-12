package cache

import (
	"container/list"
	"errors"
	"sync"
	"unsafe"
)

var (
	ErrorNil   = errors.New("nil")
	OutOfRange = errors.New("size is not enough")
)

type Cache interface {
	Set(key string, value interface{}) error
	Get(key string) (value interface{}, err error)
	Delete(key string) error
}

type LRUCache struct {
	cacheList *list.List
	cache     map[string]*list.Element
	size      int64
	used      int64
	mutex     sync.RWMutex
}

type cacheValue struct {
	key string
	value interface{}
}

func NewLRUCache(size int64) *LRUCache {
	return &LRUCache{
		size:      size,
		cacheList: list.New(),
		cache:     make(map[string]*list.Element),
	}
}

func (l *LRUCache) Set(key string, value interface{}) error {
	if l == nil {
		return ErrorNil
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()

	v := cacheValue{
		key:key,
		value:value,
	}
	size := int64(unsafe.Sizeof(v))
	if l.size < size {
		return OutOfRange
	}
	for l.size < l.used+size {
		back := l.cacheList.Back()
		l.used -= int64(unsafe.Sizeof(back.Value))
		l.cacheList.Remove(back)
		delete(l.cache, back.Value.(cacheValue).key)
	}
	if v, ok := l.cache[key]; ok {
		l.cacheList.MoveToFront(v)
		v.Value.(cacheValue).value = value
		return nil
	}
	l.cache[key] = l.cacheList.PushFront(v)
	l.used += size
	return nil
}

func (l *LRUCache) Get(key string) (interface{}, error) {
	if l == nil {
		return nil, ErrorNil
	}

	l.mutex.RLock()
	defer l.mutex.RUnlock()

	if v, ok := l.cache[key]; ok {
		if v == nil {
			return nil, ErrorNil
		}
		l.cacheList.MoveToFront(v)
		return v.Value.(cacheValue).value, nil
	}
	return nil, ErrorNil
}

func (l *LRUCache) Delete(key string) error {
	if l == nil {
		return ErrorNil
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()

	if v, ok := l.cache[key]; ok {
		l.cacheList.Remove(v)
		delete(l.cache, key)
		l.used -= int64(unsafe.Sizeof(v.Value))
	}
	return nil
}
