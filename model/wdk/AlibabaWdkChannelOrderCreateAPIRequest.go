package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkChannelOrderCreateAPIRequest 创建订单 API请求
// alibaba.wdk.channel.order.create
//
// 外部商家创建订单
type AlibabaWdkChannelOrderCreateAPIRequest struct {
	model.Params
	// 订单信息
	_orderInfo *OrderInfo
}

// NewAlibabaWdkChannelOrderCreateRequest 初始化AlibabaWdkChannelOrderCreateAPIRequest对象
func NewAlibabaWdkChannelOrderCreateRequest() *AlibabaWdkChannelOrderCreateAPIRequest {
	return &AlibabaWdkChannelOrderCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkChannelOrderCreateAPIRequest) Reset() {
	r._orderInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkChannelOrderCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.channel.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkChannelOrderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkChannelOrderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderInfo is OrderInfo Setter
// 订单信息
func (r *AlibabaWdkChannelOrderCreateAPIRequest) SetOrderInfo(_orderInfo *OrderInfo) error {
	r._orderInfo = _orderInfo
	r.Set("order_info", _orderInfo)
	return nil
}

// GetOrderInfo OrderInfo Getter
func (r AlibabaWdkChannelOrderCreateAPIRequest) GetOrderInfo() *OrderInfo {
	return r._orderInfo
}

var poolAlibabaWdkChannelOrderCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkChannelOrderCreateRequest()
	},
}

// GetAlibabaWdkChannelOrderCreateRequest 从 sync.Pool 获取 AlibabaWdkChannelOrderCreateAPIRequest
func GetAlibabaWdkChannelOrderCreateAPIRequest() *AlibabaWdkChannelOrderCreateAPIRequest {
	return poolAlibabaWdkChannelOrderCreateAPIRequest.Get().(*AlibabaWdkChannelOrderCreateAPIRequest)
}

// ReleaseAlibabaWdkChannelOrderCreateAPIRequest 将 AlibabaWdkChannelOrderCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkChannelOrderCreateAPIRequest(v *AlibabaWdkChannelOrderCreateAPIRequest) {
	v.Reset()
	poolAlibabaWdkChannelOrderCreateAPIRequest.Put(v)
}
