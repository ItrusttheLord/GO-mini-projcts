package storage

import (
	"errors"
	"sync"
)

type Storage struct {
	Map map[string]string
	mu  sync.Mutex
}

// init map
func NewStorage() *Storage {
	st := &Storage{
		Map: make(map[string]string), //create empty map
	}
	return st //return pointer
}

func (S *Storage) AddURL(short string, original string) error {
	if short == "" || original == "" {
		return errors.New("enter a valid url")
	}

	S.mu.Lock()         //lock so no other go rutines can write/read
	defer S.mu.Unlock() //check if short exist in map
	if _, ok := S.Map[short]; ok {
		return errors.New("short URL already exist")
	}
	S.Map[short] = original //assign value of original to the short url
	return nil
}

func (S *Storage) GetURL(short string) (string, bool) {
	S.mu.Lock()
	defer S.mu.Unlock() //check if the short exist
	if original, ok := S.Map[short]; ok {
		return original, true //return the original
	}
	return "", false
}

func (S *Storage) RemoveURL(short string) error {
	if short == "" {
		return errors.New("please enter a valid url")
	}
	S.mu.Lock()
	defer S.mu.Unlock()
	if _, ok := S.Map[short]; !ok {
		return errors.New("the url you have enter doesn't exist")
	}
	delete(S.Map, short) //remove the url
	return nil
}

// get all the urls
func (S *Storage) ListAllURLs() map[string]string {
	S.mu.Lock()
	defer S.mu.Unlock()
	sl := make(map[string]string)

	for i, value := range S.Map {
		sl[i] = value //populate sl with short=keys and original=values url
	}
	return sl
}
