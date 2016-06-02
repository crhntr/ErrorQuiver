package ErrorQuiver_test

import (
	"errors"
	"github.com/hunter-r-christopher/ErrorQuiver"
	"testing"
)

func TestNewErrorQuiver(t *testing.T) {
	errs := ErrorQuiver.New()

	if errs == nil {
		t.Error("NewErrorQuiver() should return a non nil pointer")
	}

	if errs.Len() != 0 {
		t.Error("NewErrorQuiver() should have a length of 0")
	}
}

func TestErrorQuiver_Add(t *testing.T) {
	errs := ErrorQuiver.New()

	err := errors.New("something went wrong")
	errs.Add(err)

	if errs.Len() != 1 {
		t.Error("after Add(err) && err != nil, the len of an ErrorQuiver should be 1")
	}
}

func TestErrorQuiver_Add2(t *testing.T) {
	errs := ErrorQuiver.New()

	errs.Add(nil)

	if errs.Len() != 0 {
		t.Error("after Add(nil), the len of an ErrorQuiver should be 0")
	}
}

func TestErrorQuiver_AddNew(t *testing.T) {
	errs := ErrorQuiver.New()

	errs.AddNew("something went wrong")

	if errs.Len() != 1 {
		t.Error("after AddNew(\"some error message\"), the len of an ErrorQuiver should be 1")
	}
}

func TestErrorQuiver_AddNew2(t *testing.T) {
	errs := ErrorQuiver.New()

	errs.AddNew("")

	if errs.Len() != 0 {
		t.Error("after AddNew(\"\"), the len of an ErrorQuiver should be 0")
	}
}

func TestErrorQuiver_Error(t *testing.T) {
	errs := ErrorQuiver.New()

	err1 := "err msg 1"
	err2 := "err msg 2"
	err3 := "err msg 3"
	err4 := "err msg 4"
	err5 := "err msg 5"
	err6 := "err msg 6"

	errs.AddNew(err1)
	errs.AddNew(err2)
	errs.AddNew(err3)
	errs.AddNew(err4)
	errs.AddNew(err5)
	errs.AddNew(err6)

	str := errs.Error()

	if str == "" {
		t.Error("Error() should return a non empty string")
	}
}

func TestErrorQuiver_Error2(t *testing.T) {
	errsN := ErrorQuiver.New()
	errsL := ErrorQuiver.New()

	err1 := "err msg 1"
	err2 := "err msg 2"
	err3 := "err msg 3"
	errA := "err msg A"
	errB := "err msg B"
	errC := "err msg C"

	errsN.AddNew(err1)
	errsN.AddNew(err2)
	errsN.AddNew(err3)
	errsL.AddNew(errA)
	errsL.AddNew(errB)
	errsL.AddNew(errC)
	errsN.Add(errsL)

	str := func() string {
		defer func() {
			if r := recover(); r != nil {
				t.Error("Error() should not panic or ")
			}
		}()

		return errsN.Error()
	}()

	if str == "" {
		t.Error("Error() should return a non empty string")
	}
}

func TestErrorQuiver_HaveBeenAdded(t *testing.T) {
	errs := ErrorQuiver.New()

	if errs.HaveBeenAdded() {
		t.Error("HaveBeenAdded() should return false if no errors have been added")
	}

	errs.AddNew("err msg 1")

	if !errs.HaveBeenAdded() {
		t.Error("HaveBeenAdded() should return true an error have been added")
	}

	errs.AddNew("err msg 2")
	errs.AddNew("err msg 3")
	errs.AddNew("err msg 4")
	errs.AddNew("err msg 5")
	errs.AddNew("err msg 6")

	if !errs.HaveBeenAdded() {
		t.Error("HaveBeenAdded() should return true serror have been added")
	}
}
