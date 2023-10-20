package taotv

import (
	"sync"
)

// TaobaoTaotvVideoPlaylistOttnavGetResult 结构体
type TaobaoTaotvVideoPlaylistOttnavGetResult struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model *PlayListNavRbo `json:"model,omitempty" xml:"model,omitempty"`
	// httpStatusCode
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoTaotvVideoPlaylistOttnavGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoTaotvVideoPlaylistOttnavGetResult)
	},
}

// GetTaobaoTaotvVideoPlaylistOttnavGetResult() 从对象池中获取TaobaoTaotvVideoPlaylistOttnavGetResult
func GetTaobaoTaotvVideoPlaylistOttnavGetResult() *TaobaoTaotvVideoPlaylistOttnavGetResult {
	return poolTaobaoTaotvVideoPlaylistOttnavGetResult.Get().(*TaobaoTaotvVideoPlaylistOttnavGetResult)
}

// ReleaseTaobaoTaotvVideoPlaylistOttnavGetResult 释放TaobaoTaotvVideoPlaylistOttnavGetResult
func ReleaseTaobaoTaotvVideoPlaylistOttnavGetResult(v *TaobaoTaotvVideoPlaylistOttnavGetResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.HttpStatusCode = 0
	v.Success = false
	poolTaobaoTaotvVideoPlaylistOttnavGetResult.Put(v)
}
