package trade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTradeOrderQueryAPIRequest 查询外部交易订单接口 API请求
// alibaba.wdk.trade.order.query
//
// 通过该接口可以在盒马查询交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
type AlibabaWdkTradeOrderQueryAPIRequest struct {
	model.Params
	// 订单查询
	_query *TradeOrderQuery
}

// NewAlibabaWdkTradeOrderQueryRequest 初始化AlibabaWdkTradeOrderQueryAPIRequest对象
func NewAlibabaWdkTradeOrderQueryRequest() *AlibabaWdkTradeOrderQueryAPIRequest {
	return &AlibabaWdkTradeOrderQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkTradeOrderQueryAPIRequest) Reset() {
	r._query = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkTradeOrderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trade.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkTradeOrderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkTradeOrderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 订单查询
func (r *AlibabaWdkTradeOrderQueryAPIRequest) SetQuery(_query *TradeOrderQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaWdkTradeOrderQueryAPIRequest) GetQuery() *TradeOrderQuery {
	return r._query
}

var poolAlibabaWdkTradeOrderQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkTradeOrderQueryRequest()
	},
}

// GetAlibabaWdkTradeOrderQueryRequest 从 sync.Pool 获取 AlibabaWdkTradeOrderQueryAPIRequest
func GetAlibabaWdkTradeOrderQueryAPIRequest() *AlibabaWdkTradeOrderQueryAPIRequest {
	return poolAlibabaWdkTradeOrderQueryAPIRequest.Get().(*AlibabaWdkTradeOrderQueryAPIRequest)
}

// ReleaseAlibabaWdkTradeOrderQueryAPIRequest 将 AlibabaWdkTradeOrderQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkTradeOrderQueryAPIRequest(v *AlibabaWdkTradeOrderQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkTradeOrderQueryAPIRequest.Put(v)
}
