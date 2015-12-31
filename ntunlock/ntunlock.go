package ntunlock

import (
	"os"
	"log"
	"path/filepath"
	"fmt"
	"bufio"
	"encoding/json"
	"io/ioutil"
)

type Settings struct {
	SearchPaths []string
}

// Path Glob under Steam installation with the Steam Cloud files
var SteamAppPath string = "userdata/*/242680/remote"

func GetSettings() *Settings {
	s := GetPlatformSettings()

	for _, path := range os.Args[1:] {
		s.SearchPaths = append(s.SearchPaths, path)
	}

	return s
}

func Read(msg string) string {
	fmt.Print(msg)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}

func YesNo(msg string) string {
	for {
		answer := Read(fmt.Sprintf("%s (y/n) ", msg))

		if answer[:1] == "y" {
			return "y"
		} else if answer[:1] == "n" {
			return "n"
		} else {
			fmt.Println("Unrecognized answer, try y or n.")
		}
	}
}

func Run(s *Settings) {
	var files []string = []string{}

	for _, search := range s.SearchPaths {
		pattern := fmt.Sprintf("%s%cnuclearthrone.sav", search, os.PathSeparator)

		fmt.Println("Searching", pattern)

		matches, err := filepath.Glob(pattern)

		if err != nil {
			log.Fatalf("Got error %s when trying to check for files in %s", err, search)
		}

		for _, entry := range matches {
			files = append(files, entry)
		}
	}

	if len(files) == 0 {
		fmt.Println("")
		fmt.Println("Make sure you've run Nuclear Throne at least once before running this tool.")
		fmt.Println("Alternatively give more paths to search for nuclearthrone.sav as commandline arguments.")
		fmt.Println("")
		log.Fatalf("Failed to find any saves. No changes done.")
	}

	fmt.Println("")
	fmt.Println("Found the following saves:")
	for _, file := range files {
		fmt.Println(" -", file)
	}

	fmt.Println("")
	es := GetEditSettings()

	fmt.Println("")
	fmt.Println("Are you sure these are the changes you want to do?")
	fmt.Println("")
	fmt.Println("Understand that the save file is protected by Vlambeer DRM and you accept the terms by editing the save file:")
	fmt.Println("YUNG VENUZ WILL SHANK YOU IF YOU EDIT THIS FILE.")
	fmt.Println("")

	Read("Press <Enter> to continue or CTRL+C to abort.")

	fmt.Println("")

	for _, file := range files {
		s := NewSave()

		backup := fmt.Sprintf("%s.bak", file)
		fmt.Printf("Creating %s\n", backup)
		os.Link(file, backup)

		contents, err := ioutil.ReadFile(file)

		if err != nil {
			log.Fatalf("Error reading %s: %s", file, err)
		}

		json.Unmarshal(contents, &s)

		Edit(s, es)

		contents, err = json.Marshal(&s)

		if err != nil {
			log.Fatalf("Error regenerating %s: %s", file, err)
		}

		fmt.Printf("Writing %s\n", file)

		ioutil.WriteFile(file, contents, 0644)
	}

	fmt.Println("")
	fmt.Println("Changes done. Give it a go.")
}
