package main

import (
	"context"
	"errors"
	"sync"
	"time"
)

const (
	KVStorageWrapperTypeInMemory = 1 << iota
)

type KVStorageWrapper struct {
	sync.RWMutex
	storageObject KVStorage
}

func NewKVStorageWrapper(storageType int) (*KVStorageWrapper, error) {
	var obj = new(KVStorageWrapper)
	switch storageType {
	case KVStorageWrapperTypeInMemory:
		obj.storageObject = NewKVStorageInMemory()
	default:
		return nil, errors.New("Wrong storage type")
	}
	return obj, nil
}

func (obj *KVStorageWrapper) Get(ctx context.Context, key string) (interface{}, error) {
	obj.RLock()
	defer obj.RUnlock()
	res, err := obj.storageObject.Get(ctx, key)
	return res, err
}

func (obj *KVStorageWrapper) Put(ctx context.Context, key string, val interface{}) error {
	obj.Lock()
	defer obj.Unlock()
	// TODO: experimental timeout to check sync
	time.Sleep(time.Second * 5)
	err := obj.storageObject.Put(ctx, key, val)
	return err
}

func (obj *KVStorageWrapper) Delete(ctx context.Context, key string) error {
	obj.Lock()
	defer obj.Unlock()
	err := obj.storageObject.Delete(ctx, key)
	return err
}
