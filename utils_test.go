package masutils

import (
	"reflect"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  bool
	}{
		{"empty", "", true},
		{"space", " ", false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := IsEmpty(c.input)
			if got != c.want {
				t.Errorf("IsEmpty(%q) == %v, want %v", c.input, got, c.want)
			}
		})
	}
}

func TestIsNotEmpty(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  bool
	}{
		{"empty", "", false},
		{"space", " ", true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := IsNotEmpty(c.input)
			if got != c.want {
				t.Errorf("IsNotEmpty(%q) == %v, want %v", c.input, got, c.want)
			}
		})
	}
}

func TestIsNil(t *testing.T) {
	cases := []struct {
		name  string
		input interface{}
		want  bool
	}{
		{"nil", nil, true},
		{"not nil", "", false},
		{"not nil", " ", false},
		{"not nil", 1, false},
		{"not nil", 1.0, false},
		{"not nil", true, false},
		{"not nil", false, false},
		{"not nil", []string{}, false},
		{"not nil", []int{}, false},
		{"not nil", map[string]string{}, false},
		{"not nil", map[string]int{}, false},
		{"not nil", make(chan int), false},
		{"not nil", make(chan string), false},
		{"not nil", make(chan bool), false},
		{"not nil", make(chan interface{}), false},
		{"not nil", make(chan struct{}), false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := IsNil(c.input)
			if got != c.want {
				t.Errorf("IsNil(%q) == %v, want %v", c.input, got, c.want)
			}
		})
	}
}

func TestIsNotNil(t *testing.T) {
	cases := []struct {
		name  string
		input interface{}
		want  bool
	}{
		{"nil", nil, false},
		{"not nil", "", true},
		{"not nil", " ", true},
		{"not nil", 1, true},
		{"not nil", 1.0, true},
		{"not nil", true, true},
		{"not nil", false, true},
		{"not nil", []string{}, true},
		{"not nil", []int{}, true},
		{"not nil", map[string]string{}, true},
		{"not nil", map[string]int{}, true},
		{"not nil", make(chan int), true},
		{"not nil", make(chan string), true},
		{"not nil", make(chan bool), true},
		{"not nil", make(chan interface{}), true},
		{"not nil", make(chan struct{}), true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := IsNotNil(c.input)
			if got != c.want {
				t.Errorf("IsNotNil(%q) == %v, want %v", c.input, got, c.want)
			}
		})
	}
}

func TestRemoveFromSlice(t *testing.T) {
	cases := []struct {
		name  string
		input []string
		want  []string
	}{
		{"empty", []string{}, []string{}},
		{"one", []string{"a"}, []string{}},
		{"two", []string{"a", "b"}, []string{"b"}},
		{"three", []string{"a", "b", "c"}, []string{"b", "c"}},
		{"four", []string{"a", "b", "c", "d"}, []string{"b", "c", "d"}},
		{"five", []string{"a", "b", "c", "d", "e"}, []string{"b", "c", "d", "e"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := RemoveFromSlice(c.input, "a")
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("RemoveFromSlice(%q) == %v, want %v", c.input, got, c.want)
			}
		})
	}
}

func TestRemoveExtraSpaces(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", ""},
		{"space", " ", ""},
		{"space", " a ", "a"},
		{"space", " a  b ", "a b"},
		{"space", "  a   b   c ", "a b c"},
		{"space", "   a     b    c    d ", "a b c d"},
		{"space", "     a    b    c  d        e ", "a b c d e"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := RemoveExtraSpaces(c.input)
			if got != c.want {
				t.Errorf("RemoveExtraSpaces(%q) == %v, want %v", c.input, got, c.want)
			}
		})
	}
}
