package domain

import (
	"database/sql"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Login struct {
	Username string         `db:"username"`
	PlayerId sql.NullString `db:"player_id"`
	TeamId   sql.NullString `db:"team_id"`
	Role     string         `db:"role"`
}

func (l Login) ClaimsForAccessToken() AccessTokenClaims {
	if l.PlayerId.Valid && l.TeamId.Valid {
		return l.claimsForUser()
	} else {
		return l.claimsForAdmin()
	}
}

func (l Login) claimsForUser() AccessTokenClaims {
	return AccessTokenClaims{
		PlayerId: l.PlayerId.String,
		TeamId:   l.TeamId.String,
		Username: l.Username,
		Role:     l.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ACCESS_TOKEN_DURATION).Unix(),
		},
	}
}

func (l Login) claimsForAdmin() AccessTokenClaims {
	return AccessTokenClaims{
		Username: l.Username,
		Role:     l.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ACCESS_TOKEN_DURATION).Unix(),
		},
	}
}
