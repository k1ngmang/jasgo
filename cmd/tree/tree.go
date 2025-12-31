package main

import (
	"fmt"
	j "jasgo/core"
)

func main() {
	class := j.NewClass("Tree", j.Public).
		Source("Tree.java").
		DefaultConstructor()

	// public static void main(String[] args)
	main := class.MainMethod().
		SetLimits(6, 4)

	/*
	   locals:
	   0 = args
	   1 = height
	   2 = i (row)
	   3 = j (inner)
	*/

	rowLoop := main.NewLabel("row")
	spaceLoop := main.NewLabel("space")
	starLoop := main.NewLabel("star")
	rowEnd := main.NewLabel("row_end")
	spaceEnd := main.NewLabel("space_end")
	starEnd := main.NewLabel("star_end")

	main.
		// int height = 30;
		Ldc(30).
		Istore(1).

		// i = 0
		Ldc(0).
		Istore(2).

		// for (i < height)
		Label(rowLoop).
		Iload(2).
		Iload(1).
		IfIcmpGe(rowEnd).

		// ----- print spaces -----
		Ldc(0).
		Istore(3).
		Label(spaceLoop).
		Iload(3).
		Iload(1).
		Iload(2).
		Isub().
		Ldc(1).
		Isub().
		IfIcmpGe(spaceEnd).
		GetStatic("java/lang/System", "out", j.ObjectType("java/io/PrintStream")).
		Ldc(" ").
		InvokeVirtual("java/io/PrintStream", "print", "(Ljava/lang/String;)V").
		Iinc(3, 1).
		Goto(spaceLoop).
		Label(spaceEnd).

		// ----- print stars -----
		Ldc(0).
		Istore(3).
		Label(starLoop).
		Iload(3).
		Iload(2).
		Ldc(2).
		Imul().
		Ldc(1).
		Iadd().
		IfIcmpGe(starEnd).
		GetStatic("java/lang/System", "out", j.ObjectType("java/io/PrintStream")).
		Ldc("*").
		InvokeVirtual("java/io/PrintStream", "print", "(Ljava/lang/String;)V").
		Iinc(3, 1).
		Goto(starLoop).
		Label(starEnd).

		// println()
		GetStatic("java/lang/System", "out", j.ObjectType("java/io/PrintStream")).
		InvokeVirtual("java/io/PrintStream", "println", "()V").

		// i++
		Iinc(2, 1).
		Goto(rowLoop).
		Label(rowEnd).
		Return()

	fmt.Println(class.Generate())
}
