package useless

import (
	"github.com/baraa-almasri/useless/songs"
	"os/exec"
	"runtime"
)

// a YTPlayer plays youtube songs using the system's default browser
type YTPlayer struct {
	songSource songs.Songs
}

// NewYTPlayer returns a new YTPlayer instance
func NewYTPlayer(source songs.Songs) *YTPlayer {
	return &YTPlayer{songSource: source}
}

// PlayRandomSong plays a random song from the given music source
func (yt *YTPlayer) PlayRandomSong() error {
	return yt.PlaySong(yt.songSource.GetRandomSong())
}

// PlaySong launches default browser and plays the given url
func (yt *YTPlayer) PlaySong(song string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows": // fuck this shit :)
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin": // macOS :)
		cmd = "open"
	default: // linux and the other boys
		cmd = "xdg-open"
	}

	args = append(args, song)

	return exec.Command(cmd, args...).Start()
}
