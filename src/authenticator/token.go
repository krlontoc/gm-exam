package authenticator

import (
	"errors"
	"io/ioutil"
	"time"

	cfg "gm-exam/config"
	mdls "gm-exam/src/models"

	"github.com/golang-jwt/jwt/v4"
)

const (
	prvKey = "storage/keys/private_key.pem"
	pubKey = "storage/keys/public_key.pem"
)

type TokenClaim struct {
	Vld     bool      `json:"valid"`
	Session mdls.User `json:"session"`
	jwt.StandardClaims
}

func GenerateToken(session mdls.User) (string, error) {
	bPrvKey, err := ioutil.ReadFile(prvKey)
	if err != nil {
		return "", errors.New(cfg.InternalError)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(bPrvKey)
	if err != nil {
		return "", errors.New(cfg.InternalError)
	}

	claims := TokenClaim{
		true,
		session,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", errors.New(cfg.InternalError)
	}

	return tokenString, nil
}

func ValidateToken(authToken string) (bool, mdls.User, error) {
	bPubKey, err := ioutil.ReadFile(pubKey)
	if err != nil {
		return false, mdls.User{}, err
	}

	signingKey, err := jwt.ParseRSAPublicKeyFromPEM(bPubKey)
	if err != nil {
		return false, mdls.User{}, errors.New(cfg.InternalError)
	}

	token, err := jwt.ParseWithClaims(authToken, &TokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New(cfg.InvalidToken)
		}
		return signingKey, nil
	})
	if err != nil {
		return false, mdls.User{}, err
	}

	claims, ok := token.Claims.(*TokenClaim)
	if !ok || !token.Valid || !claims.Vld {
		return false, mdls.User{}, errors.New(cfg.InvalidToken)
	}

	// valid token
	return true, claims.Session, nil
}
