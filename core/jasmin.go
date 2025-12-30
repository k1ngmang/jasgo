package internal

import (
	"fmt"
	"strings"
)

type ClassFile struct {
	name        string
	superClass  string
	accessFlags []AccessFlag
	interfaces  []string
	source      string
	fields      []*Field
	methods     []*Method
}

func NewClass(name string, flags ...AccessFlag) *ClassFile {
	return &ClassFile{
		name:        name,
		superClass:  "java/lang/Object",
		accessFlags: flags,
		interfaces:  make([]string, 0),
		fields:      make([]*Field, 0),
		methods:     make([]*Method, 0),
	}
}

func (c *ClassFile) Super(superClass string) *ClassFile {
	c.superClass = superClass
	return c
}

func (c *ClassFile) Implements(interfaceName string) *ClassFile {
	c.interfaces = append(c.interfaces, interfaceName)
	return c
}

func (c *ClassFile) Source(filename string) *ClassFile {
	c.source = filename
	return c
}

func (c *ClassFile) AddField(field *Field) *ClassFile {
	c.fields = append(c.fields, field)
	return c
}

func (c *ClassFile) Field(name, descriptor string, flags ...AccessFlag) *ClassFile {
	return c.AddField(NewField(name, descriptor, flags...))
}

func (c *ClassFile) AddMethod(method *Method) *ClassFile {
	c.methods = append(c.methods, method)
	return c
}

func (c *ClassFile) DefaultConstructor() *ClassFile {
	m := NewMethod("<init>", "()V", Public).
		SetLimits(1, 1).
		Aload(0).
		InvokeSpecial(c.superClass, "<init>", "()V").
		Return()
	return c.AddMethod(m)
}

func (c *ClassFile) MainMethod() *Method {
	m := NewMethod("main", "([Ljava/lang/String;)V", Public, Static).
		SetLimits(10, 10)
	c.AddMethod(m)
	return m
}

func (c *ClassFile) Generate() string {
	var sb strings.Builder

	if c.source != "" {
		fmt.Fprintf(&sb, ".source %s\n", c.source)
	}

	sb.WriteString(".class ")
	for _, flag := range c.accessFlags {
		sb.WriteString(string(flag))
		sb.WriteString(" ")
	}
	sb.WriteString(c.name)
	sb.WriteString("\n")

	fmt.Fprintf(&sb, ".super %s\n", c.superClass)

	for _, iface := range c.interfaces {
		fmt.Fprintf(&sb, ".implements %s\n", iface)
	}

	sb.WriteString("\n")

	for _, field := range c.fields {
		sb.WriteString(field.Generate())
		sb.WriteString("\n")
	}

	if len(c.fields) > 0 {
		sb.WriteString("\n")
	}

	for i, method := range c.methods {
		sb.WriteString(method.Generate())
		if i < len(c.methods)-1 {
			sb.WriteString("\n\n")
		} else {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}
