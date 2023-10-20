package taotv

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaotvVideoPlaylistGetAPIRequest 根据频道ID获取频道下节目单以及当前播放 API请求
// taobao.taotv.video.playlist.get
//
// 根据频道ID获取频道下节目单以及当前播放
type TaobaoTaotvVideoPlaylistGetAPIRequest struct {
	model.Params
	// 系统信息
	_systemInfo string
	// 播单id
	_playListId int64
}

// NewTaobaoTaotvVideoPlaylistGetRequest 初始化TaobaoTaotvVideoPlaylistGetAPIRequest对象
func NewTaobaoTaotvVideoPlaylistGetRequest() *TaobaoTaotvVideoPlaylistGetAPIRequest {
	return &TaobaoTaotvVideoPlaylistGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTaotvVideoPlaylistGetAPIRequest) Reset() {
	r._systemInfo = ""
	r._playListId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTaotvVideoPlaylistGetAPIRequest) GetApiMethodName() string {
	return "taobao.taotv.video.playlist.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTaotvVideoPlaylistGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTaotvVideoPlaylistGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemInfo is SystemInfo Setter
// 系统信息
func (r *TaobaoTaotvVideoPlaylistGetAPIRequest) SetSystemInfo(_systemInfo string) error {
	r._systemInfo = _systemInfo
	r.Set("system_info", _systemInfo)
	return nil
}

// GetSystemInfo SystemInfo Getter
func (r TaobaoTaotvVideoPlaylistGetAPIRequest) GetSystemInfo() string {
	return r._systemInfo
}

// SetPlayListId is PlayListId Setter
// 播单id
func (r *TaobaoTaotvVideoPlaylistGetAPIRequest) SetPlayListId(_playListId int64) error {
	r._playListId = _playListId
	r.Set("play_list_id", _playListId)
	return nil
}

// GetPlayListId PlayListId Getter
func (r TaobaoTaotvVideoPlaylistGetAPIRequest) GetPlayListId() int64 {
	return r._playListId
}

var poolTaobaoTaotvVideoPlaylistGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTaotvVideoPlaylistGetRequest()
	},
}

// GetTaobaoTaotvVideoPlaylistGetRequest 从 sync.Pool 获取 TaobaoTaotvVideoPlaylistGetAPIRequest
func GetTaobaoTaotvVideoPlaylistGetAPIRequest() *TaobaoTaotvVideoPlaylistGetAPIRequest {
	return poolTaobaoTaotvVideoPlaylistGetAPIRequest.Get().(*TaobaoTaotvVideoPlaylistGetAPIRequest)
}

// ReleaseTaobaoTaotvVideoPlaylistGetAPIRequest 将 TaobaoTaotvVideoPlaylistGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTaotvVideoPlaylistGetAPIRequest(v *TaobaoTaotvVideoPlaylistGetAPIRequest) {
	v.Reset()
	poolTaobaoTaotvVideoPlaylistGetAPIRequest.Put(v)
}
