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