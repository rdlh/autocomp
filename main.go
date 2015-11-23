package main

import (
	"log"
	"net/http"
  "github.com/fzzy/radix/redis"
  "github.com/fzzy/radix/extra/pool"
)

var (
  redisPool     *pool.Pool
  redisURL      = "localhost:6379"
  redisPoolSize = 10
)

func main() {
  df := func(network, addr string) (*redis.Client, error) {
    client, err := redis.Dial(network, addr)
    // fmt.Println("DIaling")
    if err != nil {
      return nil, err
    }
    err = client.Cmd("SELECT", 8).Err
    if err != nil {
      return nil, err
    }
    // err = client.Cmd("FLUSHDB").Err
    if err != nil {
      return nil, err
    }
    return client, nil
  }

  redisPool, _ = pool.NewCustomPool("tcp", redisURL, redisPoolSize, df)
  // if err != nil {
  //   // handle err
  // }

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
