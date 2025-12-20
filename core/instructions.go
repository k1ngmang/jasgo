package internal

import "fmt"

func (m *Method) Aload(index int) *Method {
	if index >= 0 && index <= 3 {
		return m.Instruction(fmt.Sprintf("aload_%d", index))
	}
	return m.Instruction(fmt.Sprintf("aload %d", index))
}

func (m *Method) Astore(index int) *Method {
	if index >= 0 && index <= 3 {
		return m.Instruction(fmt.Sprintf("astore_%d", index))
	}
	return m.Instruction(fmt.Sprintf("astore %d", index))
}

func (m *Method) Iload(index int) *Method {
	if index >= 0 && index <= 3 {
		return m.Instruction(fmt.Sprintf("iload_%d", index))
	}
	return m.Instruction(fmt.Sprintf("iload %d", index))
}

func (m *Method) Istore(index int) *Method {
	if index >= 0 && index <= 3 {
		return m.Instruction(fmt.Sprintf("istore_%d", index))
	}
	return m.Instruction(fmt.Sprintf("istore %d", index))
}

func (m *Method) Fload(index int) *Method {
	return m.Instruction(fmt.Sprintf("fload %d", index))
}

func (m *Method) Fstore(index int) *Method {
	return m.Instruction(fmt.Sprintf("fstore %d", index))
}

func (m *Method) Dload(index int) *Method {
	return m.Instruction(fmt.Sprintf("dload %d", index))
}

func (m *Method) Dstore(index int) *Method {
	return m.Instruction(fmt.Sprintf("dstore %d", index))
}

func (m *Method) Lload(index int) *Method {
	return m.Instruction(fmt.Sprintf("lload %d", index))
}

func (m *Method) Lstore(index int) *Method {
	return m.Instruction(fmt.Sprintf("lstore %d", index))
}

func (m *Method) Ldc(value interface{}) *Method {
	switch v := value.(type) {
	case string:
		return m.Instruction(fmt.Sprintf("ldc \"%s\"", v))
	case int:
		if v >= -1 && v <= 5 {
			return m.Instruction(fmt.Sprintf("iconst_%d", v))
		} else if v >= -128 && v <= 127 {
			return m.Instruction(fmt.Sprintf("bipush %d", v))
		} else if v >= -32768 && v <= 32767 {
			return m.Instruction(fmt.Sprintf("sipush %d", v))
		}
		return m.Instruction(fmt.Sprintf("ldc %d", v))
	case float32:
		return m.Instruction(fmt.Sprintf("ldc %f", v))
	case float64:
		return m.Instruction(fmt.Sprintf("ldc2_w %f", v))
	default:
		return m.Instruction(fmt.Sprintf("ldc %v", v))
	}
}

func (m *Method) Iconst(value int) *Method {
	if value == -1 {
		return m.Instruction("iconst_m1")
	} else if value >= 0 && value <= 5 {
		return m.Instruction(fmt.Sprintf("iconst_%d", value))
	}
	return m.Ldc(value)
}

func (m *Method) AconstNull() *Method {
	return m.Instruction("aconst_null")
}

func (m *Method) Iadd() *Method { return m.Instruction("iadd") }
func (m *Method) Isub() *Method { return m.Instruction("isub") }
func (m *Method) Imul() *Method { return m.Instruction("imul") }
func (m *Method) Idiv() *Method { return m.Instruction("idiv") }
func (m *Method) Irem() *Method { return m.Instruction("irem") }
func (m *Method) Ineg() *Method { return m.Instruction("ineg") }

func (m *Method) Fadd() *Method { return m.Instruction("fadd") }
func (m *Method) Fsub() *Method { return m.Instruction("fsub") }
func (m *Method) Fmul() *Method { return m.Instruction("fmul") }
func (m *Method) Fdiv() *Method { return m.Instruction("fdiv") }

func (m *Method) Dadd() *Method { return m.Instruction("dadd") }
func (m *Method) Dsub() *Method { return m.Instruction("dsub") }
func (m *Method) Dmul() *Method { return m.Instruction("dmul") }
func (m *Method) Ddiv() *Method { return m.Instruction("ddiv") }

