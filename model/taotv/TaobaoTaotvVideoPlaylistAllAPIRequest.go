package taotv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaotvVideoPlaylistAllAPIRequest 获取播单列表 API请求
// taobao.taotv.video.playlist.all
//
// 根据牌照和视频源等获取播单列表
type TaobaoTaotvVideoPlaylistAllAPIRequest struct {
	model.Params
	// 系统信息
	_systemInfo string
}

// NewTaobaoTaotvVideoPlaylistAllRequest 初始化TaobaoTaotvVideoPlaylistAllAPIRequest对象
func NewTaobaoTaotvVideoPlaylistAllRequest() *TaobaoTaotvVideoPlaylistAllAPIRequest {
	return &TaobaoTaotvVideoPlaylistAllAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTaotvVideoPlaylistAllAPIRequest) GetApiMethodName() string {
	return "taobao.taotv.video.playlist.all"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTaotvVideoPlaylistAllAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTaotvVideoPlaylistAllAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemInfo is SystemInfo Setter
// 系统信息
func (r *TaobaoTaotvVideoPlaylistAllAPIRequest) SetSystemInfo(_systemInfo string) error {
	r._systemInfo = _systemInfo
	r.Set("system_info", _systemInfo)
	return nil
}

// GetSystemInfo SystemInfo Getter
func (r TaobaoTaotvVideoPlaylistAllAPIRequest) GetSystemInfo() string {
	return r._systemInfo
}
