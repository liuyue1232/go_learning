package string_test

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s)

	s = "Hello"
	t.Log(len(s))

	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))

	c := []rune(s)
	t.Log(c)
}

func TestWithString(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c, %[1]d", c)
	}
}
