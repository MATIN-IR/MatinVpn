package auth

import "errors"

// نمونه‌ی ساده – بعداً به DB وصل می‌شه
var users = map[string]string{
	"admin": "matin123",
}

func Authenticate(username, password string) error {
	pass, ok := users[username]
	if !ok {
		return errors.New("user not found")
	}
	if pass != password {
		return errors.New("invalid password")
	}
	return nil
}
