package main

import (
	"runtime"
	"sync"
)

func main() {
	//1.设置只有一个P可以工作
	runtime.GOMAXPROCS(1)

	wg := sync.WaitGroup{}
	var n = 10
	// 2.启动n个 goroutine打印,哪个goroutine最先打印
	wg.Add(n)

	// 创建goroutine go func -> runtime.newproc
	for i := 1; i <= n; i++ {

		go func(i int) {
			//waitGroup 计数器减一,可以忽略
			defer wg.Done()
			println("I am goroutine ", i)
		}(i)

	}

	wg.Wait()
}

//#### **1. 并发环境设置**
//- `runtime.GOMAXPROCS(1)` 强制 Go 运行时仅使用一个 P（逻辑处理器），即所有 goroutine 在单线程下通过协作式调度运行。
//- 主 goroutine 启动 10 个子 goroutine，每个子 goroutine 打印自己的编号（1~10）。

//#### **2. 调度器的关键机制**
//- **`runnext` 槽位**：每个 P 有一个高优先级的 `runnext` 槽位，新创建的 goroutine 优先放入此槽位。
//  - 若 `runnext` 已占用，原 goroutine 会被移动到本地队列的尾部，新 goroutine 占据 `runnext`。
//- **本地队列**：一个先进先出（FIFO）队列，用于存储待运行的 goroutine。

//#### **3. goroutine 创建过程**
//- **循环创建 goroutine**：
//  - 创建第一个 goroutine（`i=1`）时，直接放入 `runnext`。
//  - 创建后续 goroutine（`i=2` 到 `i=10`）时，将当前 `runnext` 中的 goroutine（如 `i=2` 创建时，`runnext` 中是 `i=1`）移动到本地队列尾部，然后将新 goroutine（如 `i=2`）放入 `runnext`。
//- **最终状态**：
//  - `runnext` 中是最后一个创建的 goroutine（`i=10`）。
//  - 本地队列按创建顺序依次存放 `i=1` 到 `i=9`。

//#### **4. 调度顺序分析**
//1. **主 goroutine 完成循环**：
//   - 调用 `wg.Wait()` 后，主 goroutine 阻塞并释放 P。
//2. **调度器开始执行**：
//   - 优先执行 `runnext` 中的 `i=10`，输出第一行。
//   - 从本地队列头部依次取出 `i=1` 到 `i=9`，按 FIFO 顺序输出。

//#### **5. 结果验证**
//- **输出顺序为 `10, 1, 2, ..., 9`**：
//  - `i=10` 位于 `runnext`，最先执行。
//  - 本地队列中的 `i=1`~`i=9` 按创建顺序依次执行。

//### **关键点总结**
//- **`runnext` 优先级**：新创建的 goroutine 优先进入 `runnext`，导致最后创建的 goroutine 最先执行。
//- **本地队列顺序**：早先创建的 goroutine 按 FIFO 顺序排列在本地队列中。
//- **单 P 环境**：无并行抢占，调度顺序完全由 `runnext` 和本地队列的机制决定。
//
//此现象揭示了 Go 调度器在高并发场景下的底层行为，尤其在单线程调度时需注意 `runnext` 对执行顺序的影响。
