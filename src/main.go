package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/blacknikka/go-k8s/domain"
	"github.com/blacknikka/go-k8s/infrastructure"
	"github.com/blacknikka/go-k8s/random"
	"github.com/gomodule/redigo/redis"
)

var redisConn redis.Conn

func main() {
	redisPool := infrastructure.NewRedis()
	redisConn = redisPool.NewRedisPool()

	rnd := random.NewRandom()
	doRandom(rnd)
}

func doRandom(rnd random.Random) {
	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-t.C:
			randomizedInt := rnd.GetRandomValue()

			repo := infrastructure.NewRandomRepository(redisConn)

			randomItem := domain.RandomItem{
				Name:   "name" + strconv.FormatInt(randomizedInt, 10),
				Random: strconv.FormatInt(randomizedInt, 10),
			}

			err := repo.Save(randomItem)

			if err != nil {
				fmt.Printf("error %v\n", err)
			} else {
				fmt.Println(randomItem.ToString())
			}
		}
	}
	redisConn.Close()
	t.Stop()
}
