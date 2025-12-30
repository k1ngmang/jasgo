.source BubbleSort.java
.class public BubbleSort
.super java/lang/Object

.method public <init>()V
    .limit stack 1
    .limit locals 1
    aload_0
    invokespecial java/lang/Object/<init>()V
    return
.end method

.method public static sort([I)V
    .limit stack 5
    .limit locals 5
    ; int n = arr.length
    aload_0
    arraylength
    istore_1
    ; for (int i = 0; i < n - 1; i++)
    iconst_0
    istore_2
outer_0:
    iload_2
    iload_1
    iconst_1
    isub
    if_icmpge outer_end_1
    ; for (int j = 0; j < n - i - 1; j++)
    iconst_0
    istore_3
inner_2:
    iload_3
    iload_1
    iload_2
    isub
    iconst_1
    isub
    if_icmpge inner_end_3
    ; if (arr[j] > arr[j + 1])
    aload_0
    iload_3
    iaload
    aload_0
    iload_3
    iconst_1
    iadd
    iaload
    if_icmple no_swap_4
    ; swap arr[j] and arr[j+1]
    aload_0
    iload_3
    iaload
    istore 4
    aload_0
    iload_3
    aload_0
    iload_3
    iconst_1
    iadd
    iaload
    iastore
    aload_0
    iload_3
    iconst_1
    iadd
    iload 4
    iastore
no_swap_4:
    iinc 3 1
    goto inner_2
inner_end_3:
    iinc 2 1
    goto outer_0
outer_end_1:
    return
.end method

.method public static printArray([I)V
    .limit stack 4
    .limit locals 3
    ; Print [
    getstatic java/lang/System/out Ljava/io/PrintStream;
    ldc "["
    invokevirtual java/io/PrintStream/print(Ljava/lang/String;)V
    ; for (int i = 0; i < arr.length; i++)
    iconst_0
    istore_1
    aload_0
    arraylength
    istore_2
loop_0:
    iload_1
    iload_2
    if_icmpge end_1
    ; Print arr[i]
    getstatic java/lang/System/out Ljava/io/PrintStream;
    aload_0
    iload_1
    iaload
    invokevirtual java/io/PrintStream/print(I)V
    ; Print comma if not last
    iload_1
    iload_2
    iconst_1
    isub
    if_icmpge end_1
    getstatic java/lang/System/out Ljava/io/PrintStream;
    ldc ", "
    invokevirtual java/io/PrintStream/print(Ljava/lang/String;)V
    iinc 1 1
    goto loop_0
end_1:
    getstatic java/lang/System/out Ljava/io/PrintStream;
    ldc "]"
    invokevirtual java/io/PrintStream/println(Ljava/lang/String;)V
    return
.end method

.method public static main([Ljava/lang/String;)V
    .limit stack 6
    .limit locals 2
    ; int[] arr = {64, 34, 25, 12, 22, 11, 90}
    bipush 7
    newarray int
    dup
    iconst_0
    bipush 64
    iastore
    dup
    iconst_1
    bipush 34
    iastore
    dup
    iconst_2
    bipush 25
    iastore
    dup
    iconst_3
    bipush 12
    iastore
    dup
    iconst_4
    bipush 22
    iastore
    dup
    iconst_5
    bipush 11
    iastore
    dup
    bipush 6
    bipush 90
    iastore
    astore_1
    ; Print before sorting
    getstatic java/lang/System/out Ljava/io/PrintStream;
    ldc "Before sorting:"
    invokevirtual java/io/PrintStream/println(Ljava/lang/String;)V
    aload_1
    invokestatic BubbleSort/printArray([I)V
    ; Sort the array
    aload_1
    invokestatic BubbleSort/sort([I)V
    ; Print after sorting
    getstatic java/lang/System/out Ljava/io/PrintStream;
    ldc "After sorting:"
    invokevirtual java/io/PrintStream/println(Ljava/lang/String;)V
    aload_1
    invokestatic BubbleSort/printArray([I)V
    return
.end method
