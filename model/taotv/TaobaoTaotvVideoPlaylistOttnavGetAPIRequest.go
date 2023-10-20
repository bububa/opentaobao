package taotv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaotvVideoPlaylistOttnavGetAPIRequest ott播单 API请求
// taobao.taotv.video.playlist.ottnav.get
//
// 根据聚焦播单ID拿到下面播单视频，根据左侧播单ID列表批量拿到播单信息
type TaobaoTaotvVideoPlaylistOttnavGetAPIRequest struct {
	model.Params
	// 播单列表
	_playListNav []string
	// 系统信息
	_systemInfo string
	// 播单id
	_playListId int64
}

// NewTaobaoTaotvVideoPlaylistOttnavGetRequest 初始化TaobaoTaotvVideoPlaylistOttnavGetAPIRequest对象
func NewTaobaoTaotvVideoPlaylistOttnavGetRequest() *TaobaoTaotvVideoPlaylistOttnavGetAPIRequest {
	return &TaobaoTaotvVideoPlaylistOttnavGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTaotvVideoPlaylistOttnavGetAPIRequest) GetApiMethodName() string {
	return "taobao.taotv.video.playlist.ottnav.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTaotvVideoPlaylistOttnavGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTaotvVideoPlaylistOttnavGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPlayListNav is PlayListNav Setter
// 播单列表
func (r *TaobaoTaotvVideoPlaylistOttnavGetAPIRequest) SetPlayListNav(_playListNav []string) error {
	r._playListNav = _playListNav
	r.Set("play_list_nav", _playListNav)
	return nil
}

// GetPlayListNav PlayListNav Getter
func (r TaobaoTaotvVideoPlaylistOttnavGetAPIRequest) GetPlayListNav() []string {
	return r._playListNav
}

// SetSystemInfo is SystemInfo Setter
// 系统信息
func (r *TaobaoTaotvVideoPlaylistOttnavGetAPIRequest) SetSystemInfo(_systemInfo string) error {
	r._systemInfo = _systemInfo
	r.Set("system_info", _systemInfo)
	return nil
}

// GetSystemInfo SystemInfo Getter
func (r TaobaoTaotvVideoPlaylistOttnavGetAPIRequest) GetSystemInfo() string {
	return r._systemInfo
}

// SetPlayListId is PlayListId Setter
// 播单id
func (r *TaobaoTaotvVideoPlaylistOttnavGetAPIRequest) SetPlayListId(_playListId int64) error {
	r._playListId = _playListId
	r.Set("play_list_id", _playListId)
	return nil
}

// GetPlayListId PlayListId Getter
func (r TaobaoTaotvVideoPlaylistOttnavGetAPIRequest) GetPlayListId() int64 {
	return r._playListId
}
