package internal

type AccessFlag string

const (
	Public       AccessFlag = "public"
	Private      AccessFlag = "private"
	Protected    AccessFlag = "protected"
	Static       AccessFlag = "static"
	Final        AccessFlag = "final"
	Super        AccessFlag = "super"
	Synchronized AccessFlag = "synchronized"
	Volatile     AccessFlag = "volatile"
	Transient    AccessFlag = "transient"
	Native       AccessFlag = "native"
	Interface    AccessFlag = "interface"
	Abstract     AccessFlag = "abstract"
	Strict       AccessFlag = "strict"
)

const (
	TVoid    = "V"
	TInt     = "I"
	TLong    = "J"
	TFloat   = "F"
	TDouble  = "D"
	TByte    = "B"
	TChar    = "C"
	TShort   = "S"
	TBoolean = "Z"
)

func ArrayOf(typeDesc string) string {
	return "[" + typeDesc
}

func ObjectType(className string) string {
	return "L" + className + ";"
}

func MethodDescriptor(returnType string, paramTypes ...string) string {
	params := ""
	for _, p := range paramTypes {
		params += p
	}
	return "(" + params + ")" + returnType
}
