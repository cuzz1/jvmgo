package rtda

import "jvmgo/ch06/rtda/heap"

type Slot struct {
	// 用于存整数类型
	num int32
	// 用于存放引用类型
	ref *heap.Object
}