func (m *Method) Ladd() *Method { return m.Instruction("ladd") }
func (m *Method) Lsub() *Method { return m.Instruction("lsub") }
func (m *Method) Lmul() *Method { return m.Instruction("lmul") }
func (m *Method) Ldiv() *Method { return m.Instruction("ldiv") }

func (m *Method) Iand() *Method { return m.Instruction("iand") }
func (m *Method) Ior() *Method  { return m.Instruction("ior") }
func (m *Method) Ixor() *Method { return m.Instruction("ixor") }
func (m *Method) Ishl() *Method { return m.Instruction("ishl") }
func (m *Method) Ishr() *Method { return m.Instruction("ishr") }

func (m *Method) IfIcmpEq(label string) *Method {
	return m.Instruction(fmt.Sprintf("if_icmpeq %s", label))
}

func (m *Method) IfIcmpNe(label string) *Method {
	return m.Instruction(fmt.Sprintf("if_icmpne %s", label))
}

func (m *Method) IfIcmpLt(label string) *Method {
	return m.Instruction(fmt.Sprintf("if_icmplt %s", label))
}

func (m *Method) IfIcmpLe(label string) *Method {
	return m.Instruction(fmt.Sprintf("if_icmple %s", label))
}

func (m *Method) IfIcmpGt(label string) *Method {
	return m.Instruction(fmt.Sprintf("if_icmpgt %s", label))
}

func (m *Method) IfIcmpGe(label string) *Method {
	return m.Instruction(fmt.Sprintf("if_icmpge %s", label))
}

func (m *Method) Ifeq(label string) *Method {
	return m.Instruction(fmt.Sprintf("ifeq %s", label))
}

func (m *Method) Ifne(label string) *Method {
	return m.Instruction(fmt.Sprintf("ifne %s", label))
}

func (m *Method) Iflt(label string) *Method {
	return m.Instruction(fmt.Sprintf("iflt %s", label))
}

func (m *Method) Ifle(label string) *Method {
	return m.Instruction(fmt.Sprintf("ifle %s", label))
}

func (m *Method) Ifgt(label string) *Method {
	return m.Instruction(fmt.Sprintf("ifgt %s", label))
}

func (m *Method) Ifge(label string) *Method {
	return m.Instruction(fmt.Sprintf("ifge %s", label))
}

func (m *Method) IfNull(label string) *Method {
	return m.Instruction(fmt.Sprintf("ifnull %s", label))
}

func (m *Method) IfNonNull(label string) *Method {
	return m.Instruction(fmt.Sprintf("ifnonnull %s", label))
}

func (m *Method) Goto(label string) *Method {
	return m.Instruction(fmt.Sprintf("goto %s", label))
}

func (m *Method) Lcmp() *Method  { return m.Instruction("lcmp") }
func (m *Method) Fcmpl() *Method { return m.Instruction("fcmpl") }
func (m *Method) Fcmpg() *Method { return m.Instruction("fcmpg") }
func (m *Method) Dcmpl() *Method { return m.Instruction("dcmpl") }
func (m *Method) Dcmpg() *Method { return m.Instruction("dcmpg") }


func (m *Method) InvokeVirtual(className, methodName, descriptor string) *Method {
	return m.Instruction(fmt.Sprintf("invokevirtual %s/%s%s", className, methodName, descriptor))
}

func (m *Method) InvokeStatic(className, methodName, descriptor string) *Method {
	return m.Instruction(fmt.Sprintf("invokestatic %s/%s%s", className, methodName, descriptor))
}

func (m *Method) InvokeSpecial(className, methodName, descriptor string) *Method {
	return m.Instruction(fmt.Sprintf("invokespecial %s/%s%s", className, methodName, descriptor))
}

func (m *Method) InvokeInterface(className, methodName, descriptor string, count int) *Method {
	return m.Instruction(fmt.Sprintf("invokeinterface %s/%s%s %d", className, methodName, descriptor, count))
}

func (m *Method) GetField(className, fieldName, descriptor string) *Method {
	return m.Instruction(fmt.Sprintf("getfield %s/%s %s", className, fieldName, descriptor))
}

