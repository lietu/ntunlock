package ntunlock

import (
	"os"
	"fmt"
)

func GetPlatformSettings() *Settings {
	s := Settings{[]string{}}
	s.SearchPaths = append(s.SearchPaths, fmt.Sprintf("%s/.config/nuclearthrone", os.Getenv("HOME")))
	s.SearchPaths = append(s.SearchPaths, fmt.Sprintf("%s/.steam/steam/%s", os.Getenv("HOME"), SteamAppPath))

	return &s
}
