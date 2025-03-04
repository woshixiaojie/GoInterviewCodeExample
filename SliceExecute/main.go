package main

import "fmt"

func main() {

	doAppend := func(s []int) {
		s = append(s, 1)
		printLengthAndCapacity(s)
	}

	s := make([]int, 8, 8)

	doAppend(s[:4])
	printLengthAndCapacity(s)

	doAppend(s)
	printLengthAndCapacity(s)
}

func printLengthAndCapacity(s []int) {
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d \n", len(s), cap(s))
}

//**第一次调用 `doAppend(s[:4])` 的输出：**

//1. **函数内部 `doAppend` 的打印：**
//   ```
//   [0 0 0 0 1]
//   len=5 cap=8
//   ```
//   - **原因**：传入 `s[:4]`（长度4，容量8）。追加元素后，底层数组的索引4被设为1，长度变为5。此时原数组的索引4也被修改为1，但 `s[:4]` 的长度在外部仍为4。

//2. **主函数第一次打印 `s`：**
//   ```
//   [0 0 0 0 1 0 0 0]
//   len=8 cap=8
//   ```
//   - **原因**：`s` 的底层数组被修改（索引4为1），但 `s` 的长度和容量仍为8。打印时显示所有8个元素，其中索引4是1，其余为0。

//**第二次调用 `doAppend(s)` 的输出：**
//
//3. **函数内部 `doAppend` 的打印：**
//   ```
//   [0 0 0 0 1 0 0 0 1]
//   len=9 cap=16
//   ```
//   - **原因**：传入的 `s` 长度和容量均为8。追加元素时容量不足，触发扩容（容量翻倍为16）。新切片指向新数组，前8个元素与原数组相同（索引4为1），追加的1成为第9个元素。

//4. **主函数第二次打印 `s`：**
//   ```
//   [0 0 0 0 1 0 0 0]
//   len=8 cap=8
//   ```
//   - **原因**：原 `s` 未扩容，仍指向旧数组，长度和容量不变。打印结果与第一次主函数打印相同。

//**关键点总结：**
//- **切片传递机制**：函数参数是切片副本，共享底层数组，但长度和容量独立。
//- **扩容影响**：若 `append` 导致扩容，新切片指向新数组，原切片不受影响。
//- **子切片操作**：`s[:4]` 的容量为8（原数组容量），因此追加元素会修改原数组的索引4。
