package std

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

var chanSum int

func _TestChan1(t *testing.T) {
	c1 := make(chan bool)
	f, err := os.OpenFile("std.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	l := logrus.New()
	l.SetOutput(f)
	l.SetLevel(logrus.DebugLevel)
	if err != nil {
		t.Log(err)
	}
	go func() {
		for i := 1; i <= 1000; i++ {
			chanSum += i
		}
		c1 <- true
	}()
	<-c1
	fmt.Println("end of file")
}

func _TestChansDanger(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func _TestChansDangerWithMute(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

// 协程安全Demo
func _TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex    // 互斥锁
	var wg sync.WaitGroup // 等待队列
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1) // 加个任务
		go func() {
			defer func() {
				mut.Unlock() //函数调用完成后：解锁，保证协程安全
			}()
			mut.Lock() // 函数将要调用前：加锁，保证协程安全
			counter++
			wg.Done() // 做完任务
		}()
	}
	wg.Wait() //等待所有任务执行完毕
	t.Logf("counter = %d", counter)
}

var tsand = 1

var c = 1

func _TestRoutineOrder(t *testing.T) {
	t.Logf("%f", 1e9)
	a, b := 20, 30 // declare variables a and b
	fmt.Println("a is", a, "b is", b)
	b, c := 40, 50 // b is already declared but c is new
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90 // assign new values to already declared variables b and c
	fmt.Println("changed b is", b, "c is", c)
	var wg sync.WaitGroup
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("total count %d", counter)
}

type addsForSum struct {
	sum    int
	first  int
	second int
}

func _TestFor(t *testing.T) {
	var wg sync.WaitGroup
	var wg1 sync.WaitGroup
	var rst = make(chan []addsForSum, 100)
	// defer close(rst)
	for i := 0; i <= 10; i++ {
		wg1.Add(1)
		wg.Add(1)
		go func(i int) {
			defer func() {
				wg.Done()
				wg1.Done()
			}()
			rst <- somefor(i)
		}(i)
		wg1.Wait()
	}
	wg.Wait()
	// close(rst)
	// for v := range rst {
	// 	fmt.Println(v)
	// }
	close(rst)
	for {
		v, ok := <-rst
		if !ok {
			break
		}
		fmt.Println(v)
	}
}

func somefor(i int) (r []addsForSum) {
	r = make([]addsForSum, 0)
	for x := 0; x <= i; x++ {
		y := i - x
		r = append(r, addsForSum{
			i, x, y,
		})
	}
	return
}

//无buffer的chan的特性
func _TestChanNoBuffer(t *testing.T) {
	ch := make(chan int)
	go func() {
		time.Sleep(3e9)
		ch <- 1
	}()
	<-ch
	fmt.Println("main rt done")
}

//有buffer的 chan特性 communcation
func _TestChanWithBuffer(t *testing.T) {
	ch := make(chan int, 3)
	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(3e9)
			fmt.Println(111)
			ch <- 1
		}
	}()
	<-ch
}

// 协程不安全demo
func TestCounterThreadUnsafe(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

// 协程不安全demo + sync.Mutex 锁
func _TestCounterThreadSafeNotGood(t *testing.T) {
	counter := 0
	var mut sync.Mutex
	for i := 0; i < 5000; i++ {
		go func() {
			defer mut.Unlock()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

// 协程安全Demo
func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex    // 互斥锁
	var wg sync.WaitGroup // 等待队列
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1) // 加个任务
		go func() {
			defer func() {
				mut.Unlock() //函数调用完成后：解锁，保证协程安全
			}()
			mut.Lock() // 函数将要调用前：加锁，保证协程安全
			counter++
			wg.Done() // 做完任务
		}()
	}
	wg.Wait() //等待所有任务执行完毕
	t.Logf("counter = %d", counter)
}

func TestSelect(t *testing.T) {
	ch := make(chan int, 3)
	for i := 1; i <= 3; i++ {
		ch <- i
	}
	select {
	case ch <- 1:
		fmt.Println("send 1 to chan")
	case ch <- 1:
		fmt.Println("receive 1 from chan")
	default:
		fmt.Println("default")
	}
}

func _TestBit(t *testing.T) {
	c := make(chan int)
	go func() {
		for {
			select {
			case c <- 0:
				time.Sleep(1e9)
			case c <- 1:
				time.Sleep(2 * 1e9)
			}
		}
	}()
	for v := range c {
		fmt.Printf("%d", v)
	}
}
