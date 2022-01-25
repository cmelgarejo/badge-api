package data

import (
	"database/sql"
	"io/ioutil"
	"os"

	// database driver
	_ "github.com/lib/pq"
)

func getConnection() (*sql.DB, error) {
	uri := os.Getenv("DATABASE_URL")
	return sql.Open("postgres", uri)
}

// RunMigrations creates all the tables in the database
func RunMigrations(db *sql.DB) error {
	b, err := ioutil.ReadFile("./database/models.sql")
	if err != nil {
		return err
	}

	rows, err := db.Query(string(b))
	if err != nil {
		return err
	}

	return rows.Close()
}
