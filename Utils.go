package main

import (
	"errors"
	"strings"
)

func getKey(path string) (string, error) {
	pathCunks := strings.Split(path, `/`)
	if len(pathCunks) < 4 {
		return "", errors.New(`wrong path`)
	}
	return pathCunks[3], nil
}

func getKeyValue(path string) (string, string, error) {
	pathCunks := strings.Split(path, `/`)
	if len(pathCunks) < 5 {
		return "", "", errors.New(`wrong path`)
	}
	return pathCunks[3], pathCunks[4], nil
}
