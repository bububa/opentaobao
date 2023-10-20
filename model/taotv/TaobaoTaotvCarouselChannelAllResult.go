package taotv

import (
	"sync"
)

// TaobaoTaotvCarouselChannelAllResult 结构体
type TaobaoTaotvCarouselChannelAllResult struct {
	// 频道列表
	ModelList []TaobaoTaotvCarouselChannelAllModel `json:"model_list,omitempty" xml:"model_list>taobao_taotv_carousel_channel_all_model,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// httpStatusCode
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoTaotvCarouselChannelAllResult = sync.Pool{
	New: func() any {
		return new(TaobaoTaotvCarouselChannelAllResult)
	},
}

// GetTaobaoTaotvCarouselChannelAllResult() 从对象池中获取TaobaoTaotvCarouselChannelAllResult
func GetTaobaoTaotvCarouselChannelAllResult() *TaobaoTaotvCarouselChannelAllResult {
	return poolTaobaoTaotvCarouselChannelAllResult.Get().(*TaobaoTaotvCarouselChannelAllResult)
}

// ReleaseTaobaoTaotvCarouselChannelAllResult 释放TaobaoTaotvCarouselChannelAllResult
func ReleaseTaobaoTaotvCarouselChannelAllResult(v *TaobaoTaotvCarouselChannelAllResult) {
	v.ModelList = v.ModelList[:0]
	v.MsgCode = ""
	v.MsgInfo = ""
	v.HttpStatusCode = 0
	v.Success = false
	poolTaobaoTaotvCarouselChannelAllResult.Put(v)
}
