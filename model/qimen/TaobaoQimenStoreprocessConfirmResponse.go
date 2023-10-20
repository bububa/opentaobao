package qimen

import (
	"sync"
)

// TaobaoQimenStoreprocessConfirmResponse 结构体
type TaobaoQimenStoreprocessConfirmResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTaobaoQimenStoreprocessConfirmResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenStoreprocessConfirmResponse)
	},
}

// GetTaobaoQimenStoreprocessConfirmResponse() 从对象池中获取TaobaoQimenStoreprocessConfirmResponse
func GetTaobaoQimenStoreprocessConfirmResponse() *TaobaoQimenStoreprocessConfirmResponse {
	return poolTaobaoQimenStoreprocessConfirmResponse.Get().(*TaobaoQimenStoreprocessConfirmResponse)
}

// ReleaseTaobaoQimenStoreprocessConfirmResponse 释放TaobaoQimenStoreprocessConfirmResponse
func ReleaseTaobaoQimenStoreprocessConfirmResponse(v *TaobaoQimenStoreprocessConfirmResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolTaobaoQimenStoreprocessConfirmResponse.Put(v)
}
