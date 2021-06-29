package campus

// CheckUserMenuParam 
type CheckUserMenuParam struct {
    // 菜单url
    MenuUrl   string `json:"menu_url,omitempty" xml:"menu_url,omitempty"`
    // 用户账号
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
