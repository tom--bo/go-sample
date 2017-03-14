package main

import (
  "fmt"
  "time"
  "strconv"
  "sync"
  "github.com/garyburd/redigo/redis"
)
// リファレンス
// http://godoc.org/github.com/garyburd/redigo/redis
// https://godoc.org/github.com/hjr265/redsync.go/redsync#Mutex

var pool *redis.Pool
var LockKeyPrefix = "lock:"

type RedisLock struct {
  lockInterface
  isLocked bool
  lockedValue string
  lockKey string
}

type lockInterface interface {
  lock()
  unlock()
}

func (r *RedisLock) lock() {
  if r.isLocked {
    return
  }
  lockKey := LockKeyPrefix + r.lockKey
  max := (time.Now().UnixNano() / int64(time.Millisecond)) + 30000
  redis_c := pool.Get()
  defer redis_c.Close()
  for ;; {
    now :=  time.Now().UnixNano() / int64(time.Millisecond)
    value := strconv.FormatInt(now, 10)
    result, _ := redis.String(redis_c.Do("SET", lockKey, value, "NX"))
    if result == "OK"{
      redis_c.Do("EXPIRE", lockKey, 30)
      r.isLocked = true
      r.lockedValue = value
      break
    }
    if max <= now {
      locked, err := redis.Int64(redis_c.Do("GET", lockKey))
      if err != nil {
        elapsedTime := now - locked
        if elapsedTime >= 30*1000 {
          redis_c.Do("DEL", lockKey)
        }
      }
    }
  }
}

func (r *RedisLock) unlock(){
  if !r.isLocked {
    return
  }
  lockKey := LockKeyPrefix + r.lockKey
  redis_c := pool.Get()
  defer redis_c.Close()
  value, _ := redis.String(redis_c.Do("GET", lockKey))
  if value == r.lockedValue {
    redis_c.Do("DEL", lockKey)
  }
  r.isLocked = false
  r.lockedValue = "";
}

func newPool() *redis.Pool {
  return &redis.Pool{
    MaxIdle:     1000,
    IdleTimeout: 240 * time.Second,
    Dial: func() (redis.Conn, error) {
      //c, err := redis.Dial("unix", "/tmp/redis.sock")
      c, err := redis.Dial("tcp", "127.0.0.1:6379")
      if err != nil {
        return nil, err
      }
      return c, err
    },
    TestOnBorrow: func(c redis.Conn, t time.Time) error {
      if time.Since(t) < time.Minute {
        return nil
      }
      _, err := c.Do("PING")
      return err
    },
  }
}

func update() {
  key := "test"
  var lock RedisLock
  lock.isLocked = false
  lock.lockedValue = ""
  lock.lockKey = key

  redis_c := pool.Get()
  defer redis_c.Close()
  lock.lock()
  defer lock.unlock()

  // update処理
  cnt, err := redis.Int(redis_c.Do("GET", key))
  if err != nil {
    fmt.Println(err)
  }
  cnt++
  time.Sleep(10 * time.Millisecond)
  redis_c.Do("SET", "test", strconv.Itoa(cnt))
}

func main() {
  pool = newPool()
  start := time.Now()

  // test用に
  // 他にもプロセス立ち上げても大丈夫
  wg := new(sync.WaitGroup)
  wg.Add(1)
  go func() {
    for i:= 0; i < 100; i++  {
      update()
    }
    wg.Done()
  }()
  for i:= 0; i < 100; i++  {
    update()
  }

  wg.Wait()
  end := time.Now();
  fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
}

