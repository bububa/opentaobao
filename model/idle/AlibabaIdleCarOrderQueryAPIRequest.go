package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlecarorderqueryAPIRequest 二手车寄卖查询订单接口 API请求
// alibaba.idle.car.order.query
//
// 二手车寄卖查询订单接口
type AlibabaidlecarorderqueryAPIRequest struct {
	model.Params
	// 查询订单入参
	_query *ConsignmentOrderQuery
}

// NewAlibabaidlecarorderqueryRequest 初始化AlibabaidlecarorderqueryAPIRequest对象
func NewAlibabaidlecarorderqueryRequest() *AlibabaidlecarorderqueryAPIRequest {
	return &AlibabaidlecarorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidlecarorderqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.car.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidlecarorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidlecarorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询订单入参
func (r *AlibabaidlecarorderqueryAPIRequest) SetQuery(_query *ConsignmentOrderQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaidlecarorderqueryAPIRequest) GetQuery() *ConsignmentOrderQuery {
	return r._query
}
