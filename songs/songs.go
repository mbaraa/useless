package songs

// Songs interface represents a song list of a certain type
type Songs interface {
	GetRandomSong() string
	GetSong(key string) string
}
