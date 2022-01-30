package users

import (
	"crypto/md5"
	"encoding/hex"
)

type RegisterUser struct {
	LoginInfo   User        `json:"login_info"`
	ProfileInfo UserProfile `json:"user_profile"`
}

func HashPassword(str string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
