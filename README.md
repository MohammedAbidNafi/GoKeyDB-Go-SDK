# GoKeyDB-Go-SDK

GoKeyDB is a lightweight, SQLite-backed key-value store library for Go applications. It provides a simple and efficient way to integrate persistent storage for managing key-value pairs. This library is ideal for small-scale applications or projects where setting up a full-fledged database is overkill.

## Features

- **Lightweight**: Minimal dependencies, only SQLite is required.
- **Simple API**: Easy-to-use methods for CRUD operations.
- **Thread-Safe**: Built-in synchronization ensures safe access in concurrent applications. (Coming soon)
- **In-Memory Caching**: Frequently accessed data is cached for faster lookups. (Coming soon)

## Installation

Install the library via `go get`:

```bash
go get github.com/MohammedAbidNafi/GoKeyDB-Go-SDK
```

## Getting Started

### Import the Library

```go
import "github.com/MohammedAbidNafi/GoKeyDB-Go-SDK"
```

### Example Usage

```go
package main

import (
	"log"
	"github.com/MohammedAbidNafi/GoKeyDB-Go-SDK"
)

func main() {
	// Initialize the database
	err := GoKeyDB.Initialize("mydatabase")
	if err != nil {
		log.Fatal(err)
	}

	// Add a key-value pair
	GoKeyDB.Put("mydatabase","username", "john_doe")

	// Retrieve a value
	value, exists := GoKeyDB.Get("mydatabase","username")
	if exists {
		log.Printf("Value: %s\n", value)
	}

	// List all key-value pairs
	GoKeyDB.List("mydatabase")

	// Delete a key-value pair
	GoKeyDB.Delete("mydatabase","username")
}
```

### API Reference

#### `Initialize(dbFileName string) error`
Initializes the SQLite database. Creates a file named `<dbFileName>.sqlite` if it doesn't already exist.

- **Parameters**: `dbFileName` - the name of the SQLite database file (without extension).
- **Returns**: An error if the initialization fails.

#### `Put(key string, value string) error`
Adds or updates a key-value pair in the database.

- **Parameters**:
  - `key`: The key to store.
  - `value`: The value associated with the key.
- **Returns**: An error if the operation fails.

#### `Get(dbname, key string) (string, bool)`
Retrieves the value associated with a key.

- **Parameters**:
- `dbname` - the name of the database (without extension)
- `key` - the key to retrieve.
- **Returns**:
  - `string`: The value associated with the key.
  - `bool`: `true` if the key exists, `false` otherwise.

#### `Delete(dbname, key string) error`
Deletes a key-value pair from the database.

- **Parameters**:
- `dbname` - the name of the database (without extension)
- `key` - the key to delete.
- **Returns**: An error if the operation fails.

#### `List(dbname string) ([]map[string]string, error)`
Prints all key-value pairs from the database and the in-memory cache.

- **Parameters**: `dbname` - the db name
- **Returns**:
- `array`: This contains all the keys that are available in the db
- `error`: An error if operation fails

## Advanced Features (Coming Soon)

### Query Feature
Ability to Query the value of the keys in the db

### In-Memory Caching
GoKeyDB uses an in-memory cache to speed up read operations. Any data retrieved from the database is automatically cached, reducing database queries for subsequent accesses.

### Thread Safety
All database operations are protected by a read-write mutex, ensuring safe usage in multi-threaded applications.

## Use Cases

- Storing user preferences or configuration settings.
- A simple key value database
- Managing lightweight key-value storage in CLI tools or small-scale projects.

## Dependencies

- [SQLite](https://www.sqlite.org/): A lightweight and embedded SQL database engine.
- [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3): SQLite driver for Go.

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a feature branch: `git checkout -b feat/feature-name`
3. Commit your changes: `git commit -m 'feat: Add feature'`
4. Push to the branch: `git push origin feat/feature-name`
5. Open a pull request.

## License

This project is licensed under the GPL-3.0 License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to the Go community for their support and libraries.
- Special thanks to the developers of SQLite and its Go driver for making lightweight database storage simple.

