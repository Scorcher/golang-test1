package main

import "context"

type KVStorageInMemory struct {
	data map[string]interface{}
}

func NewKVStorageInMemory() *KVStorageInMemory {
	obj := new(KVStorageInMemory)
	obj.data = make(map[string]interface{})
	return obj
}

func (obj *KVStorageInMemory) Get(ctx context.Context, key string) (interface{}, error) {
	return obj.data[key], nil
}

func (obj *KVStorageInMemory) Put(ctx context.Context, key string, val interface{}) error {
	obj.data[key] = val
	return nil
}

func (obj *KVStorageInMemory) Delete(ctx context.Context, key string) error {
	delete(obj.data, key)
	return nil
}
