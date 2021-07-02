package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest 翱象商家自有渠道 订单创建 API请求
// alibaba.tcls.aelophy.merchant.channel.order.create
//
// 翱象小程序渠道订单创建
type AlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest struct {
	model.Params
	// 订单信息
	_orderInfo *OrderInfo
}

// NewAlibabaTclsAelophyMerchantChannelOrderCreateRequest 初始化AlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest对象
func NewAlibabaTclsAelophyMerchantChannelOrderCreateRequest() *AlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest {
	return &AlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.channel.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderInfo is OrderInfo Setter
// 订单信息
func (r *AlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest) SetOrderInfo(_orderInfo *OrderInfo) error {
	r._orderInfo = _orderInfo
	r.Set("order_info", _orderInfo)
	return nil
}

// GetOrderInfo OrderInfo Getter
func (r AlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest) GetOrderInfo() *OrderInfo {
	return r._orderInfo
}
