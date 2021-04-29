package boolpflags

import (
	"github.com/spf13/pflag"
)

type BoolInterfaceValue struct {
	BoolValue
	defaultVal interface{}
	pointer    interface{}
}

func (s *BoolInterfaceValue) Set(val string) error {
	s.pointer = s.defaultVal
	return nil
}

func newBoolInterfaceValue(val interface{}, p interface{}) *BoolInterfaceValue {
	return &BoolInterfaceValue{
		defaultVal: val,
		pointer:    p,
	}
}

// BoolStringVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func BoolInterfaceVar(f *pflag.FlagSet, p interface{}, name string, value interface{}, usage string) {
	BoolInterfaceVarP(f, p, name, "", value, usage)
}

// BoolStringVar is like BoolVar, but accepts a shorthand letter that can be used after a single dash.
func BoolInterfaceVarP(f *pflag.FlagSet, p interface{}, name, shorthand string, value interface{}, usage string) {
	flag := f.VarPF(newBoolInterfaceValue(value, p), name, shorthand, usage)
	flag.NoOptDefVal = "true"
}
