package larkiot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaolarkiotordergetcinemasAPIRequest 获取iot渠道开放的影院 API请求
// taobao.lark.iot.order.getcinemas
//
// iot渠道拉取有权限访问的影院
type TaobaolarkiotordergetcinemasAPIRequest struct {
	model.Params
	// 渠道编码
	_channelCode string
}

// NewTaobaolarkiotordergetcinemasRequest 初始化TaobaolarkiotordergetcinemasAPIRequest对象
func NewTaobaolarkiotordergetcinemasRequest() *TaobaolarkiotordergetcinemasAPIRequest {
	return &TaobaolarkiotordergetcinemasAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaolarkiotordergetcinemasAPIRequest) GetApiMethodName() string {
	return "taobao.lark.iot.order.getcinemas"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaolarkiotordergetcinemasAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaolarkiotordergetcinemasAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannelCode is ChannelCode Setter
// 渠道编码
func (r *TaobaolarkiotordergetcinemasAPIRequest) SetChannelCode(_channelCode string) error {
	r._channelCode = _channelCode
	r.Set("channel_code", _channelCode)
	return nil
}

// GetChannelCode ChannelCode Getter
func (r TaobaolarkiotordergetcinemasAPIRequest) GetChannelCode() string {
	return r._channelCode
}
