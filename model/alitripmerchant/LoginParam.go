package alitripmerchant

// LoginParam 结构体
type LoginParam struct {
	// 用户信息初始化向量
	InfoIv string `json:"info_iv,omitempty" xml:"info_iv,omitempty"`
	// 微信认证code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 加密的手机信息
	EncryptedPhone string `json:"encrypted_phone,omitempty" xml:"encrypted_phone,omitempty"`
	// 验签数据完整签名
	Signature string `json:"signature,omitempty" xml:"signature,omitempty"`
	// 用户ip
	Ip string `json:"ip,omitempty" xml:"ip,omitempty"`
	// 用户信息的加密数据
	EncryptedInfo string `json:"encrypted_info,omitempty" xml:"encrypted_info,omitempty"`
	// 解密用户手机号的初始化向量
	PhoneIv string `json:"phone_iv,omitempty" xml:"phone_iv,omitempty"`
	// 用户的基本信息
	RawData string `json:"raw_data,omitempty" xml:"raw_data,omitempty"`
	// 用户注册来源
	Channel int64 `json:"channel,omitempty" xml:"channel,omitempty"`
	// 是否是老版本
	OldVersion bool `json:"old_version,omitempty" xml:"old_version,omitempty"`
	// 新版本用户信息
	NewUserinfo *NewUserInfo `json:"new_userinfo,omitempty" xml:"new_userinfo,omitempty"`
}
