package auth 

import (
	"testing"
)

func TestHashPassword(t *testing.T){
	hash, err  := HashPassword("password")
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}
	if hash == "" {
		t.Error("Hash is empty")
	}
	if hash == "password" {
		t.Error("expected hash to be different from password")
	}
}

func TestComparePasswords(t *testing.T){
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}
	if !ComparePasswords(hash, []byte("password")) {
		t.Error("expected password to match hash")
	}
	if ComparePasswords(hash, []byte("wrongpassword")) {
		t.Error("expected password to not match hash")
	}
}