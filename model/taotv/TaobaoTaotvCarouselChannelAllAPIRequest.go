package taotv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaotvCarouselChannelAllAPIRequest 获取所有频道列表 API请求
// taobao.taotv.carousel.channel.all
//
// 获取所有频道列表，按照序号升序
type TaobaoTaotvCarouselChannelAllAPIRequest struct {
	model.Params
	// 系统信息
	_systemInfo string
}

// NewTaobaoTaotvCarouselChannelAllRequest 初始化TaobaoTaotvCarouselChannelAllAPIRequest对象
func NewTaobaoTaotvCarouselChannelAllRequest() *TaobaoTaotvCarouselChannelAllAPIRequest {
	return &TaobaoTaotvCarouselChannelAllAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTaotvCarouselChannelAllAPIRequest) GetApiMethodName() string {
	return "taobao.taotv.carousel.channel.all"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTaotvCarouselChannelAllAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSystemInfo is SystemInfo Setter
// 系统信息
func (r *TaobaoTaotvCarouselChannelAllAPIRequest) SetSystemInfo(_systemInfo string) error {
	r._systemInfo = _systemInfo
	r.Set("system_info", _systemInfo)
	return nil
}

// GetSystemInfo SystemInfo Getter
func (r TaobaoTaotvCarouselChannelAllAPIRequest) GetSystemInfo() string {
	return r._systemInfo
}
