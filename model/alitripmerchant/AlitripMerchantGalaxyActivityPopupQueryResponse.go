package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyActivityPopupQueryResponse 结构体
type AlitripMerchantGalaxyActivityPopupQueryResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 首页抽奖弹窗数据
	Content *ActivityDrawPopupVo `json:"content,omitempty" xml:"content,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyActivityPopupQueryResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyActivityPopupQueryResponse)
	},
}

// GetAlitripMerchantGalaxyActivityPopupQueryResponse() 从对象池中获取AlitripMerchantGalaxyActivityPopupQueryResponse
func GetAlitripMerchantGalaxyActivityPopupQueryResponse() *AlitripMerchantGalaxyActivityPopupQueryResponse {
	return poolAlitripMerchantGalaxyActivityPopupQueryResponse.Get().(*AlitripMerchantGalaxyActivityPopupQueryResponse)
}

// ReleaseAlitripMerchantGalaxyActivityPopupQueryResponse 释放AlitripMerchantGalaxyActivityPopupQueryResponse
func ReleaseAlitripMerchantGalaxyActivityPopupQueryResponse(v *AlitripMerchantGalaxyActivityPopupQueryResponse) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Content = nil
	v.Success = false
	poolAlitripMerchantGalaxyActivityPopupQueryResponse.Put(v)
}
