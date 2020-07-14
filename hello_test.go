package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello word."
	if got := Hello(); got != want {
		t.Error("Hello() , want: ", got, want)
	}
}
