package alitripmerchant

import (
	"sync"
)

// AlitripMerchantGalaxyActivityGoodsQueryResponse 结构体
type AlitripMerchantGalaxyActivityGoodsQueryResponse struct {
	// 奖品信息数据
	Contents []ActivityDrawUserGoodsVo `json:"contents,omitempty" xml:"contents>activity_draw_user_goods_vo,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误代码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripMerchantGalaxyActivityGoodsQueryResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyActivityGoodsQueryResponse)
	},
}

// GetAlitripMerchantGalaxyActivityGoodsQueryResponse() 从对象池中获取AlitripMerchantGalaxyActivityGoodsQueryResponse
func GetAlitripMerchantGalaxyActivityGoodsQueryResponse() *AlitripMerchantGalaxyActivityGoodsQueryResponse {
	return poolAlitripMerchantGalaxyActivityGoodsQueryResponse.Get().(*AlitripMerchantGalaxyActivityGoodsQueryResponse)
}

// ReleaseAlitripMerchantGalaxyActivityGoodsQueryResponse 释放AlitripMerchantGalaxyActivityGoodsQueryResponse
func ReleaseAlitripMerchantGalaxyActivityGoodsQueryResponse(v *AlitripMerchantGalaxyActivityGoodsQueryResponse) {
	v.Contents = v.Contents[:0]
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Success = false
	poolAlitripMerchantGalaxyActivityGoodsQueryResponse.Put(v)
}
