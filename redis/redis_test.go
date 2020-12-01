package redis

import (
	"github.com/go-redis/redis/v8"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/tester"
)

const redisTestServer = "localhost:6379"

var newRedisStore = func(_ *testing.T) sessions.Store {
	s := redis.NewClient(&redis.Options{
		Addr: redisTestServer,
	})
	store := NewStore(s, []byte("secret"))
	return store
}

func TestRedis_SessionGetSet(t *testing.T) {
	tester.GetSet(t, newRedisStore)
}

func TestRedis_SessionDeleteKey(t *testing.T) {
	tester.DeleteKey(t, newRedisStore)
}

func TestRedis_SessionFlashes(t *testing.T) {
	tester.Flashes(t, newRedisStore)
}

func TestRedis_SessionClear(t *testing.T) {
	tester.Clear(t, newRedisStore)
}

func TestRedis_SessionOptions(t *testing.T) {
	tester.Options(t, newRedisStore)
}

func TestGetRedisStore(t *testing.T) {
	t.Run("unmatched type", func(t *testing.T) {
		type store struct{ Store }
		err, redisStore := GetRedisStore(store{})
		if err == nil || redisStore != nil {
			t.Fail()
		}
	})
}
