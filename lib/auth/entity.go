package auth

type GetCurrentUserResonse struct {
  StatusCode int `json:"status_code"`
  Data User `json:"data"`
}

type User struct {
  ID int `json:"id"`
  Name string `json:"name"`
  Role Role `json:"role"`
}

type Role struct {
  ID int `json:"id"`
  Name string `json:"name"`
}
