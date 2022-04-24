package snowflake

import (
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node = nil

func init() {
	if nodes, err := snowflake.NewNode(1); err != nil {
		panic("snowflake init faild")
	} else {
		node = nodes
	}
}

// RandomUID 简易发号器
func RandomUID() int64 {
	return node.Generate().Int64()
}
