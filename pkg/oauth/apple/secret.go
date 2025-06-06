package apple

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

/*
GenerateClientSecret generates the client secret used to make requests to the validation server.
The secret expires after 6 months

signingKey - Private key from Apple obtained by going to the keys section of the developer section
teamID - Your 10-character Team ID
clientID - Your Services ID, e.g. com.aaronparecki.services
keyID - Find the 10-char Key ID value from the portal
*/
func GenerateClientSecret(signingKey, teamID, clientID, keyID string) (string, error) {
	block, _ := pem.Decode([]byte(signingKey))
	if block == nil {
		return "", errors.New("empty block after decoding")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	// Create the Claims
	now := time.Now()
	claims := &jwt.RegisteredClaims{
		Issuer: teamID,
		IssuedAt: &jwt.NumericDate{
			Time: now,
		},
		ExpiresAt: &jwt.NumericDate{
			Time: now.Add(time.Hour*24*180 - time.Second), // 180 days
		},
		Audience: jwt.ClaimStrings{
			"https://appleid.apple.com",
		},
		Subject: clientID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header["alg"] = "ES256"
	token.Header["kid"] = keyID

	return token.SignedString(privateKey)
}
