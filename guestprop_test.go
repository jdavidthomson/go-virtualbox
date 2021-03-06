package virtualbox

import (
	"errors"
	"fmt"
	"testing"
)

func TestGuestProperty(t *testing.T) {
	Setup(t)

	t.Logf("ManageMock=%v (type=%T)", ManageMock, ManageMock)
	if ManageMock != nil {
		ManageMock.EXPECT().run("guestproperty", "set", VM, "test_key", "test_val").Return(nil)
	}
	err := SetGuestProperty(VM, "test_key", "test_val")
	if err != nil {
		t.Fatal(err)
	}
	if Verbose {
		t.Logf("OK SetGuestProperty test_key=test_val")
	}

	if ManageMock != nil {
		ManageMock.EXPECT().runOut("guestproperty", "get", VM, "test_key").Return("Value: test_val", nil).Times(1)
	}
	val, err := GetGuestProperty(VM, "test_key")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("val='%s'", val)
	if val != "test_val" {
		t.Fatal("Wrong value")
	}
	if Verbose {
		t.Logf("OK GetGuestProperty test_key=test_val")
	}

	// Now deletes it...
	if ManageMock != nil {
		ManageMock.EXPECT().run("guestproperty", "delete", VM, "test_key").Return(nil).Times(1)
	}
	err = DeleteGuestProperty(VM, "test_key")
	if err != nil {
		t.Fatal(err)
	}
	if Verbose {
		t.Logf("OK DeleteGuestProperty test_key")
	}

	// ...and check that it is  no longer readable
	if ManageMock != nil {
		ManageMock.EXPECT().runOut("guestproperty", "get", VM, "test_key").Return("", errors.New("foo")).Times(1)
	}
	_, err = GetGuestProperty(VM, "test_key")
	if err == nil {
		t.Fatal(fmt.Errorf("Failed deleting guestproperty"))
	}
	if Verbose {
		t.Logf("OK GetGuestProperty test_key=empty")
	}

	Teardown()
}
