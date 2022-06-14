/**
    * @Description:
    * @Author: daniel(2022/6/14)
**/
package greetings

import "testing"

func TestGreetings(t *testing.T) {
	want := "Greetings, gg"
	if got := Greetings("gg"); got != want {
		t.Errorf("Greetings() want:%q got:%q", want, got)
	}
}
