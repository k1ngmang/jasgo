.source LookupSwitchExample.java
.class public LookupSwitchExample
.super java/lang/Object

.method public <init>()V
    .limit stack 1
    .limit locals 1
    aload_0
    invokespecial java/lang/Object/<init>()V
    return
.end method

.method public static testLookupSwitch(I)V
    .limit stack 2
    .limit locals 2
    ; switch(value) with non-continuous keys
    iload_0
    lookupswitch
        10: case10_0
        20: case20_1
        30: case30_2
        default: default_3
case10_0:
    getstatic java/lang/System/out Ljava/io/PrintStream;
    ldc "Value is 10"
    invokevirtual java/io/PrintStream/println(Ljava/lang/String;)V
    goto default_3
case20_1:
    getstatic java/lang/System/out Ljava/io/PrintStream;
    ldc "Value is 20"
    invokevirtual java/io/PrintStream/println(Ljava/lang/String;)V
    goto default_3
case30_2:
    getstatic java/lang/System/out Ljava/io/PrintStream;
    ldc "Value is 30"
    invokevirtual java/io/PrintStream/println(Ljava/lang/String;)V
    goto default_3
default_3:
    return
.end method

.method public static main([Ljava/lang/String;)V
    .limit stack 2
    .limit locals 1
    bipush 20
    invokestatic LookupSwitchExample/testLookupSwitch(I)V
    return
.end method
