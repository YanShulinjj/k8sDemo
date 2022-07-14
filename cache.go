/* ----------------------------------
*  @author suyame 2022-07-13 20:34:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var rdb *redis.Client

func init() {
	connectRedis()
}

func connectRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		DB:       0,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("连接redis出错，错误信息：%v", err)
	}
	fmt.Println("成功连接redis")
}

func addCache(user *User) error {
	jsonstr, err := json.Marshal(&user)
	if err != nil {
		return err
	}
	// 哈希类型插入数据
	rdb.Set(ctx, user.Name, jsonstr, 0)
	if err != nil {
		return err
	}
	return nil
}

func deleteCache(name string) error {
	// 删除缓存项
	_, err := rdb.Del(ctx, name).Result()
	if err != nil {
		return err
	}
	return nil
}

func getUserFormCache(key string) (*User, error) {
	// 获取某个元素
	address, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	// 反序列化
	user := User{}
	err = json.Unmarshal([]byte(address), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//func main() {
//
//	user := new(User)
//	user.Name = "suyame"
//	user.Password = "12345"
//	jsonstr, err := json.Marshal(&user)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(string(jsonstr))
//	// 哈希类型插入数据
//	rdb.Set(ctx, "user", jsonstr, 0)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	// 获取某个元素
//	address, err := rdb.Get(ctx, "user").Result()
//	if err != nil {
//		panic(err)
//	}
//	// 反序列化
//	new_user := User{}
//	err = json.Unmarshal([]byte(address), &new_user)
//	if err != nil {
//		panic(err)
//	}
//
//	//new_user.UnmarshalJSON([]byte(address))
//	fmt.Printf("%#v", new_user)
//
//}
