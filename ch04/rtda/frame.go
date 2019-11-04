package rtda

type Frame struct {
	lower *Frame
	// 保存局部变量
	localVars LocalVars
	// 保存操作数栈指针
	operandStack *OperandStack
}

func (self Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}
