package taotv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotaotvvideoplaylistallAPIRequest 获取播单列表 API请求
// taobao.taotv.video.playlist.all
//
// 根据牌照和视频源等获取播单列表
type TaobaotaotvvideoplaylistallAPIRequest struct {
	model.Params
	// 系统信息
	_systemInfo string
}

// NewTaobaotaotvvideoplaylistallRequest 初始化TaobaotaotvvideoplaylistallAPIRequest对象
func NewTaobaotaotvvideoplaylistallRequest() *TaobaotaotvvideoplaylistallAPIRequest {
	return &TaobaotaotvvideoplaylistallAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotaotvvideoplaylistallAPIRequest) GetApiMethodName() string {
	return "taobao.taotv.video.playlist.all"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotaotvvideoplaylistallAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotaotvvideoplaylistallAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemInfo is SystemInfo Setter
// 系统信息
func (r *TaobaotaotvvideoplaylistallAPIRequest) SetSystemInfo(_systemInfo string) error {
	r._systemInfo = _systemInfo
	r.Set("system_info", _systemInfo)
	return nil
}

// GetSystemInfo SystemInfo Getter
func (r TaobaotaotvvideoplaylistallAPIRequest) GetSystemInfo() string {
	return r._systemInfo
}
