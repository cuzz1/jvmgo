package rtda

type Thread struct {
	// 寄存器（Program Counter）
	pc int
	// Java 虚拟机栈指针
	stack *Stack
}

func NewThread() *Thread                    { return &Thread{stack: newStack(1024)} }
func (self *Thread) PC() int                { return self.pc } // getter
func (self *Thread) SetPC(pc int)           { self.pc = pc }   // setter
func (self *Thread) PushFrame(frame *Frame) { self.stack.push(frame) }
func (self *Thread) PopFrame() *Frame       { return self.stack.pop() }
func (self *Thread) CurrentFrame() *Frame   { return self.stack.top() }
