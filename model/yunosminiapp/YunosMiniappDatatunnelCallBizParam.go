package yunosminiapp

import (
	"sync"
)

// YunosMiniappDatatunnelCallBizParam 结构体
type YunosMiniappDatatunnelCallBizParam struct {
	// 请求参数详细
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 业务操作
	Scene string `json:"scene,omitempty" xml:"scene,omitempty"`
}

var poolYunosMiniappDatatunnelCallBizParam = sync.Pool{
	New: func() any {
		return new(YunosMiniappDatatunnelCallBizParam)
	},
}

// GetYunosMiniappDatatunnelCallBizParam() 从对象池中获取YunosMiniappDatatunnelCallBizParam
func GetYunosMiniappDatatunnelCallBizParam() *YunosMiniappDatatunnelCallBizParam {
	return poolYunosMiniappDatatunnelCallBizParam.Get().(*YunosMiniappDatatunnelCallBizParam)
}

// ReleaseYunosMiniappDatatunnelCallBizParam 释放YunosMiniappDatatunnelCallBizParam
func ReleaseYunosMiniappDatatunnelCallBizParam(v *YunosMiniappDatatunnelCallBizParam) {
	v.Data = ""
	v.Scene = ""
	poolYunosMiniappDatatunnelCallBizParam.Put(v)
}
