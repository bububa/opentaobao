package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyMerchantChannelOrderUpdatestatusAPIRequest 翱象商家自有渠道 订单状态更新 API请求
// alibaba.tcls.aelophy.merchant.channel.order.updatestatus
//
// 订单状态变更
type AlibabaTclsAelophyMerchantChannelOrderUpdatestatusAPIRequest struct {
	model.Params
	// 修改信息
	_orderStatusInfo *OrderStatusInfo
}

// NewAlibabaTclsAelophyMerchantChannelOrderUpdatestatusRequest 初始化AlibabaTclsAelophyMerchantChannelOrderUpdatestatusAPIRequest对象
func NewAlibabaTclsAelophyMerchantChannelOrderUpdatestatusRequest() *AlibabaTclsAelophyMerchantChannelOrderUpdatestatusAPIRequest {
	return &AlibabaTclsAelophyMerchantChannelOrderUpdatestatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantChannelOrderUpdatestatusAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.channel.order.updatestatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantChannelOrderUpdatestatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTclsAelophyMerchantChannelOrderUpdatestatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderStatusInfo is OrderStatusInfo Setter
// 修改信息
func (r *AlibabaTclsAelophyMerchantChannelOrderUpdatestatusAPIRequest) SetOrderStatusInfo(_orderStatusInfo *OrderStatusInfo) error {
	r._orderStatusInfo = _orderStatusInfo
	r.Set("order_status_info", _orderStatusInfo)
	return nil
}

// GetOrderStatusInfo OrderStatusInfo Getter
func (r AlibabaTclsAelophyMerchantChannelOrderUpdatestatusAPIRequest) GetOrderStatusInfo() *OrderStatusInfo {
	return r._orderStatusInfo
}
