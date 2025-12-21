<img src="branding/logo.png" width="600">

## About
Jasgo is a type-safe framework for generating JVM (jasmin) assembly language.
It is convenient for writing compilers for the JVM.

Example:

```go
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
```

Result:
```jasmin
.source HelloWorld.java
.class public HelloWorld
.super java/lang/Object

.method public <init>()V
    .limit stack 1
    .limit locals 1
    aload_0
    invokespecial java/lang/Object/<init>()V
    return
.end method

.method public static main([Ljava/lang/String;)V
    .limit stack 2
    .limit locals 1
    ; System.out.println("Hello, World!")
    getstatic java/lang/System/out Ljava/io/PrintStream;
    ldc "Hello, World!"
    invokevirtual java/io/PrintStream/println(Ljava/lang/String;)V
    return
.end method
```