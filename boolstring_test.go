package boolpflags

import (
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func setUpFlagSet() *pflag.FlagSet {
	f := pflag.NewFlagSet("test", pflag.ContinueOnError)
	return f
}

func TestBoolStringPImplicit(t *testing.T) {
	assert := assert.New(t)

	f := setUpFlagSet()
	var bString, cString string

	BoolStringVarP(f, &bString, "flagb", "b", "good", "bool value in CommandLine")
	BoolStringVarP(f, &cString, "flagc", "c", "bad", "other bool value")

	args := []string{"--flagb"}
	if err := f.Parse(args); err != nil {
		t.Error("expected no error, got ", err)
	}

	assert.Equal("good", bString)
	assert.Equal("", cString)
}

func TestBoolStringPExplicit(t *testing.T) {
	assert := assert.New(t)

	f := setUpFlagSet()
	var bString, cString string

	BoolStringVarP(f, &bString, "flagb", "b", "good", "bool value in CommandLine")
	BoolStringVarP(f, &cString, "flagc", "c", "bad", "other bool value")

	args := []string{"--flagb=true"}
	if err := f.Parse(args); err != nil {
		t.Error("expected no error, got ", err)
	}

	assert.Equal("good", bString)
	assert.Equal("", cString)
}

func TestBoolStringUsage(t *testing.T) {
	assert := assert.New(t)

	f := setUpFlagSet()
	var bString, cString string

	BoolStringVarP(f, &bString, "flagb", "b", "good", "bool value in CommandLine")
	BoolStringVarP(f, &cString, "flagc", "c", "bad", "other bool value")

	usage := f.FlagUsages()

	fmt.Printf("Usage Output:\n\n %v\n", usage)
	assert.False(strings.Contains(usage, "(default"), "Default value is getting displayed in usage")
	assert.False(strings.Contains(usage, "[="), "Default value is getting displayed in flags")
}
