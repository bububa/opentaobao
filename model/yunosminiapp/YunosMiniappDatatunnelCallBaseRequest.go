package yunosminiapp

import (
	"sync"
)

// YunosMiniappDatatunnelCallBaseRequest 结构体
type YunosMiniappDatatunnelCallBaseRequest struct {
	// 请求基础参数
	SystemParam *SystemParam `json:"system_param,omitempty" xml:"system_param,omitempty"`
	// 请求参数
	BizParam *YunosMiniappDatatunnelCallBizParam `json:"biz_param,omitempty" xml:"biz_param,omitempty"`
}

var poolYunosMiniappDatatunnelCallBaseRequest = sync.Pool{
	New: func() any {
		return new(YunosMiniappDatatunnelCallBaseRequest)
	},
}

// GetYunosMiniappDatatunnelCallBaseRequest() 从对象池中获取YunosMiniappDatatunnelCallBaseRequest
func GetYunosMiniappDatatunnelCallBaseRequest() *YunosMiniappDatatunnelCallBaseRequest {
	return poolYunosMiniappDatatunnelCallBaseRequest.Get().(*YunosMiniappDatatunnelCallBaseRequest)
}

// ReleaseYunosMiniappDatatunnelCallBaseRequest 释放YunosMiniappDatatunnelCallBaseRequest
func ReleaseYunosMiniappDatatunnelCallBaseRequest(v *YunosMiniappDatatunnelCallBaseRequest) {
	v.SystemParam = nil
	v.BizParam = nil
	poolYunosMiniappDatatunnelCallBaseRequest.Put(v)
}
