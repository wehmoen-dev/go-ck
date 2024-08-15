package test

import (
	"testing"
)

const TestUserId = "bcfce3497b42e48f1210823471c1312f"

func TestGetUser(t *testing.T) {
	user, err := c.GetUser(TestUserId, false)
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}

	if user == nil {
		t.Fatalf("User is nil")
	}

	if user.Profile != nil {
		t.Fatalf("Profile should be nil")
	}
}

func TestGetUserWithProfile(t *testing.T) {
	user, err := c.GetUser(TestUserId, true)
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}

	if user == nil {
		t.Fatalf("User is nil")
	}

	if user.Id != TestUserId {
		t.Fatalf("User id does not match")
	}

	if user.Profile == nil {
		t.Fatalf("Profile should not be nil")
	}
}
