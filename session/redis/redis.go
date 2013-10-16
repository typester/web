package redis

import (
	"bytes"
	"encoding/gob"
	redigo "github.com/garyburd/redigo/redis"
	"log"
	"time"
)

type RedisStore struct {
	addr, prefix string
}

func NewRedisStore(addr, prefix string) *RedisStore {
	return &RedisStore{addr, prefix}
}

func (store *RedisStore) connect() redigo.Conn {
	conn, err := redigo.Dial("tcp", store.addr)

	if err != nil {
		log.Panicf("Can't connect redis store: %v\n", err)
		return nil
	}

	return conn
}

func (store *RedisStore) GetData(sessionId string) interface{} {
	conn := store.connect()
	defer conn.Close()

	res, err := redigo.Bytes(conn.Do("GET", store.prefix+sessionId))

	if err != nil {
		if err == redigo.ErrNil {
			return nil
		}
		log.Panicf("Failed to get data from redis: %v\n", err)
	}

	var data map[string]interface{}
	dec := gob.NewDecoder(bytes.NewBuffer(res))
	if err = dec.Decode(&data); err != nil {
		log.Panicf("Failed to decode data from redis: %v\n", err)
	}

	return data
}

func (store *RedisStore) SetData(sessionId string, data map[string]interface{}, expires time.Duration) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(data); err != nil {
		log.Panicf("Failed to encode session data: %v\n", err)
	}

	conn := store.connect()
	defer conn.Close()

	_, err := conn.Do("SET", store.prefix+sessionId, buf.Bytes(), "EX", int64(expires/time.Second))
	if err != nil {
		log.Panicf("Failed to store session data: %v\n", err)
	}
}
