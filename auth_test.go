package sr_auth

import "testing"

var auth Auth = Auth{Key: "Test", AuthServerAddress: "http://localhost:8080"}
var userTest User = User{
	Login:     "user",
	FirstName: "user",
	LastName:  "user",
	Email:     "user@user",
}

func TestGetUserFromToken(t *testing.T) {
	username, err := auth.GetUsernameFromToken(token)
	if err != nil {
		t.Fatal(err)
	}

	if username != "test" {
		t.Fatalf("input != output")
	}
}
