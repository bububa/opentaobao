package taotv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotaotvvideoplaylistgetAPIRequest 根据频道ID获取频道下节目单以及当前播放 API请求
// taobao.taotv.video.playlist.get
//
// 根据频道ID获取频道下节目单以及当前播放
type TaobaotaotvvideoplaylistgetAPIRequest struct {
	model.Params
	// 系统信息
	_systemInfo string
	// 播单id
	_playListId int64
}

// NewTaobaotaotvvideoplaylistgetRequest 初始化TaobaotaotvvideoplaylistgetAPIRequest对象
func NewTaobaotaotvvideoplaylistgetRequest() *TaobaotaotvvideoplaylistgetAPIRequest {
	return &TaobaotaotvvideoplaylistgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotaotvvideoplaylistgetAPIRequest) GetApiMethodName() string {
	return "taobao.taotv.video.playlist.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotaotvvideoplaylistgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotaotvvideoplaylistgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemInfo is SystemInfo Setter
// 系统信息
func (r *TaobaotaotvvideoplaylistgetAPIRequest) SetSystemInfo(_systemInfo string) error {
	r._systemInfo = _systemInfo
	r.Set("system_info", _systemInfo)
	return nil
}

// GetSystemInfo SystemInfo Getter
func (r TaobaotaotvvideoplaylistgetAPIRequest) GetSystemInfo() string {
	return r._systemInfo
}

// SetPlayListId is PlayListId Setter
// 播单id
func (r *TaobaotaotvvideoplaylistgetAPIRequest) SetPlayListId(_playListId int64) error {
	r._playListId = _playListId
	r.Set("play_list_id", _playListId)
	return nil
}

// GetPlayListId PlayListId Getter
func (r TaobaotaotvvideoplaylistgetAPIRequest) GetPlayListId() int64 {
	return r._playListId
}
