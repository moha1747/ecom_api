package auth

import ("testing")

func TestCreateJW(t *testing.T) {
	secret := []byte("secret")

	token, err := CreateJWT(secret, 1)
	if err != nil {
		t.Errorf("Error creating JWT token: %v", err)
	}
	if token == "" {
		t.Error("Token is empty")
	}
}