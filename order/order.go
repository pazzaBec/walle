package order

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	typeInt = iota
	typeInt64
	typeBool
	typeString
)

//Command a struct
type Command struct {
	Name   string
	Usage  string
	flags  []flag
	args   []string
	Action func(args []string) error
	parent *Command
	childs []*Command
}

type flag struct {
	name  string
	usage string
	alias string
	stamp int
	value interface{}
}

//Execute execute the command object
func (cmd *Command) Execute() error {
	if cmd.parent == nil {
		cmd.args = os.Args
	} else {
		cmd.args = cmd.parent.args[1:]
	}
	return nil
}

func (cmd *Command) help() {
	fmt.Println(cmd.Usage)
}

//Entrys a command's args and flags for application
type Entrys struct {
	args  []string
	flags map[string]interface{}
}

//Parse get os args, and analysis
func (ens *Entrys) Parse() {
	ens.args = os.Args[1:]
	ens.flags = make(map[string]interface{})
	for i := 0; i < len(ens.args); {
		if strings.HasPrefix(ens.args[i], "-") == true {
			ens.flags[ens.args[i][1:]] = getFlag(ens.args[i+1])
			if i+2 == len(ens.args) {
				ens.args = ens.args[:i]
			} else {
				ens.args = append(ens.args[:i], ens.args[i+2:]...)
			}
		} else {
			i++
		}
	}
}

func getFlag(v string) interface{} {
	if i, err := strconv.Atoi(v); err == nil {
		return i
	}
	if v == "true" || v == "TRUE" {
		return true
	}
	if v == "false" || v == "FALSE" {
		return false
	}
	return v
}

//Len get args's length
func (ens Entrys) Len() int {
	return len(ens.args)
}

//Args (int) get the position of th args
func (ens Entrys) Args(i int) string {
	if i >= len(ens.args) {
		return ""
	}
	return ens.args[i]
}

//Flags (key string) return value of the key
func (ens Entrys) Flags(key string) interface{} {
	if v, exist := ens.flags[key]; exist {
		return v
	}
	return nil
}
