package taotv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotaotvcarouselchannelallAPIRequest 获取所有频道列表 API请求
// taobao.taotv.carousel.channel.all
//
// 获取所有频道列表，按照序号升序
type TaobaotaotvcarouselchannelallAPIRequest struct {
	model.Params
	// 系统信息
	_systemInfo string
}

// NewTaobaotaotvcarouselchannelallRequest 初始化TaobaotaotvcarouselchannelallAPIRequest对象
func NewTaobaotaotvcarouselchannelallRequest() *TaobaotaotvcarouselchannelallAPIRequest {
	return &TaobaotaotvcarouselchannelallAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotaotvcarouselchannelallAPIRequest) GetApiMethodName() string {
	return "taobao.taotv.carousel.channel.all"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotaotvcarouselchannelallAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotaotvcarouselchannelallAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemInfo is SystemInfo Setter
// 系统信息
func (r *TaobaotaotvcarouselchannelallAPIRequest) SetSystemInfo(_systemInfo string) error {
	r._systemInfo = _systemInfo
	r.Set("system_info", _systemInfo)
	return nil
}

// GetSystemInfo SystemInfo Getter
func (r TaobaotaotvcarouselchannelallAPIRequest) GetSystemInfo() string {
	return r._systemInfo
}
