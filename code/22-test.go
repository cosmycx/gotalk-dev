package hello

import "testing"

func Hello() string {

	return "Hello, world."

} // .Hello

// testing function Hello
func TestHello(t *testing.T) {

	want := "Hello, world."

	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	} // .if

} // .TestHello
