package customeauth

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/skp/app/models"
	"golang.org/x/crypto/bcrypt"
)

func New() *Customeauth {
	return &Customeauth{}
}

type Customeauth struct {
	// RedisClient *redisclient.RedisClient
}

func (customeAuth *Customeauth) CreateUserToken(usermail string, user *models.User) (string, string) {
	hashedtoken, _ := bcrypt.GenerateFromPassword([]byte(usermail), bcrypt.DefaultCost)
	token := string(hashedtoken)
	tokenlength := truncateText(token, 20)
	// fmt.Println("auth token", len(tokenlength))
	hash := sha256.Sum256([]byte(tokenlength))
	refrestoken := truncateText(hex.EncodeToString(hash[:]), 20)
	// fmt.Println("refresh token", len(refrestoken))
	return tokenlength, refrestoken
}
func truncateText(str string, max int) string {
	if max < 0 {
		return ""
	}
	r := []rune(str)
	trunc := r[:max]
	return string(trunc)
}
