package main

import (
	"log"
	"mocker/db"
)

func main() {
	mem := db.New("memory")
	_ = mem.Connect()
	defer mem.Close()

	redis := db.New("redis")
	_ = redis.Connect()
	defer redis.Close()

	Run(redis)
	Run(mem)

	//_ = mem.Set("MyKey", "A value")
	//_ = mem.Set("Another Key", "Another value")
	//val, _ := mem.Get("MyKey")
	//_ = redis.Set("memKey", val)
	//if val, err := redis.Get("MemKey"); err != nil {
	//	log.Printf("Error on get 'MemKey' from redis, Err: %s", err.Error())
	//} else {
	//	log.Printf("Val from redis is '%s'", val)
	//}

	log.Printf("Exiting...")

}

func Run(db db.Database) {

	_ = db.Set("MyKey", "A value")
	val, err := db.Get("MyKey")
	if err != nil {
		log.Printf("Error on get key from database, Err: %s", err.Error())
	} else {
		log.Printf("Key from db is '%s'", val)
	}

}
