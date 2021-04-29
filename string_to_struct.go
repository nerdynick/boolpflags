package boolpflags

import (
	"fmt"

	"github.com/spf13/pflag"
)

type StringToStruct struct {
	value  interface{}
	parser func(string) (interface{}, error)
}

func (s *StringToStruct) Set(val string) error {
	v, err := s.parser(val)
	if err != nil {
		return err
	}
	s.value = v
	return nil
}

func (s *StringToStruct) Type() string   { return "stringtostruct" }
func (s *StringToStruct) String() string { return fmt.Sprintf("%v", s.value) }

func StringToStructVar(f *pflag.FlagSet, p interface{}, name string, usage string, parser func(string) (interface{}, error)) {
	StringToStructVarP(f, p, name, "", usage, parser)
}

func StringToStructVarP(f *pflag.FlagSet, p interface{}, name string, short string, usage string, parser func(string) (interface{}, error)) {
	f.VarP(&StringToStruct{
		value:  p,
		parser: parser,
	}, name, short, usage)
}
