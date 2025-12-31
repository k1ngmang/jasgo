.source Tree.java
.class public Tree
.super java/lang/Object

.method public <init>()V
    .limit stack 1
    .limit locals 1
    aload_0
    invokespecial java/lang/Object/<init>()V
    return
.end method

.method public static main([Ljava/lang/String;)V
    .limit stack 6
    .limit locals 4
    bipush 30
    istore_1
    iconst_0
    istore_2
row_0:
    iload_2
    iload_1
    if_icmpge row_end_3
    iconst_0
    istore_3
space_1:
    iload_3
    iload_1
    iload_2
    isub
    iconst_1
    isub
    if_icmpge space_end_4
    getstatic java/lang/System/out Ljava/io/PrintStream;
    ldc " "
    invokevirtual java/io/PrintStream/print(Ljava/lang/String;)V
    iinc 3 1
    goto space_1
space_end_4:
    iconst_0
    istore_3
star_2:
    iload_3
    iload_2
    iconst_2
    imul
    iconst_1
    iadd
    if_icmpge star_end_5
    getstatic java/lang/System/out Ljava/io/PrintStream;
    ldc "*"
    invokevirtual java/io/PrintStream/print(Ljava/lang/String;)V
    iinc 3 1
    goto star_2
star_end_5:
    getstatic java/lang/System/out Ljava/io/PrintStream;
    invokevirtual java/io/PrintStream/println()V
    iinc 2 1
    goto row_0
row_end_3:
    return
.end method
