package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophymerchantchannelorderupdatestatusAPIRequest 翱象商家自有渠道 订单状态更新 API请求
// alibaba.tcls.aelophy.merchant.channel.order.updatestatus
//
// 订单状态变更
type AlibabatclsaelophymerchantchannelorderupdatestatusAPIRequest struct {
	model.Params
	// 修改信息
	_orderStatusInfo *OrderStatusInfo
}

// NewAlibabatclsaelophymerchantchannelorderupdatestatusRequest 初始化AlibabatclsaelophymerchantchannelorderupdatestatusAPIRequest对象
func NewAlibabatclsaelophymerchantchannelorderupdatestatusRequest() *AlibabatclsaelophymerchantchannelorderupdatestatusAPIRequest {
	return &AlibabatclsaelophymerchantchannelorderupdatestatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaelophymerchantchannelorderupdatestatusAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.channel.order.updatestatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaelophymerchantchannelorderupdatestatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaelophymerchantchannelorderupdatestatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderStatusInfo is OrderStatusInfo Setter
// 修改信息
func (r *AlibabatclsaelophymerchantchannelorderupdatestatusAPIRequest) SetOrderStatusInfo(_orderStatusInfo *OrderStatusInfo) error {
	r._orderStatusInfo = _orderStatusInfo
	r.Set("order_status_info", _orderStatusInfo)
	return nil
}

// GetOrderStatusInfo OrderStatusInfo Getter
func (r AlibabatclsaelophymerchantchannelorderupdatestatusAPIRequest) GetOrderStatusInfo() *OrderStatusInfo {
	return r._orderStatusInfo
}
