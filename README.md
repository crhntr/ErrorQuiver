# ErrorQuiver

ErrorQuiver is a convienient and lightweight package wrapping a slice of errors.

Implements the error interface

Although this package is fun and has earned it's place in a few of my projects, check out [pkg/errors](https://github.com/pkg/errors).

**93.8% test coverage**

## Example 
See [this test](https://github.com/hunter-r-christopher/ErrorQuiver/blob/master/Example_test.go) for full working example test code which calls this function
```go

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
```