func (m *Method) PutField(className, fieldName, descriptor string) *Method {
	return m.Instruction(fmt.Sprintf("putfield %s/%s %s", className, fieldName, descriptor))
}

func (m *Method) GetStatic(className, fieldName, descriptor string) *Method {
	return m.Instruction(fmt.Sprintf("getstatic %s/%s %s", className, fieldName, descriptor))
}

func (m *Method) PutStatic(className, fieldName, descriptor string) *Method {
	return m.Instruction(fmt.Sprintf("putstatic %s/%s %s", className, fieldName, descriptor))
}

func (m *Method) New(className string) *Method {
	return m.Instruction(fmt.Sprintf("new %s", className))
}

func (m *Method) NewArray(typeCode string) *Method {
	return m.Instruction(fmt.Sprintf("newarray %s", typeCode))
}

func (m *Method) ANewArray(className string) *Method {
	return m.Instruction(fmt.Sprintf("anewarray %s", className))
}

func (m *Method) MultiANewArray(descriptor string, dimensions int) *Method {
	return m.Instruction(fmt.Sprintf("multianewarray %s %d", descriptor, dimensions))
}

func (m *Method) ArrayLength() *Method {
	return m.Instruction("arraylength")
}

func (m *Method) Iaload() *Method  { return m.Instruction("iaload") }
func (m *Method) Iastore() *Method { return m.Instruction("iastore") }
func (m *Method) Aaload() *Method  { return m.Instruction("aaload") }
func (m *Method) Aastore() *Method { return m.Instruction("aastore") }
func (m *Method) Baload() *Method  { return m.Instruction("baload") }
func (m *Method) Bastore() *Method { return m.Instruction("bastore") }
func (m *Method) Caload() *Method  { return m.Instruction("caload") }
func (m *Method) Castore() *Method { return m.Instruction("castore") }

func (m *Method) Pop() *Method   { return m.Instruction("pop") }
func (m *Method) Pop2() *Method  { return m.Instruction("pop2") }
func (m *Method) Dup() *Method   { return m.Instruction("dup") }
func (m *Method) Dup2() *Method  { return m.Instruction("dup2") }
func (m *Method) DupX1() *Method { return m.Instruction("dup_x1") }
func (m *Method) DupX2() *Method { return m.Instruction("dup_x2") }
func (m *Method) Swap() *Method  { return m.Instruction("swap") }

func (m *Method) I2f() *Method { return m.Instruction("i2f") }
func (m *Method) I2d() *Method { return m.Instruction("i2d") }
func (m *Method) I2l() *Method { return m.Instruction("i2l") }
func (m *Method) F2i() *Method { return m.Instruction("f2i") }
func (m *Method) F2d() *Method { return m.Instruction("f2d") }
func (m *Method) D2i() *Method { return m.Instruction("d2i") }
func (m *Method) D2f() *Method { return m.Instruction("d2f") }
func (m *Method) L2i() *Method { return m.Instruction("l2i") }
func (m *Method) I2b() *Method { return m.Instruction("i2b") }
func (m *Method) I2c() *Method { return m.Instruction("i2c") }
func (m *Method) I2s() *Method { return m.Instruction("i2s") }

func (m *Method) Checkcast(className string) *Method {
	return m.Instruction(fmt.Sprintf("checkcast %s", className))
}

func (m *Method) Instanceof(className string) *Method {
	return m.Instruction(fmt.Sprintf("instanceof %s", className))
}

func (m *Method) Return() *Method  { return m.Instruction("return") }
func (m *Method) Ireturn() *Method { return m.Instruction("ireturn") }
func (m *Method) Areturn() *Method { return m.Instruction("areturn") }
func (m *Method) Freturn() *Method { return m.Instruction("freturn") }
func (m *Method) Dreturn() *Method { return m.Instruction("dreturn") }
func (m *Method) Lreturn() *Method { return m.Instruction("lreturn") }

func (m *Method) Athrow() *Method { return m.Instruction("athrow") }

func (m *Method) Iinc(index, amount int) *Method {
	return m.Instruction(fmt.Sprintf("iinc %d %d", index, amount))
}

func (m *Method) Nop() *Method { return m.Instruction("nop") }
