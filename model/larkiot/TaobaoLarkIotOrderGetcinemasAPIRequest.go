package larkiot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLarkIotOrderGetcinemasAPIRequest 获取iot渠道开放的影院 API请求
// taobao.lark.iot.order.getcinemas
//
// iot渠道拉取有权限访问的影院
type TaobaoLarkIotOrderGetcinemasAPIRequest struct {
	model.Params
	// 渠道编码
	_channelCode string
}

// NewTaobaoLarkIotOrderGetcinemasRequest 初始化TaobaoLarkIotOrderGetcinemasAPIRequest对象
func NewTaobaoLarkIotOrderGetcinemasRequest() *TaobaoLarkIotOrderGetcinemasAPIRequest {
	return &TaobaoLarkIotOrderGetcinemasAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLarkIotOrderGetcinemasAPIRequest) GetApiMethodName() string {
	return "taobao.lark.iot.order.getcinemas"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLarkIotOrderGetcinemasAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLarkIotOrderGetcinemasAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannelCode is ChannelCode Setter
// 渠道编码
func (r *TaobaoLarkIotOrderGetcinemasAPIRequest) SetChannelCode(_channelCode string) error {
	r._channelCode = _channelCode
	r.Set("channel_code", _channelCode)
	return nil
}

// GetChannelCode ChannelCode Getter
func (r TaobaoLarkIotOrderGetcinemasAPIRequest) GetChannelCode() string {
	return r._channelCode
}
