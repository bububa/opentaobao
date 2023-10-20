package lstvending

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstVendingTradeflowQueryAPIRequest 自动售卖机交易流水查询 API请求
// alibaba.lst.vending.tradeflow.query
//
// 零售通自动售卖机交易流水查询接口，品牌商通过此接口同步商品交易数据。
type AlibabaLstVendingTradeflowQueryAPIRequest struct {
	model.Params
	// 交易流水查询条件
	_openTradeFlowQuery *OpenTradeFlowQuery
}

// NewAlibabaLstVendingTradeflowQueryRequest 初始化AlibabaLstVendingTradeflowQueryAPIRequest对象
func NewAlibabaLstVendingTradeflowQueryRequest() *AlibabaLstVendingTradeflowQueryAPIRequest {
	return &AlibabaLstVendingTradeflowQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstVendingTradeflowQueryAPIRequest) Reset() {
	r._openTradeFlowQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstVendingTradeflowQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.vending.tradeflow.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstVendingTradeflowQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstVendingTradeflowQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenTradeFlowQuery is OpenTradeFlowQuery Setter
// 交易流水查询条件
func (r *AlibabaLstVendingTradeflowQueryAPIRequest) SetOpenTradeFlowQuery(_openTradeFlowQuery *OpenTradeFlowQuery) error {
	r._openTradeFlowQuery = _openTradeFlowQuery
	r.Set("open_trade_flow_query", _openTradeFlowQuery)
	return nil
}

// GetOpenTradeFlowQuery OpenTradeFlowQuery Getter
func (r AlibabaLstVendingTradeflowQueryAPIRequest) GetOpenTradeFlowQuery() *OpenTradeFlowQuery {
	return r._openTradeFlowQuery
}

var poolAlibabaLstVendingTradeflowQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstVendingTradeflowQueryRequest()
	},
}

// GetAlibabaLstVendingTradeflowQueryRequest 从 sync.Pool 获取 AlibabaLstVendingTradeflowQueryAPIRequest
func GetAlibabaLstVendingTradeflowQueryAPIRequest() *AlibabaLstVendingTradeflowQueryAPIRequest {
	return poolAlibabaLstVendingTradeflowQueryAPIRequest.Get().(*AlibabaLstVendingTradeflowQueryAPIRequest)
}

// ReleaseAlibabaLstVendingTradeflowQueryAPIRequest 将 AlibabaLstVendingTradeflowQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstVendingTradeflowQueryAPIRequest(v *AlibabaLstVendingTradeflowQueryAPIRequest) {
	v.Reset()
	poolAlibabaLstVendingTradeflowQueryAPIRequest.Put(v)
}
