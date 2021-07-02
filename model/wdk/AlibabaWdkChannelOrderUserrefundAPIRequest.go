package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkChannelOrderUserrefundAPIRequest 用户发起售后退款(整单/部分) API请求
// alibaba.wdk.channel.order.userrefund
//
// 用户发起售后退款(整单/部分)
type AlibabaWdkChannelOrderUserrefundAPIRequest struct {
	model.Params
	// 退款信息
	_orderUserRefundInfo *OrderUserRefundInfo
}

// NewAlibabaWdkChannelOrderUserrefundRequest 初始化AlibabaWdkChannelOrderUserrefundAPIRequest对象
func NewAlibabaWdkChannelOrderUserrefundRequest() *AlibabaWdkChannelOrderUserrefundAPIRequest {
	return &AlibabaWdkChannelOrderUserrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkChannelOrderUserrefundAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.channel.order.userrefund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkChannelOrderUserrefundAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderUserRefundInfo Setter
// 退款信息
func (r *AlibabaWdkChannelOrderUserrefundAPIRequest) SetOrderUserRefundInfo(_orderUserRefundInfo *OrderUserRefundInfo) error {
	r._orderUserRefundInfo = _orderUserRefundInfo
	r.Set("order_user_refund_info", _orderUserRefundInfo)
	return nil
}

// Get OrderUserRefundInfo Getter
func (r AlibabaWdkChannelOrderUserrefundAPIRequest) GetOrderUserRefundInfo() *OrderUserRefundInfo {
	return r._orderUserRefundInfo
}
