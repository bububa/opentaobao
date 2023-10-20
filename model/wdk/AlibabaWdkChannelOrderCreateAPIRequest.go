package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkchannelordercreateAPIRequest 创建订单 API请求
// alibaba.wdk.channel.order.create
//
// 外部商家创建订单
type AlibabawdkchannelordercreateAPIRequest struct {
	model.Params
	// 订单信息
	_orderInfo *OrderInfo
}

// NewAlibabawdkchannelordercreateRequest 初始化AlibabawdkchannelordercreateAPIRequest对象
func NewAlibabawdkchannelordercreateRequest() *AlibabawdkchannelordercreateAPIRequest {
	return &AlibabawdkchannelordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkchannelordercreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.channel.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkchannelordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkchannelordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderInfo is OrderInfo Setter
// 订单信息
func (r *AlibabawdkchannelordercreateAPIRequest) SetOrderInfo(_orderInfo *OrderInfo) error {
	r._orderInfo = _orderInfo
	r.Set("order_info", _orderInfo)
	return nil
}

// GetOrderInfo OrderInfo Getter
func (r AlibabawdkchannelordercreateAPIRequest) GetOrderInfo() *OrderInfo {
	return r._orderInfo
}
