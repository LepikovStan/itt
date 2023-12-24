package tools

import "golang.org/x/crypto/bcrypt"

func hashPassword(password string) (string, error) {
	// Generate a bcrypt hash of the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// comparePassword compares a password with its hashed version
func comparePassword(password, hashedPassword string) error {
	// Compare the password with the hashed version
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
