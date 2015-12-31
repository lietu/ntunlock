package ntunlock

import (
	"os"
	"fmt"
)

func GetPlatformSettings() *Settings {
	s := Settings{[]string{}}
	s.SearchPaths = append(s.SearchPaths, fmt.Sprintf("%s/Library/Application Support/com.vlambeer.nuclearthrone", os.Getenv("HOME")))
	s.SearchPaths = append(s.SearchPaths, fmt.Sprintf("%s/Steam/%s", os.Getenv("HOME"), SteamAppPath))

	return &s
}
