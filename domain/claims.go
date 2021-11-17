package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const HMAC_SAMPLE_SECRET = "hmacSampleSecret"
const ACCESS_TOKEN_DURATION = time.Hour
const REFRESH_TOKEN_DURATION = time.Hour * 24 * 30

type RefreshTokenClaims struct {
	TokenType string   `json:"token_type"`
	PlayerId  string   `json:"pid"`
	PlayerIds []string `json:"players"`
	TeameIds  []string `json:"teams"`
	TeamId    string   `json:"tid"`
	Username  string   `json:"un"`
	Role      string   `json:"role"`
	jwt.StandardClaims
}

type AccessTokenClaims struct {
	PlayerId  string   `json:"player_id"`
	TeamId    string   `json:"team_id"`
	PlayerIds []string `json:"players"`
	TeameIds  []string `json:"teams"`
	Username  string   `json:"username"`
	Role      string   `json:"role"`
	jwt.StandardClaims
}

func (c AccessTokenClaims) IsUserRole() bool {
	return c.Role == "user"
}

func (a AccessTokenClaims) IsValidPlayerId(playerId string) bool {
	return a.PlayerId == playerId
}

func (a AccessTokenClaims) IsValidTeamId(teamId string) bool {
	if teamId != "" {
		teamFound := false
		for _, a := range a.TeameIds {
			if a == teamId {
				teamFound = true
				break
			}
		}
		return teamFound
	}
	return true
}

func (a AccessTokenClaims) IsValidPlayerIds(playerId string) bool {
	if playerId != "" {
		playerFound := false
		for _, a := range a.PlayerIds {
			if a == playerId {
				playerFound = true
				break
			}
		}
		return playerFound
	}
	return true
}

func (c AccessTokenClaims) IsRequestVerifiedWithTokenClaims(urlParams map[string]string) bool {
	if c.PlayerId != urlParams["player_id"] {
		return false
	}

	if !c.IsValidTeamId(urlParams["team_id"]) {
		return false
	}
	return true
}

func (c AccessTokenClaims) RefreshTokenClaims() RefreshTokenClaims {
	return RefreshTokenClaims{
		TokenType: "refresh_token",
		PlayerId:  c.PlayerId,
		TeamId:    c.TeamId,
		Username:  c.Username,
		Role:      c.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(REFRESH_TOKEN_DURATION).Unix(),
		},
	}
}

func (c RefreshTokenClaims) AccessTokenClaims() AccessTokenClaims {
	return AccessTokenClaims{
		PlayerId: c.PlayerId,
		TeamId:   c.TeamId,
		Username: c.Username,
		Role:     c.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ACCESS_TOKEN_DURATION).Unix(),
		},
	}
}
