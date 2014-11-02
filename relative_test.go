package main

import (
	"testing"
	"time"
)

func test(t *testing.T, delta int64, expected string) {
	actual := relative(time.Unix(time.Now().Unix()-delta, 0))
	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}

func TestRelative(t *testing.T) {
	test(t, 0, "now")
	test(t, 3*Minute, "3m")
	test(t, 4*Hour, "4h")
	test(t, 2*Day, "2d")
	test(t, 6*Week, "6w")
	test(t, 2*Year, "2y")
}
