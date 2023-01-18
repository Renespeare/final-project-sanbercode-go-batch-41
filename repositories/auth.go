package repositories

import (
	"database/sql"
	"errors"
	"final-project/middleware"
	userModel "final-project/models/user"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Register(db *sql.DB, user userModel.Register) (err error) {
	sqlStatement := "INSERT INTO users (name, email, password, created_at) VALUES ($1, $2, $3, $4)"

	createdAt := time.Now()
	hash, _ := HashPassword(user.Password)
	errs := db.QueryRow(sqlStatement, user.Name, user.Email, hash, createdAt)
	return errs.Err()
}

func Login(db *sql.DB, user userModel.Login) (error, interface{}) {
	sqlStatement := "SELECT id, email, password FROM users WHERE email = $1"
	row := db.QueryRow(sqlStatement, user.Email)
	
	storedUser := struct {
		ID int
		Email string
		Password string
	}{}

	err := row.Scan(&storedUser.ID, &storedUser.Email, &storedUser.Password)
	if err != nil {
		return err, nil
	}

	match := CheckPasswordHash(user.Password, storedUser.Password)
	if !match {
		return errors.New("password is incorrect"), nil
	}

	sqlStatement = "DELETE FROM user_credentials WHERE user_id = $1"
	db.QueryRow(sqlStatement, storedUser.ID)
	

	uuid := uuid.New()

	sqlStatement = "INSERT INTO user_credentials (user_id, uuid) VALUES ($1, $2)"
	insertRow := db.QueryRow(sqlStatement, storedUser.ID, uuid)
	if insertRow.Err() != nil {
		return errors.New("failed to insert user credentials"), nil
	}

	token, err := middleware.GenerateToken(storedUser.ID, uuid.String())
	if err != nil {
		return errors.New("failed to login"), nil
	}	

	return nil, token
}

func Logout(db *sql.DB, uuid string) (error) {
	sqlStatement := "DELETE FROM user_credentials WHERE uuid = $1"
	err := db.QueryRow(sqlStatement, uuid)

	return err.Err()
}
