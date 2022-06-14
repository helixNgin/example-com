/**
    * @Description:
    * @Author: daniel(2022/6/14)
**/
package hello

import "testing"

func TestGreetings(t *testing.T) {
	want := "Hello, gg"
	if got := Hello("gg"); got != want {
		t.Errorf("Hello() want:%q got:%q", want, got)
	}
}
