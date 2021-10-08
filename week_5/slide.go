package main

import (
	"sync/atomic"
	"time"
)

type SlidingWindow struct {
	Current    int64
	total      int32
	Size       int64
	Slot       []int32
	TotalLimit int32
	Limit      int32
}

//GetSlidingWindow 获取滑动窗口实例
func GetSlidingWindow(size int64, totalLimit, limit int32) SlidingWindow {
	s := SlidingWindow{Size: size, TotalLimit: totalLimit, Limit: limit}
	s.Slot = make([]int32, s.Size)
	return s
}

//Check 滑动窗口限制检查
func (s *SlidingWindow) Check() bool {
	wd := time.Now().Unix() % s.Size
	s.reset(wd)
	if atomic.LoadInt32(&s.Slot[wd]) >= s.Limit { //超出当前窗口限制
		return false
	}
	if atomic.LoadInt32(&s.total) >= s.TotalLimit { //超出总窗口限制
		return false
	}
	atomic.AddInt32(&s.Slot[wd], 1)
	atomic.AddInt32(&s.total, 1)
	return true
}

//reset 重置单个窗口数据
func (s *SlidingWindow) reset(wd int64) {
	old := atomic.SwapInt64(&s.Current, wd)
	if old == wd {
		return
	}
	nextInd := (wd+1) % s.Size
	nt := atomic.SwapInt32(&s.Slot[nextInd], 0)
	if nt > 0 {
		atomic.AddInt32(&s.total, -nt)
	}
}
