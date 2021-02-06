## 并发执行


### Join Point
> main goroutine等待 sayHello goroutine执行完毕。这里用到了
```golang
package main
import (
	"fmt"
	"sync"
)
func main() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		fmt.Println("hello")
	}
	wg.Add(1)
	go sayHello()
	wg.Wait()
}
```

### 良好地使用闭包
> 通过值传递，避免掉了最后的值都是相同的问题
```
package main
import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "grettings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}
```

### goroutine
> goroutine 使用了抢占机制，避免了goroutine永远阻塞
```golang
package main

func main() {
	go func() {}()
}
```

### goruntine使用的内在大小

```golang
package main
import (
	"runtime"
	"sync"
	"fmt"
)
func main() {
	var memConsumed = func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() {wg.Done(); <-c }
	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}
```

### goroutine切勿需要的时长
```bash
go test -bench=. -cpu=4 ./ctx-switch_test.go
BenchmarkContextSwitch-4         4840045               227 ns/op
```

```golang
package gocsp;

import (
	"sync"
	"testing"
)
func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin :=make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			c <-token
		}
	}

	receiver := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	wg.Wait()
}
```