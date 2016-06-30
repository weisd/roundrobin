package roundrobin

import (
	"testing"
)

func Test_Gcd(t *testing.T) {
	ints := []int{12, 18}

	v := ngcd(ints)
	if v != 6 {
		t.Fail()
	}
	t.Log(v)
}

func Test_RoundRobin(t *testing.T) {
	data := []RoundData{
		{"1", 1},
		{"2", 2},
		{"3", 1},
		{"4", 2},
		{"5", 1},
		{"6", 3},
		{"7", 1},
	}
	r := NewRoundRobin(data)

	c := map[string]int{}

	for i := 0; i < 100; i++ {
		ser := r.Get().(RoundData).Data.(string)
		if _, ok := c[ser]; !ok {
			c[ser] = 0
		}

		c[ser] += 1
		t.Log(ser)
	}

	t.Log(c)
}
