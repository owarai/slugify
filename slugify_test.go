package slugify

import (
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {
	s := "语言是强大的武器-—-—-如何使用它？社交工程攻击演示继续：致命的诱导"
	fmt.Println(Format(s, true))
}
