package definitions

import (
	"fmt"
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	cmds, err := LoadDefinitions("definitions.hcl")
	if err != nil {
		t.Fatalf("Could not load configuration file: %s", err)
	}
	if cmds[0].Subcommands[0] == "" {
		t.Fatalf("Subcommand value is empty")
	}
}

func TestHyphenToCamel(t *testing.T) {
	l := "list-zones"
	out := hyphenDelimToCamel(l)
	if out != "ListZones" {
		t.Fatalf("Unexpected camelcase for list-zones: %s", out)
	}
}

func TestRunTemplate(t *testing.T) {
	cmds, err := LoadDefinitions()
	if err != nil {
		t.Fatalf("Could not load configuration file: %s", err)
	}
	s, err := cmds[1].ToGo()
	if err != nil {
		t.Fatalf("Could not convert command to golang: %s", err)
	}
	fmt.Println(s)
	s, err = cmds[0].ToGo()
	if err != nil {
		t.Fatalf("Could not convert command to golang: %s", err)
	}
	fmt.Println(s)
}
