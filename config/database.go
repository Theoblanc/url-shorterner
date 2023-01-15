package config

import "os"

// DatabaseConfigInterface database config interface
type DatabaseConfigInterface interface {
	Host() string
	Port() string
	Dbname() string
	User() string
	Password() string
	Sslmode() string
}

// Database database config struct
type Database struct {
	port     string
	host     string
	dbname   string
	user     string
	password string
	sslmode  string
}

// NewDatabaseConfig create database instance
func NewDatabaseConfig() *Database {
	port := "5432"
	host := "localhost"
	dbname := "admin"
	user := "admin"
	password := "test"
	sslmode := "disable"

	if env := os.Getenv("DB_PORT"); env != "" {
		port = env
	}
	if env := os.Getenv("DB_HOST"); env != "" {
		host = env
	}
	if env := os.Getenv("DB_NAME"); env != "" {
		dbname = env
	}
	if env := os.Getenv("DB_USER"); env != "" {
		user = env
	}
	if env := os.Getenv("DB_PASSWORD"); env != "" {
		password = env
	}
	if env := os.Getenv("DB_SSLMODE"); env != "" {
		sslmode = env
	}

	database := &Database{
		port:     port,
		host:     host,
		dbname:   dbname,
		user:     user,
		password: password,
		sslmode:  sslmode,
	}

	return database
}

// Host get database host
func (database *Database) Host() string {
	return database.host
}

// Port get database port number
func (database *Database) Port() string {
	return database.port
}

// Dbname get database name
func (database *Database) Dbname() string {
	return database.dbname
}

// User get databsae user name
func (database *Database) User() string {
	return database.user
}

// Password get database user password
func (database *Database) Password() string {
	return database.password
}

// Sslmode get database sslmode
func (database *Database) Sslmode() string {
	return database.sslmode
}
