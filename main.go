package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

const ipPort = "localhost:6379"

func main() {
	//接続を取得
	conn, _ := GetRedisConnection(ipPort)
	//接続を使ってSet
	s, err := Set(conn, "hoge", "moge")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	//接続を使ってGet
	s2, err2 := Get(conn, "hoge")
	if err2 != nil {
		fmt.Println(err)
	}
	fmt.Println(s2)
}

func GetRedisConnection(ipAndPort string) (conn redis.Conn, err error) {
	conn, err = redis.Dial("tcp", ipAndPort)
	return
}

func Set(conn redis.Conn, key, value string) (s interface{}, err error) {
	s, err = conn.Do("set", key, value)
	return
}

func Get(conn redis.Conn, key string) (s interface{}, err error) {
	s, err = conn.Do("get", key)
	return
}
