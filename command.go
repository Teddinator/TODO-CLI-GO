package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    string
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo, specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & specify a new title. id:new_title")
	flag.StringVar(&cf.Del, "del", "", "Specify todo indexes to delete e.g. 1 or 1,2,3")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a todo by index to toggle")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()

	return &cf
}

func parseIndexes(value string) ([]int, error) {
	parts := strings.Split(value, ",")
	indexes := make([]int, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)

		index, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}

		indexes = append(indexes, index)
	}

	return indexes, nil
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}

		todos.edit(index, parts[1])

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.Del != "":
		indexes, err := parseIndexes(cf.Del)
		if err != nil {
			fmt.Println("Error: Invalid index for delete")
			os.Exit(1)
		}

		if err := todos.delete(indexes...); err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}

	default:
		fmt.Println("Invalid command")
	}
}
