package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallAscpOrdersSaleCreateAPIRequest ASCP渠道中心销售单创建接口 API请求
// tmall.ascp.orders.sale.create
//
// ASCP渠道中心销售单创建接口
type TmallAscpOrdersSaleCreateAPIRequest struct {
	model.Params
	// 请求对象
	_channelOrderRequest *CreateChannelOrderRequest
}

// NewTmallAscpOrdersSaleCreateRequest 初始化TmallAscpOrdersSaleCreateAPIRequest对象
func NewTmallAscpOrdersSaleCreateRequest() *TmallAscpOrdersSaleCreateAPIRequest {
	return &TmallAscpOrdersSaleCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAscpOrdersSaleCreateAPIRequest) GetApiMethodName() string {
	return "tmall.ascp.orders.sale.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAscpOrdersSaleCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ChannelOrderRequest Setter
// 请求对象
func (r *TmallAscpOrdersSaleCreateAPIRequest) SetChannelOrderRequest(_channelOrderRequest *CreateChannelOrderRequest) error {
	r._channelOrderRequest = _channelOrderRequest
	r.Set("channel_order_request", _channelOrderRequest)
	return nil
}

// Get ChannelOrderRequest Getter
func (r TmallAscpOrdersSaleCreateAPIRequest) GetChannelOrderRequest() *CreateChannelOrderRequest {
	return r._channelOrderRequest
}
