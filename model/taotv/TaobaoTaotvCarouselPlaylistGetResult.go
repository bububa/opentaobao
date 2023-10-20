package taotv

import (
	"sync"
)

// TaobaoTaotvCarouselPlaylistGetResult 结构体
type TaobaoTaotvCarouselPlaylistGetResult struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回数据
	Model *CarouselChannelRbo `json:"model,omitempty" xml:"model,omitempty"`
	// httpStatusCode
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoTaotvCarouselPlaylistGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoTaotvCarouselPlaylistGetResult)
	},
}

// GetTaobaoTaotvCarouselPlaylistGetResult() 从对象池中获取TaobaoTaotvCarouselPlaylistGetResult
func GetTaobaoTaotvCarouselPlaylistGetResult() *TaobaoTaotvCarouselPlaylistGetResult {
	return poolTaobaoTaotvCarouselPlaylistGetResult.Get().(*TaobaoTaotvCarouselPlaylistGetResult)
}

// ReleaseTaobaoTaotvCarouselPlaylistGetResult 释放TaobaoTaotvCarouselPlaylistGetResult
func ReleaseTaobaoTaotvCarouselPlaylistGetResult(v *TaobaoTaotvCarouselPlaylistGetResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.HttpStatusCode = 0
	v.Success = false
	poolTaobaoTaotvCarouselPlaylistGetResult.Put(v)
}
