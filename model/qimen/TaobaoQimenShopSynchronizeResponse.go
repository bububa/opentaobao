package qimen

import (
	"sync"
)

// TaobaoQimenShopSynchronizeResponse 结构体
type TaobaoQimenShopSynchronizeResponse struct {
	// success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTaobaoQimenShopSynchronizeResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenShopSynchronizeResponse)
	},
}

// GetTaobaoQimenShopSynchronizeResponse() 从对象池中获取TaobaoQimenShopSynchronizeResponse
func GetTaobaoQimenShopSynchronizeResponse() *TaobaoQimenShopSynchronizeResponse {
	return poolTaobaoQimenShopSynchronizeResponse.Get().(*TaobaoQimenShopSynchronizeResponse)
}

// ReleaseTaobaoQimenShopSynchronizeResponse 释放TaobaoQimenShopSynchronizeResponse
func ReleaseTaobaoQimenShopSynchronizeResponse(v *TaobaoQimenShopSynchronizeResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolTaobaoQimenShopSynchronizeResponse.Put(v)
}
