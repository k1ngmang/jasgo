package main

import (
	"fmt"
	j "jasgo/core"
)

func main() {
	class := j.NewClass("BubbleSort", j.Public).
		Source("BubbleSort.java").
		DefaultConstructor()

	// void sort(int[] arr)
	sort := j.NewMethod("sort", "([I)V", j.Public, j.Static).
		SetLimits(5, 5)

	// Local variables:
	// 0 = arr
	// 1 = n (length)
	// 2 = i (outer loop)
	// 3 = j (inner loop)
	// 4 = temp (for swap)

	outerLoop := sort.NewLabel("outer")
	outerEnd := sort.NewLabel("outer_end")
	innerLoop := sort.NewLabel("inner")
	innerEnd := sort.NewLabel("inner_end")
	noSwap := sort.NewLabel("no_swap")

	sort.
		Comment("int n = arr.length").
		Aload(0).
		ArrayLength().
		Istore(1).
		Comment("for (int i = 0; i < n - 1; i++)").
		Ldc(0).
		Istore(2).
		Label(outerLoop).
		Iload(2).
		Iload(1).
		Ldc(1).
		Isub().
		IfIcmpGe(outerEnd).
		Comment("for (int j = 0; j < n - i - 1; j++)").
		Ldc(0).
		Istore(3).
		Label(innerLoop).
		Iload(3).
		Iload(1).
		Iload(2).
		Isub().
		Ldc(1).
		Isub().
		IfIcmpGe(innerEnd).
		Comment("if (arr[j] > arr[j + 1])").
		Aload(0).
		Iload(3).
		Iaload().
		Aload(0).
		Iload(3).
		Ldc(1).
		Iadd().
		Iaload().
		IfIcmpLe(noSwap).
		Comment("swap arr[j] and arr[j+1]").
		Aload(0).
		Iload(3).
		Iaload().
		Istore(4).
		Aload(0).
		Iload(3).
		Aload(0).
		Iload(3).
		Ldc(1).
		Iadd().
		Iaload().
		Iastore().
		Aload(0).
		Iload(3).
		Ldc(1).
		Iadd().
		Iload(4).
		Iastore().
		Label(noSwap).
		Iinc(3, 1).
		Goto(innerLoop).
		Label(innerEnd).
		Iinc(2, 1).
		Goto(outerLoop).
		Label(outerEnd).
		Return()
	class.AddMethod(sort)

	// void printArray(int[] arr)
	printArr := j.NewMethod("printArray", "([I)V", j.Public, j.Static).
		SetLimits(4, 3)

	printLoop := printArr.NewLabel("loop")
	printEnd := printArr.NewLabel("end")

	printArr.
		Comment("Print [").
		GetStatic("java/lang/System", "out", j.ObjectType("java/io/PrintStream")).
		Ldc("[").
		InvokeVirtual("java/io/PrintStream", "print", "(Ljava/lang/String;)V").
		Comment("for (int i = 0; i < arr.length; i++)").
		Ldc(0).
		Istore(1).
		Aload(0).
		ArrayLength().
		Istore(2).
		Label(printLoop).
		Iload(1).
		Iload(2).
		IfIcmpGe(printEnd).
		Comment("Print arr[i]").
		GetStatic("java/lang/System", "out", j.ObjectType("java/io/PrintStream")).
		Aload(0).
		Iload(1).
		Iaload().
		InvokeVirtual("java/io/PrintStream", "print", "(I)V").
		Comment("Print comma if not last").
		Iload(1).
		Iload(2).
		Ldc(1).
		Isub().
		IfIcmpGe(printEnd).
		GetStatic("java/lang/System", "out", j.ObjectType("java/io/PrintStream")).
		Ldc(", ").
		InvokeVirtual("java/io/PrintStream", "print", "(Ljava/lang/String;)V").
		Iinc(1, 1).
		Goto(printLoop).
		Label(printEnd).
		GetStatic("java/lang/System", "out", j.ObjectType("java/io/PrintStream")).
		Ldc("]").
		InvokeVirtual("java/io/PrintStream", "println", "(Ljava/lang/String;)V").
		Return()
	class.AddMethod(printArr)

	// Main method
	main := class.MainMethod().SetLimits(6, 2)
	main.
		Comment("int[] arr = {64, 34, 25, 12, 22, 11, 90}").
		Ldc(7).
		NewArray("int").
		Dup().
		Ldc(0).
		Ldc(64).
		Iastore().
		Dup().
		Ldc(1).
		Ldc(34).
		Iastore().
		Dup().
		Ldc(2).
		Ldc(25).
		Iastore().
		Dup().
		Ldc(3).
		Ldc(12).
		Iastore().
		Dup().
		Ldc(4).
		Ldc(22).
		Iastore().
		Dup().
		Ldc(5).
		Ldc(11).
		Iastore().
		Dup().
		Ldc(6).
		Ldc(90).
		Iastore().
		Astore(1).
		Comment("Print before sorting").
		GetStatic("java/lang/System", "out", j.ObjectType("java/io/PrintStream")).
		Ldc("Before sorting:").
		InvokeVirtual("java/io/PrintStream", "println", "(Ljava/lang/String;)V").
		Aload(1).
		InvokeStatic("BubbleSort", "printArray", "([I)V").
		Comment("Sort the array").
		Aload(1).
		InvokeStatic("BubbleSort", "sort", "([I)V").
		Comment("Print after sorting").
		GetStatic("java/lang/System", "out", j.ObjectType("java/io/PrintStream")).
		Ldc("After sorting:").
		InvokeVirtual("java/io/PrintStream", "println", "(Ljava/lang/String;)V").
		Aload(1).
		InvokeStatic("BubbleSort", "printArray", "([I)V").
		Return()

	fmt.Println(class.Generate())
}
