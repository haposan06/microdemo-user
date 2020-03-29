package users

import (
	"fmt"
	"testing"
)

func TestCreateUser(t *testing.T) {
	u := New()
	if len(u.Addresses) != 0 && len(u.Cards) != 0 {
		t.Error("Expected zero length addresses and cards")
	}
}

func TestValidateUser(t *testing.T) {
	u := New()
	err := u.Validate()
	if err.Error() != fmt.Sprintf(ErrMissingField, "FirstName") {
		t.Error("Expected missing first name error")
	}
	u.FirstName = "test"
	err = u.Validate()
	if err.Error() != fmt.Sprintf(ErrMissingField, "LastName") {
		t.Error("Expected missing last name error")
	}
	u.LastName = "test"
	err = u.Validate()
	if err.Error() != fmt.Sprintf(ErrMissingField, "Username") {
		t.Error("Expected missing username error")
	}
	u.Username = "test"
	err = u.Validate()
	if err.Error() != fmt.Sprintf(ErrMissingField, "Password") {
		t.Error("Expected missing password error")
	}
	u.Password = "test"
	err = u.Validate()
	if err != nil {
		t.Error(err)
	}
}

func TestMaskCC(t *testing.T) {
	u := New()
	u.Cards = append(u.Cards, Card{LongNum: "abcdefg"})
	u.Cards = append(u.Cards, Card{LongNum: "hijklomqrps"})
	u.maskCCs()
	if u.Cards[0].LongNum != "***defg" {
		t.Error("Masking not working on 1st card")
	}
	if u.Cards[1].LongNum != "*******qrps" {
		t.Error("Masking not working on 2st card")
	}
}
