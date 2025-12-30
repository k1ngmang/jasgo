package main

import (
	"fmt"
	j "jasgo/core"
)

func main() {
	class := j.NewClass("LookupSwitchExample", j.Public).
		Source("LookupSwitchExample.java").
		DefaultConstructor()

	// Метод void testLookupSwitch(int value)
	test := j.NewMethod("testLookupSwitch", "(I)V", j.Public, j.Static).
		SetLimits(2, 2) // max stack, max locals

	case10 := test.NewLabel("case10")
	case20 := test.NewLabel("case20")
	case30 := test.NewLabel("case30")
	defaultLabel := test.NewLabel("default")

	// Мапа кейсов
	cases := map[int]string{
		10: case10,
		20: case20,
		30: case30,
	}

	test.
		Comment("switch(value) with non-continuous keys").
		Iload(0). // загружаем аргумент value
		LookupSwitch(defaultLabel, cases).
		Label(case10).
		GetStatic("java/lang/System", "out", j.ObjectType("java/io/PrintStream")).
		Ldc("Value is 10").
		InvokeVirtual("java/io/PrintStream", "println", "(Ljava/lang/String;)V").
		Goto(defaultLabel).
		Label(case20).
		GetStatic("java/lang/System", "out", j.ObjectType("java/io/PrintStream")).
		Ldc("Value is 20").
		InvokeVirtual("java/io/PrintStream", "println", "(Ljava/lang/String;)V").
		Goto(defaultLabel).
		Label(case30).
		GetStatic("java/lang/System", "out", j.ObjectType("java/io/PrintStream")).
		Ldc("Value is 30").
		InvokeVirtual("java/io/PrintStream", "println", "(Ljava/lang/String;)V").
		Goto(defaultLabel).
		Label(defaultLabel).
		Return()

	class.AddMethod(test)

	// Main method
	main := class.MainMethod().SetLimits(2, 1)
	main.
		Ldc(20). // попробуем значение 20
		InvokeStatic("LookupSwitchExample", "testLookupSwitch", "(I)V").
		Return()

	fmt.Println(class.Generate())
}
