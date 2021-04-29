package boolpflags

type BoolValue struct {
}

func (s *BoolValue) Type() string     { return "bool" }
func (s *BoolValue) String() string   { return "false" }
func (s *BoolValue) IsBoolFlag() bool { return true }
