package persistence

import (
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func dbPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".quickrelay.db"), nil
}

func OpenConnection() (*sqlx.DB, error) {
	path, err := dbPath()
	if err != nil {
		return nil, err
	}

	return sqlx.Open("sqlite3", path)
}
