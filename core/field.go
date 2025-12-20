package internal

import (
	"strings"
)

type Field struct {
	name        string
	descriptor  string
	accessFlags []AccessFlag
	initValue   *string
}

func NewField(name, descriptor string, flags ...AccessFlag) *Field {
	return &Field{
		name:        name,
		descriptor:  descriptor,
		accessFlags: flags,
	}
}

func (f *Field) WithValue(value string) *Field {
	f.initValue = &value
	return f
}

func (f *Field) Generate() string {
	var sb strings.Builder
	sb.WriteString(".field ")
	
	for _, flag := range f.accessFlags {
		sb.WriteString(string(flag))
		sb.WriteString(" ")
	}
	
	sb.WriteString(f.name)
	sb.WriteString(" ")
	sb.WriteString(f.descriptor)
	
	if f.initValue != nil {
		sb.WriteString(" = ")
		sb.WriteString(*f.initValue)
	}
	
	return sb.String()
}