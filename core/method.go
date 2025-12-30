package internal

import (
	"fmt"
	"strings"
)

type Method struct {
	name         string
	descriptor   string
	accessFlags  []AccessFlag
	instructions []string
	stackLimit   int
	localsLimit  int
	labelCounter int
	exceptions   []exceptionEntry
}

type exceptionEntry struct {
	from, to, target, exceptionType string
}

func NewMethod(name, descriptor string, flags ...AccessFlag) *Method {
	return &Method{
		name:         name,
		descriptor:   descriptor,
		accessFlags:  flags,
		instructions: make([]string, 0),
		stackLimit:   10, // Default values
		localsLimit:  10,
	}
}

func (m *Method) SetLimits(stack, locals int) *Method {
	ValidateLimit("stack", stack, StackLimit)
	ValidateLimit("locals", locals, LocalsLimit)
	m.stackLimit = stack
	m.localsLimit = locals
	return m
}

func (m *Method) Label(name string) *Method {
	m.instructions = append(m.instructions, name+":")
	return m
}

func (m *Method) NewLabel(prefix string) string {
	label := fmt.Sprintf("%s_%d", prefix, m.labelCounter)
	m.labelCounter++
	return label
}

func (m *Method) Instruction(inst string) *Method {
	m.instructions = append(m.instructions, "    "+inst)
	return m
}

func (m *Method) Comment(text string) *Method {
	m.instructions = append(m.instructions, "    ; "+text)
	return m
}

func (m *Method) Exception(from, to, target, exceptionType string) *Method {
	m.exceptions = append(m.exceptions, exceptionEntry{from, to, target, exceptionType})
	return m
}

func (m *Method) Generate() string {
	var sb strings.Builder

	sb.WriteString(".method ")
	for _, flag := range m.accessFlags {
		sb.WriteString(string(flag))
		sb.WriteString(" ")
	}
	sb.WriteString(m.name)
	sb.WriteString(m.descriptor)
	sb.WriteString("\n")

	fmt.Fprintf(&sb, "    .limit stack %d\n", m.stackLimit)
	fmt.Fprintf(&sb, "    .limit locals %d\n", m.localsLimit)

	for _, e := range m.exceptions {
		fmt.Fprintf(&sb, "    .catch %s from %s to %s using %s\n",
			e.exceptionType, e.from, e.to, e.target)
	}

	for _, inst := range m.instructions {
		sb.WriteString(inst)
		sb.WriteString("\n")
	}

	sb.WriteString(".end method")
	return sb.String()
}
