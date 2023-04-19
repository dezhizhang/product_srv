package utils

import (
	"github.com/bwmarrin/snowflake"
	"log"
)

func SnowflakeId() (string, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Printf("生成雪花算法失败")
		return "", err
	}

	id := node.Generate().String()
	return id, nil
}
