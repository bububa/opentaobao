package lsttrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsttradeorderquerychangeAPIRequest 订单id批量查询（品牌商视角） API请求
// alibaba.lst.trade.order.querychange
//
// 根据品牌和时间段查询有变更记录的订单id
type AlibabalsttradeorderquerychangeAPIRequest struct {
	model.Params
	// 入参包装类
	_query *LstOrderQuery
}

// NewAlibabalsttradeorderquerychangeRequest 初始化AlibabalsttradeorderquerychangeAPIRequest对象
func NewAlibabalsttradeorderquerychangeRequest() *AlibabalsttradeorderquerychangeAPIRequest {
	return &AlibabalsttradeorderquerychangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalsttradeorderquerychangeAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.order.querychange"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalsttradeorderquerychangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalsttradeorderquerychangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 入参包装类
func (r *AlibabalsttradeorderquerychangeAPIRequest) SetQuery(_query *LstOrderQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabalsttradeorderquerychangeAPIRequest) GetQuery() *LstOrderQuery {
	return r._query
}
