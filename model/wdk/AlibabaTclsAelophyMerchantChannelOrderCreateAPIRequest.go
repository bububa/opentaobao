package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophymerchantchannelordercreateAPIRequest 翱象商家自有渠道 订单创建 API请求
// alibaba.tcls.aelophy.merchant.channel.order.create
//
// 翱象小程序渠道订单创建
type AlibabatclsaelophymerchantchannelordercreateAPIRequest struct {
	model.Params
	// 订单信息
	_orderInfo *OrderInfo
}

// NewAlibabatclsaelophymerchantchannelordercreateRequest 初始化AlibabatclsaelophymerchantchannelordercreateAPIRequest对象
func NewAlibabatclsaelophymerchantchannelordercreateRequest() *AlibabatclsaelophymerchantchannelordercreateAPIRequest {
	return &AlibabatclsaelophymerchantchannelordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaelophymerchantchannelordercreateAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.channel.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaelophymerchantchannelordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaelophymerchantchannelordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderInfo is OrderInfo Setter
// 订单信息
func (r *AlibabatclsaelophymerchantchannelordercreateAPIRequest) SetOrderInfo(_orderInfo *OrderInfo) error {
	r._orderInfo = _orderInfo
	r.Set("order_info", _orderInfo)
	return nil
}

// GetOrderInfo OrderInfo Getter
func (r AlibabatclsaelophymerchantchannelordercreateAPIRequest) GetOrderInfo() *OrderInfo {
	return r._orderInfo
}
