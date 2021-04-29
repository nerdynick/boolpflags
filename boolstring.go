package boolpflags

import (
	"github.com/spf13/pflag"
)

type BoolStringValue struct {
	BoolValue
	defaultVal string
	pointer    *string
}

func (s *BoolStringValue) Set(val string) error {
	*s.pointer = s.defaultVal
	return nil
}

func newBoolStringValue(val string, p *string) *BoolStringValue {
	return &BoolStringValue{
		defaultVal: val,
		pointer:    p,
	}
}

// BoolStringVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func BoolStringVar(f *pflag.FlagSet, p *string, name string, value string, usage string) {
	BoolStringVarP(f, p, name, "", value, usage)
}

// BoolStringVar is like BoolVar, but accepts a shorthand letter that can be used after a single dash.
func BoolStringVarP(f *pflag.FlagSet, p *string, name, shorthand string, value string, usage string) {
	flag := f.VarPF(newBoolStringValue(value, p), name, shorthand, usage)
	flag.NoOptDefVal = "true"
}
