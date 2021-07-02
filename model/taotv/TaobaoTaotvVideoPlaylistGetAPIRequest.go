package taotv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaotvVideoPlaylistGetAPIRequest 根据频道ID获取频道下节目单以及当前播放 API请求
// taobao.taotv.video.playlist.get
//
// 根据频道ID获取频道下节目单以及当前播放
type TaobaoTaotvVideoPlaylistGetAPIRequest struct {
	model.Params
	// 播单id
	_playListId int64
	// 系统信息
	_systemInfo string
}

// NewTaobaoTaotvVideoPlaylistGetRequest 初始化TaobaoTaotvVideoPlaylistGetAPIRequest对象
func NewTaobaoTaotvVideoPlaylistGetRequest() *TaobaoTaotvVideoPlaylistGetAPIRequest {
	return &TaobaoTaotvVideoPlaylistGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTaotvVideoPlaylistGetAPIRequest) GetApiMethodName() string {
	return "taobao.taotv.video.playlist.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTaotvVideoPlaylistGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
