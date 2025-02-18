package sercurity

import (
	"fmt"
	usersdto "github.com/MinhSang97/Server-Game/usecases/dto/users_dto"
	"github.com/golang-jwt/jwt"
	"sync"
	"time"
)

const SECRET_KEY_USERS = "userssecretkeylear"

type JwtCustomClaimsUsers struct {
	UserId string
	Role   string
	jwt.StandardClaims
}

var tokenStoreUsers = struct {
	sync.RWMutex
	tokens map[string]string
	expiry map[string]int64
}{tokens: make(map[string]string), expiry: make(map[string]int64)}

func GenTokenUsers(user usersdto.Users) (string, error) {
	tokenStoreUsers.RLock()
	existingToken, exists := tokenStoreUsers.tokens[user.UserId]
	expiryTime, _ := tokenStoreUsers.expiry[user.UserId]
	tokenStoreUsers.RUnlock()

	if exists && time.Now().Unix() < expiryTime {
		return existingToken, nil
	}

	claims := &JwtCustomClaimsUsers{
		UserId: user.UserId,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SECRET_KEY_USERS))
	if err != nil {
		fmt.Println("Lỗi tạo token:", err.Error())
		return "Tạo Token Lỗi!", err
	}

	tokenStoreUsers.Lock()
	tokenStoreUsers.tokens[user.UserId] = result
	tokenStoreUsers.expiry[user.UserId] = claims.ExpiresAt
	tokenStoreUsers.Unlock()

	return result, nil
}
