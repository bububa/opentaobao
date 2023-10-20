package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest) Reset() {
	r._orderInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.channel.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTclsAelophyMerchantChannelOrderCreateRequest()
	},
}

// GetAlibabaTclsAelophyMerchantChannelOrderCreateRequest 从 sync.Pool 获取 AlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest
func GetAlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest() *AlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest {
	return poolAlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest.Get().(*AlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest)
}

// ReleaseAlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest 将 AlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest(v *AlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest) {
	v.Reset()
	poolAlibabaTclsAelophyMerchantChannelOrderCreateAPIRequest.Put(v)
}
