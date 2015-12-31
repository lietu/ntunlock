package main

import (
	"fmt"
	"github.com/lietu/ntunlock/ntunlock"
)

func main() {
	s := ntunlock.GetSettings()

	fmt.Println("Welcome to NTUnlock!")
	fmt.Println("")
	fmt.Println("This tool will look for your Nuclear Throne save file + the Steam Cloud copy")
	fmt.Println("and can be used to unlock e.g. all characters, and all crowns.")
	fmt.Println("")
	fmt.Println("Will search for Nuclear Throne (nuclearthrone.sav) in the following directories:")

	for _, value := range s.SearchPaths {
		fmt.Println(" -", value)
	}

	fmt.Println("")
	fmt.Println("If you need to add more paths to scan, just give them as command line arguments.")
	fmt.Println("")

	ntunlock.Run(s)
}
