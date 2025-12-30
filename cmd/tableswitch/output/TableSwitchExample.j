.source TableSwitchExample.java
.class public TableSwitchExample
.super java/lang/Object

.method public <init>()V
    .limit stack 1
    .limit locals 1
    aload_0
    invokespecial java/lang/Object/<init>()V
    return
.end method

.method public static testSwitch(I)V
    .limit stack 2
    .limit locals 2
    ; switch(value)
    iload_0
    tableswitch 0
        case0_0
        case1_1
        case2_2
        default: default_3
case0_0:
    getstatic java/lang/System/out Ljava/io/PrintStream;
    ldc "Value is 0"
    invokevirtual java/io/PrintStream/println(Ljava/lang/String;)V
    goto default_3
case1_1:
    getstatic java/lang/System/out Ljava/io/PrintStream;
    ldc "Value is 1"
    invokevirtual java/io/PrintStream/println(Ljava/lang/String;)V
    goto default_3
case2_2:
    getstatic java/lang/System/out Ljava/io/PrintStream;
    ldc "Value is 2"
    invokevirtual java/io/PrintStream/println(Ljava/lang/String;)V
    goto default_3
default_3:
    return
.end method

.method public static main([Ljava/lang/String;)V
    .limit stack 2
    .limit locals 1
    iconst_1
    invokestatic TableSwitchExample/testSwitch(I)V
    return
.end method
