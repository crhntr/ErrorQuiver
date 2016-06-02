package ErrorQuiver_test

import (
	"errors"
	"github.com/hunter-r-christopher/ErrorQuiver"
	"strings"
	"testing"
	"time"
)

var (
	attempt int = 0
)

func TestExample(t *testing.T) {
	user := User{Name: "christopher-r-hunter", Email: "example#example.com", Password: "BadPa$$", Foo: "baz"}

	// First attempt
	err := user.Validate()
	t.Log(err.Error())

	user.Name = "chris"
	user.Email = "example@example.com"
	user.Password = "$ti!!ABadPa$$w0rd"
	user.Foo = "bar"

	// Second attempt
	err = user.Validate()
	t.Log(err.Error())
	attempt++
}

type User struct {
	Name     string
	Password string
	Email    string
	Foo      string
}

func (user *User) Validate() error {
	errs := ErrorQuiver.New()

	if len(user.Password) < 8 {
		errs.AddNew("passwords must be greater than or equal to 8 characters long")
	}

	err := ValidateName(user.Name)
	errs.Add(err)

	// or if you prefer in one line
	errs.Add(ValidateEmail(user.Email))

	// since Add does a nil check this would be inefficient
	// prefer one of the above methods
	if err := ValidateFoo(user.Foo); err != nil {
		errs.Add(err)
	}

	// See below for easier example
	if errs.HaveBeenAdded() {
		return errs
	}

	// Do some resource expensive error checking
	// Such as checking the database for existing
	// user with same credentials

	errs.Add(user.expensiveValidation())

	// Returns nil if errs.Len() == 0
	return errs.Return()
}

func (user *User) expensiveValidation() error {
	time.Sleep(time.Second)
	return errors.New("database timeout")
}

func ValidateName(name string) error {
	if strings.ContainsAny(name, "-") {
		return errors.New("name has non-letter characters")
	}
	return nil
}

func ValidateEmail(email string) error {
	if !strings.ContainsAny(email, "@") {
		return errors.New("invalid email")
	}
	return nil
}

func ValidateFoo(foo string) error {
	if foo != "bar" {
		return errors.New("something baz happened")
	}
	return nil
}
