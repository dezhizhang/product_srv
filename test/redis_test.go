package test

import (
	"fmt"
	"log"
	"sales-product-srv/driver"
	"testing"
)

//func TestSetKey(t *testing.T) {
//	err := driver.RDB.Set("name", "刘德华", 0).Err()
//	if err != nil {
//		log.Printf("设置出错%s", err)
//	}
//	result, err := driver.RDB.Get("name").Result()
//	if err != nil {
//		log.Printf("获取出错%s", err)
//	}
//	fmt.Println(result)
//}

func TestSetKey(t *testing.T) {
	err := driver.RDB.Set("name", "刘德华", 0).Err()
	if err != nil {
		log.Printf("设置出错%s", err)
	}
	result, err := driver.RDB.Get("name").Result()
	if err != nil {
		log.Printf("获取出错%s", err)
	}
	fmt.Printf("result1=%s", result)
	err = driver.RDB.Del("name").Err()
	if err != nil {
		log.Printf("删除失败%s", err)
	}
	result, err = driver.RDB.Get("name").Result()
	if err != nil {
		log.Printf("获取出错%s", err)
	}

	//fmt.Printf()
	fmt.Printf("result2=%s", result)
}
