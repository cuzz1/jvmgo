package classfile

// 字段或方法的名称和描述符
// 字段描述  字段类型      方法描述   方法
// S        short       ()V      void run()
// [[D      double[][]  (FF)F    int max(float x, float y)
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
