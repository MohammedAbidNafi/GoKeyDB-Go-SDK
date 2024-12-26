package GoKeyDB

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	memory = make(map[string]string)
	db     *sql.DB
	mu     sync.RWMutex
)

// Initialize initializes the SQLite database
func Initialize(dbFileName string) error {
	var err error
	db, err = sql.Open("sqlite3", dbFileName+".sqlite")
	if err != nil {
		log.Printf("Failed to open database: %v", err)
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS `+dbFileName+` (
		key TEXT PRIMARY KEY,
		value TEXT
	)`, dbFileName)
	if err != nil {
		log.Printf("Failed to create table: %v", err)
		return err
	}

	log.Println("Database initialized successfully")
	return nil
}

// Put adds or updates a key-value pair
func Put(dbname, key, value string) error {
	mu.Lock()
	defer mu.Unlock()
	var err error
	db, err = sql.Open("sqlite3", dbname+".sqlite")
	if err != nil {
		log.Printf("Failed to open database: %v", err)
		return err
	}
	if db == nil {
		return fmt.Errorf("database not initialized")
	}

	memory[key] = value
	_, errfromdb := db.Exec(`INSERT INTO `+dbname+` (key, value) VALUES (?, ?)
		ON CONFLICT(key) DO UPDATE SET value=excluded.value`, key, value)
	if errfromdb != nil {
		log.Printf("Failed to insert key-value pair: %v", errfromdb)
	}
	return errfromdb
}

// Get retrieves the value for a key
func Get(dbname, key string) (string, bool) {
	mu.RLock()
	defer mu.RUnlock()

	var err error
	db, err = sql.Open("sqlite3", dbname+".sqlite")
	if err != nil {
		log.Printf("Failed to open database: %v", err)
		return "", false
	}
	if db == nil {
		log.Printf("Database not initialized")

	}

	value, exists := memory[key]
	if !exists {
		row := db.QueryRow("SELECT value FROM "+dbname+" WHERE key = ?", key)
		err := row.Scan(&value)
		if err != nil {
			if err != sql.ErrNoRows {
				log.Printf("Failed to retrieve key-value pair: %v", err)
			}
			return "", false
		}
		memory[key] = value
		exists = true
	}
	return value, exists
}

// Delete removes a key-value pair
func Delete(dbname, key string) error {
	mu.Lock()
	defer mu.Unlock()

	var err error
	db, err = sql.Open("sqlite3", dbname+".sqlite")
	if err != nil {
		log.Printf("Failed to open database: %v", err)

	}
	if db == nil {
		log.Printf("Database not initialized")

	}

	delete(memory, key)
	_, err = db.Exec("DELETE FROM "+dbname+" WHERE key = ?", key)
	if err != nil {
		log.Printf("Failed to delete key-value pair: %v", err)
	}
	return err
}

// List prints all key-value pairs
func List(dbname string) {
	mu.RLock()
	defer mu.RUnlock()
	var err error
	db, err = sql.Open("sqlite3", dbname+".sqlite")
	if err != nil {
		log.Printf("Failed to open database: %v", err)

	}
	if db == nil {
		log.Printf("Database not initialized")

	}

	for key, value := range memory {
		fmt.Printf("%s: %s\n", key, value)
	}
	rows, err := db.Query("SELECT key, value FROM " + dbname)
	if err != nil {
		log.Printf("Failed to query database: %v", err)
	}
	// Display all the keys
	for rows.Next() {
		var key, value string
		if err := rows.Scan(&key, &value); err != nil {
			log.Printf("Failed to scan row: %v", err)
		}
		fmt.Printf("%s: %s\n", key, value)
	}

}
