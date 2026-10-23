package toolkit

import "testing"

func TestTools_RandomString(t *testing.T) {
	var testTools Tools

	testLengths := []int{0, 1, 10}
	for _, val := range testLengths {
		s := testTools.RandomString(val)
		if len(s) != val {
			t.Error("wrong length random string returned")
		}

	}

}
