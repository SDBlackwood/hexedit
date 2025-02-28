package main

import (
	"os"
	"testing"
)

func TestMainFunc(t *testing.T) {
	tests := []struct {
		name string // description of this test case
	}{
    {
      name: "it_can_read_an_extant_file",
    },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
      // Run the comand on the command line
      os.Args = append(os.Args, "-filepath=../test/data/basic.txt")
      
      main()
		})
	}
}

