package cells

type Flag int
type Flags int

const (
	Line Flag = 1
)

func NewFlags(flags ...Flag) Flags {
	newFlags := Flags(0)
	for _, flag := range flags {
		newFlags = newFlags.Add(flag)
	}
	return newFlags
}

func (f Flags) Add(flag Flag) Flags {
	return Flags(int(f) | int(flag))
}

func (f Flags) HasFlag(flag Flag) bool {
	return int(f)&int(flag) != 0
}
