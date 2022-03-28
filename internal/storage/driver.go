package storage

import (
	"fmt"

	"github.com/AnimeKaizoku/pasty/internal/config"
	"github.com/AnimeKaizoku/pasty/internal/shared"
	"github.com/AnimeKaizoku/pasty/internal/storage/file"
	"github.com/AnimeKaizoku/pasty/internal/storage/mongodb"
	"github.com/AnimeKaizoku/pasty/internal/storage/postgres"
	"github.com/AnimeKaizoku/pasty/internal/storage/s3"
)

// Current holds the current storage driver
var Current Driver

// Driver represents a storage driver
type Driver interface {
	Initialize() error
	Terminate() error
	ListIDs() ([]string, error)
	Get(id string) (*shared.Paste, error)
	Save(paste *shared.Paste) error
	Delete(id string) error
	Cleanup() (int, error)
}

// Load loads the current storage driver
func Load() error {
	// Define the driver to use
	driver, err := GetDriver(config.Current.StorageType)
	if err != nil {
		return err
	}

	// Initialize the driver
	err = driver.Initialize()
	if err != nil {
		return err
	}
	Current = driver
	return nil
}

// GetDriver returns the driver with the given type if it exists
func GetDriver(storageType shared.StorageType) (Driver, error) {
	switch storageType {
	case shared.StorageTypeFile:
		return new(file.FileDriver), nil
	case shared.StorageTypePostgres:
		return new(postgres.PostgresDriver), nil
	case shared.StorageTypeMongoDB:
		return new(mongodb.MongoDBDriver), nil
	case shared.StorageTypeS3:
		return new(s3.S3Driver), nil
	default:
		return nil, fmt.Errorf("invalid storage type '%s'", storageType)
	}
}
