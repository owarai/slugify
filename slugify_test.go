package slugify

import (
	"strconv"
	"testing"
)

func TestFormat(t *testing.T) {
	type args struct {
		s            string
		allowUnicode bool
	}
	tests := []struct {
		args args
		want string
	}{
		{args: args{"语言-—-—-如何使用？演示：诱导", true}, want: "语言-如何使用-演示-诱导"},
		{args: args{"hello, 你好，world! 世界！", true}, want: "hello-你好-world-世界"},
		{args: args{"hello, 你好，world! 世界！", false}, want: "hello-world"},
		{args: args{"hello, ハロー，world! ワールド！", true}, want: "hello-ハロー-world-ワールド"},
		{args: args{"hello, ハロー，world! ワールド！", false}, want: "hello-world"},
	}
	for _, tt := range tests {
		t.Run(strconv.FormatBool(tt.args.allowUnicode)+":"+tt.args.s, func(t *testing.T) {
			if got := Format(tt.args.s, tt.args.allowUnicode); got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
