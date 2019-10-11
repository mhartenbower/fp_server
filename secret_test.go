package main

import (
	"testing"
	"time"
)

func TestGetSecret(t *testing.T) {
	secret := GetSecret(1)

	if secret.Ciphertext == "" {
		t.Errorf("Did not get a secret back")
	}

	if secret.ID == 0 {
		t.Errorf("Did not get a secret ID back")
	}

	if secret.UserID == "" {
		t.Errorf("Did not get a user ID back")
	}

	if secret.CreatedAt.IsZero() {
		t.Errorf("Did not get a creation time back")
	}

	if secret.UpdatedAt.IsZero() {
		t.Errorf("Did not get an update time back")
	}
}

func TestCreateSecret(t *testing.T) {
	s := &Secret{
		123,
		"abcd",
		"abcde",
		time.Now(),
		time.Now(),
		time.Now(),
	}

	err := CreateSecret(s)
	if err != nil {
		t.Errorf("Failed to create secret: %s", err)
	}
}
