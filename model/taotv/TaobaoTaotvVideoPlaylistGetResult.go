package taotv

import (
	"sync"
)

// TaobaoTaotvVideoPlaylistGetResult 结构体
type TaobaoTaotvVideoPlaylistGetResult struct {
	// 数据列表
	ModelList []TaobaoTaotvVideoPlaylistGetModel `json:"model_list,omitempty" xml:"model_list>taobao_taotv_video_playlist_get_model,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// httpStatusCode
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoTaotvVideoPlaylistGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoTaotvVideoPlaylistGetResult)
	},
}

// GetTaobaoTaotvVideoPlaylistGetResult() 从对象池中获取TaobaoTaotvVideoPlaylistGetResult
func GetTaobaoTaotvVideoPlaylistGetResult() *TaobaoTaotvVideoPlaylistGetResult {
	return poolTaobaoTaotvVideoPlaylistGetResult.Get().(*TaobaoTaotvVideoPlaylistGetResult)
}

// ReleaseTaobaoTaotvVideoPlaylistGetResult 释放TaobaoTaotvVideoPlaylistGetResult
func ReleaseTaobaoTaotvVideoPlaylistGetResult(v *TaobaoTaotvVideoPlaylistGetResult) {
	v.ModelList = v.ModelList[:0]
	v.MsgCode = ""
	v.MsgInfo = ""
	v.HttpStatusCode = 0
	v.Success = false
	poolTaobaoTaotvVideoPlaylistGetResult.Put(v)
}
