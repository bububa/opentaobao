package btrip

// OpenUserInfo 
type OpenUserInfo struct {
    // 出行人名称
    UserName   string `json:"user_name,omitempty" xml:"user_name,omitempty"`
    // 出行人id
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
