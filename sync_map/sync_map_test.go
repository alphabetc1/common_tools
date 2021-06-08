package metricstool

import (
	"strconv"
	"sync"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/go-playground/assert.v1"
)

var (
	cntNum int = 100
)

func BenchmarkSyncMapSlice(b *testing.B) {
	b.ResetTimer()
	syncMapSlice := newSyncMapSlice()
	for i := 0; i < b.N; i++ {
		for j := 0; j < cntNum; j++ {
			go func(num int) {
				key := strconv.Itoa(num)
				syncMapSlice.getAndAdd(key, 1)
			}(j)
		}
	}
	b.StopTimer()
}

func BenchmarkSyncMap(b *testing.B) {
	b.ResetTimer()
	syncMap := NewSyncMap(DefaultSyncCaps)
	for i := 0; i < b.N; i++ {
		for j := 0; j < cntNum; j++ {
			go func(num int) {
				key := strconv.Itoa(num)
				syncMap.GetAndAdd(key, 1)
			}(j)
		}
	}
	b.StopTimer()
}

func TestSyncMap(t *testing.T) {
	Convey("TestSyncMap", t, func() {
		syncMap := NewSyncMap(DefaultSyncCaps)
		syncMapSlice := newSyncMapSlice()
		var wg sync.WaitGroup
		wg.Add(cntNum * cntNum)
		for i := 0; i < cntNum; i++ {
			for j := 0; j < cntNum; j++ {
				go func(num int) {
					key := strconv.Itoa(num)
					syncMap.GetAndAdd(key, 1)
					syncMapSlice.getAndAdd(key, 1)
					wg.Done()
				}(j)
			}
		}
		wg.Wait()
		sum := 0
		Convey("Check num for each key", func() {
			for key, val1 := range syncMapSlice.schema2conn {
				val2 := syncMap.GetAndAdd(key, 0)
				assert.Equal(t, val1, val2)
				sum += val2
			}
			Convey("Check total", func() {
				So(sum == cntNum*cntNum, ShouldBeTrue)
			})
		})
	})
}
