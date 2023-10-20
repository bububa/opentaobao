package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTradeOrderBalanceBillQueryAPIRequest 分页拉取订单数据 API请求
// alibaba.wdk.trade.order.balance.bill.query
//
// 提供接口供外部调用，分页拉取订单数据
type AlibabaWdkTradeOrderBalanceBillQueryAPIRequest struct {
	model.Params
	// 入参
	_orderBalanceBillRequest *OrderBalanceBillRequest
}

// NewAlibabaWdkTradeOrderBalanceBillQueryRequest 初始化AlibabaWdkTradeOrderBalanceBillQueryAPIRequest对象
func NewAlibabaWdkTradeOrderBalanceBillQueryRequest() *AlibabaWdkTradeOrderBalanceBillQueryAPIRequest {
	return &AlibabaWdkTradeOrderBalanceBillQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkTradeOrderBalanceBillQueryAPIRequest) Reset() {
	r._orderBalanceBillRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkTradeOrderBalanceBillQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trade.order.balance.bill.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkTradeOrderBalanceBillQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkTradeOrderBalanceBillQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderBalanceBillRequest is OrderBalanceBillRequest Setter
// 入参
func (r *AlibabaWdkTradeOrderBalanceBillQueryAPIRequest) SetOrderBalanceBillRequest(_orderBalanceBillRequest *OrderBalanceBillRequest) error {
	r._orderBalanceBillRequest = _orderBalanceBillRequest
	r.Set("order_balance_bill_request", _orderBalanceBillRequest)
	return nil
}

// GetOrderBalanceBillRequest OrderBalanceBillRequest Getter
func (r AlibabaWdkTradeOrderBalanceBillQueryAPIRequest) GetOrderBalanceBillRequest() *OrderBalanceBillRequest {
	return r._orderBalanceBillRequest
}

var poolAlibabaWdkTradeOrderBalanceBillQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkTradeOrderBalanceBillQueryRequest()
	},
}

// GetAlibabaWdkTradeOrderBalanceBillQueryRequest 从 sync.Pool 获取 AlibabaWdkTradeOrderBalanceBillQueryAPIRequest
func GetAlibabaWdkTradeOrderBalanceBillQueryAPIRequest() *AlibabaWdkTradeOrderBalanceBillQueryAPIRequest {
	return poolAlibabaWdkTradeOrderBalanceBillQueryAPIRequest.Get().(*AlibabaWdkTradeOrderBalanceBillQueryAPIRequest)
}

// ReleaseAlibabaWdkTradeOrderBalanceBillQueryAPIRequest 将 AlibabaWdkTradeOrderBalanceBillQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkTradeOrderBalanceBillQueryAPIRequest(v *AlibabaWdkTradeOrderBalanceBillQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkTradeOrderBalanceBillQueryAPIRequest.Put(v)
}
