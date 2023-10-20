package lstvending

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstvendingtradeflowqueryAPIRequest 自动售卖机交易流水查询 API请求
// alibaba.lst.vending.tradeflow.query
//
// 零售通自动售卖机交易流水查询接口，品牌商通过此接口同步商品交易数据。
type AlibabalstvendingtradeflowqueryAPIRequest struct {
	model.Params
	// 交易流水查询条件
	_openTradeFlowQuery *OpenTradeFlowQuery
}

// NewAlibabalstvendingtradeflowqueryRequest 初始化AlibabalstvendingtradeflowqueryAPIRequest对象
func NewAlibabalstvendingtradeflowqueryRequest() *AlibabalstvendingtradeflowqueryAPIRequest {
	return &AlibabalstvendingtradeflowqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstvendingtradeflowqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.vending.tradeflow.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstvendingtradeflowqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstvendingtradeflowqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenTradeFlowQuery is OpenTradeFlowQuery Setter
// 交易流水查询条件
func (r *AlibabalstvendingtradeflowqueryAPIRequest) SetOpenTradeFlowQuery(_openTradeFlowQuery *OpenTradeFlowQuery) error {
	r._openTradeFlowQuery = _openTradeFlowQuery
	r.Set("open_trade_flow_query", _openTradeFlowQuery)
	return nil
}

// GetOpenTradeFlowQuery OpenTradeFlowQuery Getter
func (r AlibabalstvendingtradeflowqueryAPIRequest) GetOpenTradeFlowQuery() *OpenTradeFlowQuery {
	return r._openTradeFlowQuery
}
