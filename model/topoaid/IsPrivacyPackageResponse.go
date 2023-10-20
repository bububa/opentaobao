package topoaid

import (
	"sync"
)

// IsPrivacyPackageResponse 结构体
type IsPrivacyPackageResponse struct {
	// 用户授权信息
	UserAuthInfos []UserAuthInfo `json:"user_auth_infos,omitempty" xml:"user_auth_infos>user_auth_info,omitempty"`
	// 是否为隐私包裹
	PrivacyPackage bool `json:"privacy_package,omitempty" xml:"privacy_package,omitempty"`
}

var poolIsPrivacyPackageResponse = sync.Pool{
	New: func() any {
		return new(IsPrivacyPackageResponse)
	},
}

// GetIsPrivacyPackageResponse() 从对象池中获取IsPrivacyPackageResponse
func GetIsPrivacyPackageResponse() *IsPrivacyPackageResponse {
	return poolIsPrivacyPackageResponse.Get().(*IsPrivacyPackageResponse)
}

// ReleaseIsPrivacyPackageResponse 释放IsPrivacyPackageResponse
func ReleaseIsPrivacyPackageResponse(v *IsPrivacyPackageResponse) {
	v.UserAuthInfos = v.UserAuthInfos[:0]
	v.PrivacyPackage = false
	poolIsPrivacyPackageResponse.Put(v)
}
