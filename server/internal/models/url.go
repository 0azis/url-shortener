package models

type Url struct {
	ID     string `json:"UUID"`
	UserID int    `json:"-"`
	Origin string `json:"origin"`
}

type UrlCredentials struct {
	UserID int    `json:"userID"`
	Origin string `json:"origin"`
}

func (u UrlCredentials) Validate() bool {
	if u.Origin == "" {
		return false
	}

	return true
}
