package libs

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(hashPassword, dbPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(dbPassword))

	if err != nil {
		return err
	}

	return nil
}
