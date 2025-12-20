package main

import (
	"fmt"
	j "jasgo/core"
)

func main() {
	class := j.NewClass("HelloWorld", j.Public).
		Source("HelloWorld.java").
		DefaultConstructor()

	main := class.MainMethod()
	main.SetLimits(2, 1).
		Comment("System.out.println(\"Hello, World!\")").
		GetStatic("java/lang/System", "out", "Ljava/io/PrintStream;").
		Ldc("Hello, World!").
		InvokeVirtual("java/io/PrintStream", "println", "(Ljava/lang/String;)V").
		Return()

	fmt.Println(class.Generate())
}
