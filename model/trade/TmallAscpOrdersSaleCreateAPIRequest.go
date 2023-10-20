package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallascporderssalecreateAPIRequest ASCP渠道中心销售单创建接口 API请求
// tmall.ascp.orders.sale.create
//
// ASCP渠道中心销售单创建接口
type TmallascporderssalecreateAPIRequest struct {
	model.Params
	// 请求对象
	_channelOrderRequest *CreateChannelOrderRequest
}

// NewTmallascporderssalecreateRequest 初始化TmallascporderssalecreateAPIRequest对象
func NewTmallascporderssalecreateRequest() *TmallascporderssalecreateAPIRequest {
	return &TmallascporderssalecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallascporderssalecreateAPIRequest) GetApiMethodName() string {
	return "tmall.ascp.orders.sale.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallascporderssalecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallascporderssalecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannelOrderRequest is ChannelOrderRequest Setter
// 请求对象
func (r *TmallascporderssalecreateAPIRequest) SetChannelOrderRequest(_channelOrderRequest *CreateChannelOrderRequest) error {
	r._channelOrderRequest = _channelOrderRequest
	r.Set("channel_order_request", _channelOrderRequest)
	return nil
}

// GetChannelOrderRequest ChannelOrderRequest Getter
func (r TmallascporderssalecreateAPIRequest) GetChannelOrderRequest() *CreateChannelOrderRequest {
	return r._channelOrderRequest
}
