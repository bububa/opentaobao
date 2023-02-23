package topoaid

// IsPrivacyPackageResponse 结构体
type IsPrivacyPackageResponse struct {
	// 用户授权信息
	UserAuthInfos []UserAuthInfo `json:"user_auth_infos,omitempty" xml:"user_auth_infos>user_auth_info,omitempty"`
	// 是否为隐私包裹
	PrivacyPackage bool `json:"privacy_package,omitempty" xml:"privacy_package,omitempty"`
}
