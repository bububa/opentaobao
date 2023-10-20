package taotv

import (
	"sync"
)

// TaobaoTaotvVideoPlaylistAllResult 结构体
type TaobaoTaotvVideoPlaylistAllResult struct {
	// model
	ModelList []TaobaoTaotvVideoPlaylistAllModel `json:"model_list,omitempty" xml:"model_list>taobao_taotv_video_playlist_all_model,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// httpStatusCode
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoTaotvVideoPlaylistAllResult = sync.Pool{
	New: func() any {
		return new(TaobaoTaotvVideoPlaylistAllResult)
	},
}

// GetTaobaoTaotvVideoPlaylistAllResult() 从对象池中获取TaobaoTaotvVideoPlaylistAllResult
func GetTaobaoTaotvVideoPlaylistAllResult() *TaobaoTaotvVideoPlaylistAllResult {
	return poolTaobaoTaotvVideoPlaylistAllResult.Get().(*TaobaoTaotvVideoPlaylistAllResult)
}

// ReleaseTaobaoTaotvVideoPlaylistAllResult 释放TaobaoTaotvVideoPlaylistAllResult
func ReleaseTaobaoTaotvVideoPlaylistAllResult(v *TaobaoTaotvVideoPlaylistAllResult) {
	v.ModelList = v.ModelList[:0]
	v.MsgCode = ""
	v.MsgInfo = ""
	v.HttpStatusCode = 0
	v.Success = false
	poolTaobaoTaotvVideoPlaylistAllResult.Put(v)
}
