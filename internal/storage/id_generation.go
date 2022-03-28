package storage

import (
	"github.com/AnimeKaizoku/pasty/internal/config"
	"github.com/AnimeKaizoku/pasty/internal/utils"
)

// AcquireID generates a new unique ID
func AcquireID() (string, error) {
	for {
		id := utils.RandomString(config.Current.IDCharacters, config.Current.IDLength)
		paste, err := Current.Get(id)
		if err != nil {
			return "", err
		}
		if paste == nil {
			return id, nil
		}
	}
}
