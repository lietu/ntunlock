package ntunlock

import (
	"os"
	"fmt"
	"strings"
)

func GetPlatformSettings() *Settings {
	s := Settings{[]string{}}

	s.SearchPaths = append(s.SearchPaths, fmt.Sprintf("%s/nuclearthrone", os.Getenv("LOCALAPPDATA")))
	s.SearchPaths = append(s.SearchPaths, fmt.Sprintf("%s/Steam/%s", os.Getenv("ProgramFiles"), SteamAppPath))

	pf86 := os.Getenv("ProgramFiles(x86)")
	if pf86 != "" {
		s.SearchPaths = append(s.SearchPaths, fmt.Sprintf("%s/Steam/%s", pf86, SteamAppPath))
	}

	paths := s.SearchPaths
	s.SearchPaths = []string{}
	for _, v := range paths {
		s.SearchPaths = append(s.SearchPaths, strings.Replace(v, "/", "\\", -1))
	}

	return &s
}
