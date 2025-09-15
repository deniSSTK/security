package encrypt

import (
	"strings"
	"testing"
)

func TestEncryptEqual(t *testing.T) {
	for _, test := range testEncryptTests {
		t.Run(test.value, func(t *testing.T) {
			value1 := Encrypt(test.value, test.key)
			value2 := Encrypt(test.value, test.key)

			if value1 != value2 {
				t.Errorf("encrypt failed, expected %s, got %s", value1, value2)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	for i, test := range testEncryptTests {
		t.Run(string(rune(i)), func(t *testing.T) {
			encrypted := Encrypt(test.value, test.key)

			if Decrypt(encrypted, test.key) != test.value {
				t.Errorf("decrypt failed, expected %s, got %s", test.value, encrypted)
			}
		})
	}
}

var testEncryptTests = []struct {
	value string
	key   string
}{
	{value: "main.go", key: "1f2e3d4c5b6a7980"},
	{value: "lorem", key: "2a3b4c5d6e7f8091"},
	{value: "M#M@dmaml%^)!_>>1njd>$>@$K1nFNalak", key: "3c4d5e6f7a8091b2"},
	{value: " ", key: "4d5e6f7a8091b2c3"},
	{value: "", key: "5e6f7a8091b2c3d4"},
	{value: "😊😂👍", key: "6f7a8091b2c3d4e5"},
	{value: "中文测试", key: "7a8091b2c3d4e5f6"},
	{value: "longstring" + strings.Repeat("a", 100), key: "8091b2c3d4e5f607"},
	{value: "!@#$%^&*()_+-={}[]|:;<>,.?/", key: "91b2c3d4e5f60718"},
	{value: "anotherTest", key: "a2b3c4d5e6f70819"},
}
