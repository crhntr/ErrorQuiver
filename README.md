# ErrorSlice

ErrorSlice is a lightweight package to make hadling multiple errors that need to be returned together easier.

Implements the error interface

**100% test coverage**

It is usefull when validating struct atrributes when you want to return a list of errors to the user.
## For Example 
See Example_test for full working test code which calls this function
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