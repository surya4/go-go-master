package models

import (
	"fmt"
	"log"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbHost = "tcp(127.0.0.1:3306)"
	dbName = "users_srv"
	dbUser = "root"
	dbPass = "root"
)

// Connection function
func Connection() (db *sql.DB) {
	fmt.Println("foo Connection model")
	db, err := sql.Open("mysql", fmt.Sprintf("%s/%s/%s", dbName, dbUser, dbPass))
	fmt.Println("err ->", err)
	fmt.Println("db ->", db)

	if err != nil {
		log.Fatal(err)
	}

	return db
	// defer db.Close()

	// return db
}

// User struct
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// don't import this value
	CreatedAt time.Time ``
	UpdatedAt time.Time `json:"updated_at"`
}

// Insert create update user
// func (u *User) Insert() error {
// 	u.CreatedAt = time.Now()
// 	u.UpdatedAt = time.Now()

// 	db := Connection()

// 	_, err := db.Prepare("INSERT INTO users (user_id, name, created_at, updated_at) VALUES (?,?,?,?)",
// 		u.UserID, u.Name, u.CreatedAt, u.UpdatedAt)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	return nil
// }

// GetUsers user
func (u *User) GetUsers(userID int) error {

	fmt.Println("foo GetUsers model")

	db := Connection()
	defer db.Close()

	fmt.Println("hello userID -->", userID)

	users := db.QueryRow("SELECT user_id, name from users where user_id = ?", userID).Scan(&u.ID, &u.Name)

	fmt.Println("hello users -->", users)

	// return err
	// if err != nil {
	// 	return err
	// }
	return users
}
