package snowflake

import (
	"fmt"
	"testing"
)

func TestNewSnow(t *testing.T) {
	fmt.Println(RandomUID()) // go test -v snowflake_test.go snowflake.go
}
