package models

import (
	u "github.com/aliwert/go-jwt-example/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

type Token struct {
	UserId   uint
	Username string
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

// Incoming data validation function

func (account *Account) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email is incorrect"), false
	}
	if len(account.Password) < 8 {
		return u.Message(false, "Password is too short"), false
	}
	temp := &Account{}

	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "A connection error occurred. Please try again"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email is already taken"), false
	}
	return u.Message(false, "Requirement passed"), true
}

// User account creation function
func (account *Account) Create() map[string]interface{} {
	if resp, ok := account.Validate(); !ok {
		return resp
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)
	GetDB().Create(account)

	if account.ID <= 0 {
		return u.Message(false, "Failed to create account")
	}
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
}
