package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkChannelOrderUserrefundAPIRequest) Reset() {
	r._orderUserRefundInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkChannelOrderUserrefundAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.channel.order.userrefund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkChannelOrderUserrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkChannelOrderUserrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderUserRefundInfo is OrderUserRefundInfo Setter
// 退款信息
func (r *AlibabaWdkChannelOrderUserrefundAPIRequest) SetOrderUserRefundInfo(_orderUserRefundInfo *OrderUserRefundInfo) error {
	r._orderUserRefundInfo = _orderUserRefundInfo
	r.Set("order_user_refund_info", _orderUserRefundInfo)
	return nil
}

// GetOrderUserRefundInfo OrderUserRefundInfo Getter
func (r AlibabaWdkChannelOrderUserrefundAPIRequest) GetOrderUserRefundInfo() *OrderUserRefundInfo {
	return r._orderUserRefundInfo
}

var poolAlibabaWdkChannelOrderUserrefundAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkChannelOrderUserrefundRequest()
	},
}

// GetAlibabaWdkChannelOrderUserrefundRequest 从 sync.Pool 获取 AlibabaWdkChannelOrderUserrefundAPIRequest
func GetAlibabaWdkChannelOrderUserrefundAPIRequest() *AlibabaWdkChannelOrderUserrefundAPIRequest {
	return poolAlibabaWdkChannelOrderUserrefundAPIRequest.Get().(*AlibabaWdkChannelOrderUserrefundAPIRequest)
}

// ReleaseAlibabaWdkChannelOrderUserrefundAPIRequest 将 AlibabaWdkChannelOrderUserrefundAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkChannelOrderUserrefundAPIRequest(v *AlibabaWdkChannelOrderUserrefundAPIRequest) {
	v.Reset()
	poolAlibabaWdkChannelOrderUserrefundAPIRequest.Put(v)
}
