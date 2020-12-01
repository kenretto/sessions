package redis

import (
	"errors"
	sessredistore "github.com/boj/redistore"
	"github.com/go-redis/redis/v8"

	"github.com/gin-contrib/sessions"
)

type Store interface {
	sessions.Store
}

// size: maximum number of idle connections.
// network: tcp or udp
// address: host:port
// password: redis-password
// Keys are defined in pairs to allow key rotation, but the common case is to set a single
// authentication key and optionally an encryption key.
//
// The first key in a pair is used for authentication and the second for encryption. The
// encryption key can be set to nil or omitted in the last pair, but the authentication key
// is required in all pairs.
//
// It is recommended to use an authentication key with 32 or 64 bytes. The encryption key,
// if set, must be either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256 modes.
func NewStore(conn redis.Cmdable, keyPairs ...[]byte) Store {
	s := sessredistore.NewRedisStore(conn, keyPairs...)
	return &store{s}
}

type store struct {
	*sessredistore.RedisStore
}

// GetRedisStore get the actual woking store.
// Ref: https://godoc.org/github.com/boj/redistore#RediStore
func GetRedisStore(s Store) (err error, redisStore *sessredistore.RedisStore) {
	realStore, ok := s.(*store)
	if !ok {
		err = errors.New("unable to get the redis store: Store isn't *store")
		return
	}

	redisStore = realStore.RedisStore
	return
}

// SetKeyPrefix sets the key prefix in the redis database.
func SetKeyPrefix(s Store, prefix string) error {
	err, redisStore := GetRedisStore(s)
	if err != nil {
		return err
	}

	redisStore.SetKeyPrefix(prefix)
	return nil
}

func (c *store) Options(options sessions.Options) {
	c.RedisStore.Options = options.ToGorillaOptions()
}
