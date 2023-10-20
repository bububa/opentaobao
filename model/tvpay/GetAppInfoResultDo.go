package tvpay

import (
	"sync"
)

// GetAppInfoResultDo 结构体
type GetAppInfoResultDo struct {
	// 包名
	PackageName string `json:"package_name,omitempty" xml:"package_name,omitempty"`
	// 应用配置
	AppConfig *SdkAppConfigDo `json:"app_config,omitempty" xml:"app_config,omitempty"`
	// 全局配置
	GlobalConfig *SdkGlobalConfigDo `json:"global_config,omitempty" xml:"global_config,omitempty"`
	// 商户id
	PartnerId int64 `json:"partner_id,omitempty" xml:"partner_id,omitempty"`
}

var poolGetAppInfoResultDo = sync.Pool{
	New: func() any {
		return new(GetAppInfoResultDo)
	},
}

// GetGetAppInfoResultDo() 从对象池中获取GetAppInfoResultDo
func GetGetAppInfoResultDo() *GetAppInfoResultDo {
	return poolGetAppInfoResultDo.Get().(*GetAppInfoResultDo)
}

// ReleaseGetAppInfoResultDo 释放GetAppInfoResultDo
func ReleaseGetAppInfoResultDo(v *GetAppInfoResultDo) {
	v.PackageName = ""
	v.AppConfig = nil
	v.GlobalConfig = nil
	v.PartnerId = 0
	poolGetAppInfoResultDo.Put(v)
}
