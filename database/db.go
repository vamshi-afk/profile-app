package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

type ProfileData struct {
	Username string
	Name     string
	Email    string
	Bio      string
	Hobbies  string
	Friends  string
	Success  string
}

func Init() error {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	DB = db

	if err := db.Ping(); err != nil {
		return err
	}

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		name TEXT,
		email TEXT,
		bio TEXT,
		hobbies TEXT,
		friends TEXT
	);`

	if _, err := db.Exec(query); err != nil {
		return err
	}

	return nil
}

func InsertUser(username string, password string) error {
	_, err := DB.Exec(`
		INSERT INTO users (username, password, name, email, bio, hobbies, friends)
		VALUES ($1, $2, '', '', '', '', '')`, username, password)
	return err
}

func GetHashedPassword(username string) (string, error) {
	var hp string
	err := DB.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&hp)
	return hp, err
}

func GetProfile(username string) (ProfileData, error) {
	var p ProfileData
	err := DB.QueryRow(`
		SELECT username,
		       COALESCE(name, ''),
		       COALESCE(email, ''),
		       COALESCE(bio, ''),
		       COALESCE(hobbies, ''),
		       COALESCE(friends, '')
		FROM users WHERE username = $1`, username).Scan(
		&p.Username, &p.Name, &p.Email, &p.Bio, &p.Hobbies, &p.Friends,
	)
	return p, err
}

func UpdateProfile(username, name, email, bio, hobbies, friends string) error {
	_, err := DB.Exec(`
		UPDATE users 
		SET name = $1, email = $2, bio = $3, hobbies = $4, friends = $5 
		WHERE username = $6`,
		name, email, bio, hobbies, friends, username,
	)
	return err
}
