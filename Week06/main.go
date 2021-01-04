package main

import (
	"log"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/errgroup"
)

// metric interface, metric entry should implement this interface
type Metricer interface {
	Add() error
	Count() int64
}

type Windower interface {
	WindowStart() int64
	SetWindowStart(s int64)
}

type Window struct {
	num         int64
	windowStart int64
}

func (w *Window) Add() error {
	atomic.AddInt64(&w.num, 1)
	return nil
}

func (w *Window) WindowStart() int64 {
	return atomic.LoadInt64(&w.windowStart)
}

func (w *Window) SetWindowStart(s int64) {
	atomic.StoreInt64(&w.windowStart, s)
}

func (w *Window) ResetCount() {
	atomic.StoreInt64(&w.num, 0)
}

func (w *Window) Count() int64 {
	return atomic.LoadInt64(&w.num)
}

type LeapArray struct {
	//时间窗口的长度 microsecond
	windowLength int

	//采样窗口的个数
	sampleCount int

	//以毫秒为单位
	intervalInMs int

	array []*Window

	mu sync.Mutex
}

func NewLeapArray(windowLength int, intervalInSec int) *LeapArray {
	sampleCount := intervalInSec * 1000 / windowLength
	la := &LeapArray{
		windowLength: windowLength,
		intervalInMs: intervalInSec * 1000,
		sampleCount:  sampleCount,
		array:        []*Window{},
	}
	for i := 0; i < sampleCount+1; i++ {
		la.array = append(la.array, &Window{
			num:         0,
			windowStart: 0,
		})
	}
	return la
}

func (la *LeapArray) GetCurrentWindow(time int64) *Window {
	w := la.GetWindow(time)
	//判断window的状态
	time = time - time%int64(la.windowLength)
	if w.WindowStart() == 0 || time > w.WindowStart() {
		w.SetWindowStart(time)
		w.ResetCount()
	}

	return w
}

func (la *LeapArray) GetWindow(time int64) *Window {
	la.mu.Lock()
	defer la.mu.Unlock()
	timeId := time / int64(la.windowLength)
	idx := timeId % int64(la.sampleCount+1)
	time = time - time%int64(la.windowLength)
	w := la.array[idx]
	return w
}

func (la *LeapArray) Add() error {
	time := time.Now().UnixNano() / int64(time.Microsecond)
	w := la.GetCurrentWindow(time)
	w.Add()
	return nil
}

func (la *LeapArray) GetStastics() int64 {
	time := time.Now().UnixNano() / int64(time.Microsecond)
	ws := []*Window{}
	for i := 0; i < la.sampleCount; i++ {
		ws = append(ws, la.GetWindow(time-int64(i*la.windowLength)))
	}
	total := int64(0)
	for _, w := range ws {
		total += w.Count()
	}
	return total
}

func (la *LeapArray) PrintStastics() {
	time := time.Now().UnixNano() / int64(time.Microsecond)
	format := "window started at : %d; window num is %d"
	for i := 0; i < la.sampleCount; i++ {
		w := la.GetWindow(time - int64(i*la.windowLength))
		log.Printf(format, w.windowStart, w.num)
	}
	log.Printf("\n")
}

func main() {
	la := NewLeapArray(500, 2)
	var wg errgroup.Group

	wg.Go(func() error {
		for {
			la.Add()
		}
	})
	wg.Go(func() error {
		for {
			la.PrintStastics()
			time.Sleep(time.Microsecond * 100)
		}
	})
	wg.Wait()
}
