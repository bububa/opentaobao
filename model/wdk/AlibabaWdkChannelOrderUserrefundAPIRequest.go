package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkchannelorderuserrefundAPIRequest 用户发起售后退款(整单/部分) API请求
// alibaba.wdk.channel.order.userrefund
//
// 用户发起售后退款(整单/部分)
type AlibabawdkchannelorderuserrefundAPIRequest struct {
	model.Params
	// 退款信息
	_orderUserRefundInfo *OrderUserRefundInfo
}

// NewAlibabawdkchannelorderuserrefundRequest 初始化AlibabawdkchannelorderuserrefundAPIRequest对象
func NewAlibabawdkchannelorderuserrefundRequest() *AlibabawdkchannelorderuserrefundAPIRequest {
	return &AlibabawdkchannelorderuserrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkchannelorderuserrefundAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.channel.order.userrefund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkchannelorderuserrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkchannelorderuserrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderUserRefundInfo is OrderUserRefundInfo Setter
// 退款信息
func (r *AlibabawdkchannelorderuserrefundAPIRequest) SetOrderUserRefundInfo(_orderUserRefundInfo *OrderUserRefundInfo) error {
	r._orderUserRefundInfo = _orderUserRefundInfo
	r.Set("order_user_refund_info", _orderUserRefundInfo)
	return nil
}

// GetOrderUserRefundInfo OrderUserRefundInfo Getter
func (r AlibabawdkchannelorderuserrefundAPIRequest) GetOrderUserRefundInfo() *OrderUserRefundInfo {
	return r._orderUserRefundInfo
}
