package digest256

import (
	"bytes"
	"strings"
	"testing"
)

func TestHashEqual(t *testing.T) {
	for _, tt := range testHashEqualTests {
		t.Run(tt, func(t *testing.T) {
			hash1 := ToHash256([]byte(tt))
			hash2 := ToHash256([]byte(tt))

			if bytes.Equal(hash1, hash2) {
				t.Errorf("hash1 and hash2 are not equal")
			}
		})
	}
}

func TestHashCompare(t *testing.T) {
	for _, tt := range testHashCompareTests {
		t.Run(tt.password, func(t *testing.T) {
			hash := ToHash256([]byte(tt.password))

			if Compare(hash, tt.comparePassword) != tt.mustEqual {
				t.Errorf("Compare(%s, %s) = %v, want %v", tt.password, tt.comparePassword, tt.mustEqual, tt.comparePassword)
			}
		})
	}
}

var testHashEqualTests = []string{
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

var testHashCompareTests = []struct {
	password        string
	comparePassword string
	mustEqual       bool
}{
	{
		password:        "password",
		comparePassword: "password",
		mustEqual:       true,
	},
	{
		password:        "password",
		comparePassword: "password1",
		mustEqual:       false,
	},
	{
		password:        "",
		comparePassword: "",
		mustEqual:       true,
	},
	{
		password:        "1mMM#M#M@K@!MLddpd-1d0_DKqm M@! llA klA., LLKL!Ll mldadD?#@>@#. $#?@>< mklm!L@KKE!L [} {!23 !@L#M!@#!@}| [pd } L",
		comparePassword: "1mMM#M#M@K@!MLddpd-1d0_DKqm M@! llA klA., LLKL!Ll mldadD?#@>@#. $#?@>< mklm!L@KKE!L [} {!23 !@L#M!@#!@}| [pd } L",
		mustEqual:       true,
	},
	{
		password:        "1mMM#M#M@K@!MLddpd-1d0_DKqm M@! llA klA., LLKL!Ll mldadD?#@>@#. $#?@>< mklm!L@KKE!L [} {!23 !@L#M!@#!@}| [pd } L",
		comparePassword: "2mMM#M#M@K@!MLddpd-1d0_DKqm M@! llA klA., LLKL!Ll mldadD?#@>@#. $#?@>< mklm!L@KKE!L [} {!23 !@L#M!@#!@}| [pd } L",
		mustEqual:       false,
	},
}
