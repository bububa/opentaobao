package alitripmerchant

// UserCurrentStatus 结构体
type UserCurrentStatus struct {
	// token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// 用户当前状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 二次验证场景
	SceneCode string `json:"scene_code,omitempty" xml:"scene_code,omitempty"`
	// 用户姓
	ForNameVerification string `json:"for_name_verification,omitempty" xml:"for_name_verification,omitempty"`
	// 用户信息加密版
	UserEncryptInfo string `json:"user_encrypt_info,omitempty" xml:"user_encrypt_info,omitempty"`
	// 手机号更新换绑提示
	UpdateToastSign bool `json:"update_toast_sign,omitempty" xml:"update_toast_sign,omitempty"`
}
