package user

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

var (
	ErrInvalidEmail    = errors.New("invalid email format")
	ErrInvalidPassword = errors.New("password must be at least 6 characters")
)

type Email struct {
	value string
}

func NewEmail(email string) (Email, error) {
	if !regexp.MustCompile(`^[\w._%+\-]+@[\w.\-]+\.[a-zA-Z]{2,}$`).MatchString(email) {
		return Email{}, ErrInvalidEmail
	}
	return Email{value: email}, nil
}

func (e Email) String() string {
	return e.value
}

func (e Email) Value() (driver.Value, error) {
	return e.value, nil
}

func (e *Email) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to scan Email: %v", value)
	}

	valid, err := NewEmail(str)
	if err != nil {
		return err
	}

	*e = valid
	return nil
}

type Password struct {
	hashed string
}

func NewPassword(raw string) (Password, error) {
	if len(raw) < 6 {
		return Password{}, ErrInvalidPassword
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return Password{}, err
	}
	return Password{hashed: string(hashed)}, nil
}

func (p Password) String() string {
	return p.hashed
}

func (p Password) Compare(raw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p.hashed), []byte(raw)) == nil
}

func (p Password) Value() (driver.Value, error) {
	return p.hashed, nil
}

func (p *Password) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to scan Password: %v", value)
	}

	valid, err := NewPassword(str)
	if err != nil {
		return err
	}

	*p = valid
	return nil
}
