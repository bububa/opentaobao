package security

// RpUserResult 结构体
type RpUserResult struct {
	// users
	Users []string `json:"users,omitempty" xml:"users>string,omitempty"`
}
