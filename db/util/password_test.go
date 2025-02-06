package util

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "mysecurepassword"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	// Ensure the hashed password is not empty
	if hashedPassword == "" {
		t.Fatal("Hashed password is empty")
	}

	// Check that the hashed password is not the same as the original password
	if hashedPassword == password {
		t.Fatal("Hashed password should not be the same as the original password")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "mysecurepassword"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	if !CheckPasswordHash(password, hashedPassword) {
		t.Fatal("Expected password to match hashed password")
	}

	wrongPassword := "wrongpassword"
	if CheckPasswordHash(wrongPassword, hashedPassword) {
		t.Fatal("Expected wrong password to not match hashed password")
	}
}
