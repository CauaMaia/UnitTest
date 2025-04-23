package atumalaca

import "testing"

func TestVote(t *testing.T) {
	tests := []struct {
		age      int
		expected bool
	}{
		{-322, true},
		{18, true},
		{25, true},
		{69, true},
		{70, false},
	}

	for _, test := range tests {
		result := Vote(test.age)
		if result != test.expected {
			t.Errorf("Vote(%d) = %v; expected %v", test.age, result, test.expected)
		}
	}
}
