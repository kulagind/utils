package utils

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

var testedTarget = []string{"", "a", "23", " bbb "}

func TestContainsSuccess(t *testing.T) {
	cases := []string{
		"a",
		"",
		" bbb ",
	}
	for _, v := range cases {
		result := Contains(testedTarget, v)
		require.Equal(t, result, true)
	}
}

func TestContainsFailure(t *testing.T) {
	cases := []string{
		" ",
		" bbb",
		"bbb",
	}
	for _, v := range cases {
		result := Contains(testedTarget, v)
		require.Equal(t, result, false)
	}
}

var testedTargetInt = []int{0, 100, -100, math.MaxInt, math.MinInt}

func TestContainsIntSuccess(t *testing.T) {
	cases := []int{
		0,
		-100,
		math.MinInt,
	}
	for _, v := range cases {
		result := ContainsInt(testedTargetInt, v)
		require.Equal(t, result, true)
	}
}

func TestContainsIntFailure(t *testing.T) {
	cases := []int{
		1,
		99,
		math.MaxInt - 1,
		math.MinInt + 1,
	}
	for _, v := range cases {
		result := ContainsInt(testedTargetInt, v)
		require.Equal(t, result, false)
	}
}
