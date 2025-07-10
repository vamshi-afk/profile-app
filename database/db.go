package database

import "database/sql"
import _ "github.com/lib/pq"
import "os"

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
	db, err := sql.Open("sqlite3", connStr)
	if err != nil {
		return err
	}
	DB = db

	if err := db.Ping(); err != nil {
		return err
	}

	{
		query := `
		CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
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
	}
	return nil
}

func InsertUser(username string, password string) error {
	_, err := DB.Exec(`
		INSERT INTO users (username, password, name, email, bio, hobbies, friends)
		VALUES (?, ?, '', '', '', '', '')`, username, password)
	if err != nil {
		return err
	}
	return nil
}

func GetHashedPassword(username string) (string, error) {
	var hp string
	err := DB.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&hp)
	return hp, err
}

func GetProfile(username string) (ProfileData, error) {
	var p ProfileData
	err := DB.QueryRow(`
		SELECT username,
		       IFNULL(name, ''),
		       IFNULL(email, ''),
		       IFNULL(bio, ''),
		       IFNULL(hobbies, ''),
		       IFNULL(friends, '')
		FROM users WHERE username = ?`, username).Scan(
		&p.Username, &p.Name, &p.Email, &p.Bio, &p.Hobbies, &p.Friends,
	)
	return p, err
}

func UpdateProfile(username, name, email, bio, hobbies, friends string) error {
	_, err := DB.Exec("UPDATE users SET name = ?, email = ?, bio = ?, hobbies = ?, friends = ? WHERE username = ?", name, email, bio, hobbies, friends, username)
	if err != nil {
		return err
	}
	return nil
}
