package taotv

import (
	"sync"
)

// TaobaoTaotvVideoPlaylistPageResult 结构体
type TaobaoTaotvVideoPlaylistPageResult struct {
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 状态码
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// 播单信息分页信息
	Model *TaobaoTaotvVideoPlaylistPageModel `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoTaotvVideoPlaylistPageResult = sync.Pool{
	New: func() any {
		return new(TaobaoTaotvVideoPlaylistPageResult)
	},
}

// GetTaobaoTaotvVideoPlaylistPageResult() 从对象池中获取TaobaoTaotvVideoPlaylistPageResult
func GetTaobaoTaotvVideoPlaylistPageResult() *TaobaoTaotvVideoPlaylistPageResult {
	return poolTaobaoTaotvVideoPlaylistPageResult.Get().(*TaobaoTaotvVideoPlaylistPageResult)
}

// ReleaseTaobaoTaotvVideoPlaylistPageResult 释放TaobaoTaotvVideoPlaylistPageResult
func ReleaseTaobaoTaotvVideoPlaylistPageResult(v *TaobaoTaotvVideoPlaylistPageResult) {
	v.MsgInfo = ""
	v.MsgCode = ""
	v.HttpStatusCode = 0
	v.Model = nil
	v.Success = false
	poolTaobaoTaotvVideoPlaylistPageResult.Put(v)
}
