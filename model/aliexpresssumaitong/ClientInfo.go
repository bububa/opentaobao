package aliexpresssumaitong

import (
	"sync"
)

// ClientInfo 结构体
type ClientInfo struct {
	// 小程序appId
	OpenAppId string `json:"open_app_id,omitempty" xml:"open_app_id,omitempty"`
	// 客户端业务码
	OpenBizCode string `json:"open_biz_code,omitempty" xml:"open_biz_code,omitempty"`
}

var poolClientInfo = sync.Pool{
	New: func() any {
		return new(ClientInfo)
	},
}

// GetClientInfo() 从对象池中获取ClientInfo
func GetClientInfo() *ClientInfo {
	return poolClientInfo.Get().(*ClientInfo)
}

// ReleaseClientInfo 释放ClientInfo
func ReleaseClientInfo(v *ClientInfo) {
	v.OpenAppId = ""
	v.OpenBizCode = ""
	poolClientInfo.Put(v)
}
