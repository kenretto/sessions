# sessions

[![Build Status](https://travis-ci.org/gin-contrib/sessions.svg)](https://travis-ci.org/gin-contrib/sessions)
[![codecov](https://codecov.io/gh/kenretto/sessions/branch/master/graph/badge.svg)](https://codecov.io/gh/kenretto/sessions)
[![Go Report Card](https://goreportcard.com/badge/github.com/kenretto/sessions)](https://goreportcard.com/report/github.com/kenretto/sessions)
[![GoDoc](https://godoc.org/github.com/kenretto/sessions?status.svg)](https://godoc.org/github.com/kenretto/sessions)
[![Join the chat at https://gitter.im/gin-gonic/gin](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

Gin middleware for session management with multi-backend support:

- [cookie-based](#cookie-based)
- [Redis](#redis)
- [memcached](#memcached)
- [MongoDB](#mongodb)
- [memstore](#memstore)

## Usage

### Start using it

Download and install it:

```bash
$ go get github.com/kenretto/sessions
```

Import it in your code:

```go
import "github.com/kenretto/sessions"
```

## Basic Examples

### single session

```go
package main

import (
	"github.com/kenretto/sessions"
	"github.com/kenretto/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/hello", func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			session.Save()
		}

		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})
	r.Run(":8000")
}
```

### multiple sessions

```go
package main

import (
	"github.com/kenretto/sessions"
	"github.com/kenretto/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	sessionNames := []string{"a", "b"}
	r.Use(sessions.SessionsMany(sessionNames, store))

	r.GET("/hello", func(c *gin.Context) {
		sessionA := sessions.DefaultMany(c, "a")
		sessionB := sessions.DefaultMany(c, "b")

		if sessionA.Get("hello") != "world!" {
			sessionA.Set("hello", "world!")
			sessionA.Save()
		}

		if sessionB.Get("hello") != "world?" {
			sessionB.Set("hello", "world?")
			sessionB.Save()
		}

		c.JSON(200, gin.H{
			"a": sessionA.Get("hello"),
			"b": sessionB.Get("hello"),
		})
	})
	r.Run(":8000")
}
```

## Backend Examples

### cookie-based

[embedmd]:# (_example/cookie/main.go go)
```go
package main

import (
	"github.com/kenretto/sessions"
	"github.com/kenretto/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	r.Run(":8000")
}
```

### Redis

[embedmd]:# (_example/redis/main.go go)
```go
package main

import (
	"github.com/gin-gonic/gin"
	redisDriver "github.com/go-redis/redis/v8"
	"github.com/kenretto/sessions"
	"github.com/kenretto/sessions/redis"
)

func main() {
	r := gin.Default()
	storeDB := redisDriver.NewClient(&redisDriver.Options{
		Addr: "localhost:6379",
	})
	store := redis.NewStore(storeDB, []byte("secret"))	
    r.Use(sessions.Sessions("mysession", store))

	r.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	r.Run(":8000")
}
```

### memstore

[embedmd]:# (_example/memstore/main.go go)
```go
package main

import (
	"github.com/kenretto/sessions"
	"github.com/kenretto/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	r.Run(":8000")
}
```
