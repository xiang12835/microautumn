package main

import (
	_ "encoding/json"
	"log"
	"microautumn/utils/redis_model"
)

func sync_hello(dic map[string]interface{}) {

	log.Println("[sync_hello]...")
	log.Println("[recive dict]", dic)

	for key, value := range dic {
		log.Println(key, value)
	}

}

func aysnc_do(queue *redis_model.RedisQueue) {
	value := map[string]interface{}{}
	value["hello"] = 1
	value["world"] = 2

	queue.ASync(value)
}

func main() {

	queue := redis_model.NewRedisQueue("channel.test")
	aysnc_do(queue)

	//queue do work
	queue.Do(sync_hello)

}
