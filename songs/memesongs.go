package songs

import (
	"math/rand"
	"time"
)

// a MemeSongs implements Songs and holds meme songs
type MemeSongs struct {
	urls       []string
	urlsByName map[string]string
	randGen    *rand.Rand
}

// NewMemeSongs returns a new MemeSongs instance
func NewMemeSongs() *MemeSongs {
	return &MemeSongs{
		urls:       getSongsURLs(),
		urlsByName: getSongsURLsByName(),
		randGen:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// GetRandomSong returns a string that has a random meme song YT URL
func (m *MemeSongs) GetRandomSong() string {
	arrSize := len(m.urls)
	m.randGen.Seed(time.Now().UnixNano())

	return m.urls[m.randGen.Intn(arrSize)+0]
}

// GetSong returns a song with its name
func (m *MemeSongs) GetSong(key string) string {
	return m.urlsByName[key]
}

// getSongsURLs returns a string array of meme songs YT URLs
func getSongsURLs() []string {
	return []string{
		"https://www.youtube.com/watch?v=PfYnvDL0Qcw",
		"https://www.youtube.com/watch?v=L_jWHffIx5E",
		"https://www.youtube.com/watch?v=j9V78UbdzWI",
		"https://www.youtube.com/watch?v=GK2GUxOnjDQ",
		"https://www.youtube.com/watch?v=dQw4w9WgXcQ",
		"https://www.youtube.com/watch?v=jmPkNcyNINk",
		"https://youtu.be/CBrWNbjw3RA?t=92",
		"https://www.youtube.com/watch?v=HEXWRTEbj1I",
		"https://www.youtube.com/watch?v=2MtOpB5LlUA",
		"https://www.youtube.com/watch?v=1mrGdGMNsv0",
		"https://www.youtube.com/watch?v=SdcGu3qW8QU",
		"https://www.youtube.com/watch?v=Je5ISw6bqEo",
		"https://www.youtube.com/watch?v=5QCaaAyz-yA",
		"https://www.youtube.com/watch?v=wDgQdr8ZkTw",
		"https://www.youtube.com/watch?v=NUYvbT6vTPs",
		"https://www.youtube.com/watch?v=7yh9i0PAjck",
		"https://www.youtube.com/watch?v=feA64wXhbjo",
		"https://www.youtube.com/watch?v=QH2-TGUlwu4",
		"https://www.youtube.com/watch?v=cE0wfjsybIQ",
		"https://www.youtube.com/watch?v=jofNR_WkoCE",
		"https://www.youtube.com/watch?v=3WSgJCYIewM",
		"https://www.youtube.com/watch?v=x88Z5txBc7w",
		"https://www.youtube.com/watch?v=9W6AN_eQeZo",
	}
}

// getSongsURLsByName returns a map with meme songs YT URLs by names
// eg songs["ussr-anthem"] will give a YT URL with the USSR anthem
func getSongsURLsByName() map[string]string {
	return map[string]string{
		"we-are-number-one":      "https://www.youtube.com/watch?v=PfYnvDL0Qcw",
		"all-stars":              "https://www.youtube.com/watch?v=L_jWHffIx5E",
		"coffin-dance":           "https://www.youtube.com/watch?v=j9V78UbdzWI",
		"ussr-anthem":            "https://www.youtube.com/watch?v=GK2GUxOnjDQ",
		"rick-roll":              "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
		"sax-guy":                "https://www.youtube.com/watch?v=jmPkNcyNINk",
		"smooth-criminal":        "https://youtu.be/CBrWNbjw3RA?t=92",
		"what-is-love":           "https://www.youtube.com/watch?v=HEXWRTEbj1I",
		"giorono-theme":          "https://www.youtube.com/watch?v=2MtOpB5LlUA",
		"spooky-scary-skeletons": "https://www.youtube.com/watch?v=1mrGdGMNsv0",
		"creeper":                "https://www.youtube.com/watch?v=SdcGu3qW8QU",
		"lwiay":                  "https://www.youtube.com/watch?v=Je5ISw6bqEo",
		"smoke-weed-everyday":    "https://www.youtube.com/watch?v=5QCaaAyz-yA",
		"megalovania":            "https://www.youtube.com/watch?v=wDgQdr8ZkTw",
		"cat-ievan-polka":        "https://www.youtube.com/watch?v=NUYvbT6vTPs",
		"ievan-polka":            "https://www.youtube.com/watch?v=7yh9i0PAjck",
		"shooting-stars":         "https://www.youtube.com/watch?v=feA64wXhbjo",
		"nyan-cat":               "https://www.youtube.com/watch?v=QH2-TGUlwu4",
		"crab-rave":              "https://www.youtube.com/watch?v=cE0wfjsybIQ",
		"what-does-the-fox-say":  "https://www.youtube.com/watch?v=jofNR_WkoCE",
		"in-my-feelings":         "https://www.youtube.com/watch?v=3WSgJCYIewM",
		"yakkos-world":           "https://www.youtube.com/watch?v=x88Z5txBc7w",
		"you-got-that":           "https://www.youtube.com/watch?v=9W6AN_eQeZo",
	}
}
