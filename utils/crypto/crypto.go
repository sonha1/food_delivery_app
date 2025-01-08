package crypto

import (
	"encoding/hex"
	"github.com/sonha1/food_delivery_app/global"
	"golang.org/x/crypto/bcrypt"
)

func GetSaltRounds() int {
	if global.Config.Auth.SaltRounds > 10 {
		return global.Config.Auth.SaltRounds
	}

	return 10
}

func HashPassword(password string) (string, error) {
	salt := GetSaltRounds()
	hashPass, err := bcrypt.GenerateFromPassword(([]byte(password)), salt)
	if err != nil {
		//log
		return "", err
	}
	return hex.EncodeToString(hashPass[:]), nil
}

func MatchingPassword(storeHash string, password string) (bool, error) {
	hashPassword, err := HashPassword(password)
	if err != nil {
		//log
		return false, err
	}
	return storeHash == hashPassword, nil
}
