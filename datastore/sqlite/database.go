package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // driver
	"log"
	"os"
	"path"
)

const (
	driver     = "sqlite3"
	dataSource = "fabrica.db"
)

// DB local database with our custom methods.
type DB struct {
	*sql.DB
}

// NewDatabase returns an open database connection
func NewDatabase() (*DB, error) {
	// Open the database connection
	log.Println("Open database:", GetPath(dataSource))
	db, err := sql.Open(driver, GetPath(dataSource))
	if err != nil {
		log.Fatalf("Error opening the database: %v\n", err)
	}

	// Check that we have a valid database connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error accessing the database: %v\n", err)
	}

	store := &DB{db}
	store.CreateTables()

	return store, nil
}

// CreateTables creates the database tables
func (db *DB) CreateTables() error {
	if _, err := db.Exec(createRepoTableSQL); err != nil {
		return err
	}
	if _, err := db.Exec(createBuildTableSQL); err != nil {
		return err
	}
	if _, err := db.Exec(createBuildLogTableSQL); err != nil {
		return err
	}
	if _, err := db.Exec(createSettingsTableSQL); err != nil {
		return err
	}
	if _, err := db.Exec(createKeysTableSQL); err != nil {
		return err
	}
	_, _ = db.Exec(alterRepoTableSQL)
	_, _ = db.Exec(alterBuildTableSQL)
	_, _ = db.Exec(alterBuildTableSQLcontainer)
	_, _ = db.Exec(alterRepoTableKeySQL)
	return nil
}

// GetPath returns the full path to the data file
func GetPath(filename string) string {
	if len(os.Getenv("SNAP_DATA")) > 0 {
		return path.Join(os.Getenv("SNAP_DATA"), "../current", filename)
	}
	return filename
}
