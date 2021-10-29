package main

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestSlidingWindow(t *testing.T) {
	s := GetSlidingWindow(100, 100, 10)
	var succ, fail int32
	for k:=0;k<3;k++ {
		for i:=0;i<100;i++ {
			go func() {
				if s.Check() {
					atomic.AddInt32(&succ, 1)
				} else {
					atomic.AddInt32(&fail, 1)
				}
			}()
		}
		time.Sleep(time.Second)
	}
	if succ != 30 || fail != 270 {
		t.Fail()
	}
}