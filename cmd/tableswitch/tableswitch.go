package main

import (
	"fmt"
	j "jasgo/core"
)

func main() {
	class := j.NewClass("TableSwitchExample", j.Public).
		Source("TableSwitchExample.java").
		DefaultConstructor()

	// void testSwitch(int value)
	test := j.NewMethod("testSwitch", "(I)V", j.Public, j.Static).
		SetLimits(2, 2) // max stack, max locals

	label0 := test.NewLabel("case0")
	label1 := test.NewLabel("case1")
	label2 := test.NewLabel("case2")
	defaultLabel := test.NewLabel("default")

	test.
		Comment("switch(value)").
		Iload(0).
		TableSwitch(defaultLabel, 0, []string{
			label0, // case 0
			label1, // case 1
			label2, // case 2
		}).
		Label(label0).
		GetStatic("java/lang/System", "out", j.ObjectType("java/io/PrintStream")).
		Ldc("Value is 0").
		InvokeVirtual("java/io/PrintStream", "println", "(Ljava/lang/String;)V").
		Goto(defaultLabel).
		Label(label1).
		GetStatic("java/lang/System", "out", j.ObjectType("java/io/PrintStream")).
		Ldc("Value is 1").
		InvokeVirtual("java/io/PrintStream", "println", "(Ljava/lang/String;)V").
		Goto(defaultLabel).
		Label(label2).
		GetStatic("java/lang/System", "out", j.ObjectType("java/io/PrintStream")).
		Ldc("Value is 2").
		InvokeVirtual("java/io/PrintStream", "println", "(Ljava/lang/String;)V").
		Goto(defaultLabel).
		Label(defaultLabel).
		Return()

	class.AddMethod(test)

	// Main method
	main := class.MainMethod().SetLimits(2, 1)
	main.
		Ldc(1). // Testing value 1
		InvokeStatic("TableSwitchExample", "testSwitch", "(I)V").
		Return()

	fmt.Println(class.Generate())
}
