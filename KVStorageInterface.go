package main

import "context"

type KVStorage interface {
	Get(ctx context.Context, key string) (interface{}, error)

	Put(ctx context.Context, key string, val interface{}) error

	Delete(ctx context.Context, key string) error
}
