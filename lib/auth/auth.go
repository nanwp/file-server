package auth

import (
	"context"
	"encoding/json"
	"file-server/lib/config"
	"file-server/lib/helper"
	"net/http"
)

type UserContextKey string
type JWTContextKey string 

const (
  UserContext UserContextKey = "USER_CONTEXT_KEY"
  JWTContext JWTContextKey = "JWT_CONTEXT_KEY"
)

func ValidateCurrentUser(cfg config.APIConfig, r *http.Request) (*User, error) {
  token := r.Header.Get("Authorization")
  if token == "" {
    return nil, helper.NewErrNotFound("Token not found")
  }

  req, err := http.NewRequest("GET", cfg.AuthURL, nil)
  if err != nil {
    return nil, err
  }

  req.Header.Set("Authorization", token)
  client := &http.Client{}
  res, err := client.Do(req)
  if err != nil {
    return nil, err
  }

  if res.StatusCode != http.StatusOK {
    return nil, helper.NewErrUnauthorized("Unauthorized")
  }

  var response GetCurrentUserResonse
  if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
    return nil, err
  }

  return &response.Data, nil
}

func SetUserContext(r *http.Request, user *User) *http.Request {
  ctx := r.Context()
  ctx = context.WithValue(ctx, UserContext, user)
  ctx = context.WithValue(ctx, JWTContext, r.Header.Get("Authorization"))
  return r.WithContext(ctx)
}

func GetUserContext(r *http.Request) *User {
  return r.Context().Value(UserContext).(*User)
}

func GetJWTContext(r *http.Request) string {
  return r.Context().Value(JWTContext).(string)
}

