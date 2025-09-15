package hash

import (
	"strings"
	"testing"
)

func TestHashEqual(t *testing.T) {
	tests := []string{
		"MD>@!>!_aeflaefA^#DK&kadm@#la",
		"hello",
		"password",
		"123km",
		"_",
		"",
		"   ",
		"😊😂👍",
		"中文测试",
		"longstring" + strings.Repeat("a", 100),
		"!@#$%^&*()_+-={}[]|:;<>,.?/",
	}

	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			hash1 := Hash(tt)
			hash2 := Hash(tt)

			if hash1 != hash2 {
				t.Errorf("hash1 and hash2 are not equal")
			}
		})
	}
}
