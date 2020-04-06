package stringutils

import "testing"

func TestReverseRune(t *testing.T) {
	cases := []struct{
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hi ALL", "ALL hi"},
		{"", ""},
	}
	for _, c := range cases {
		got := ReverseRune(c.in)
		if got != c.want {
			t.Errorf("ReverseRune(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
