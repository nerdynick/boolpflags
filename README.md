# boolpflags

A collection of [PFlag](https://github.com/spf13/pflag) [Value Structs](https://github.com/spf13/pflag/blob/master/flag.go#L187-L191)
for creating Boolean like Flags with non-bool value types. 
These are useful when you want to provide a collection of short hand flags to define a single output.
These flag take car to not polute the Usage output with the actualy default value as well. 
