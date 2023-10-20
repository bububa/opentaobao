package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarorderqueryAPIRequest 天猫汽车整车订单查询 API请求
// tmall.car.order.query
//
// 天猫汽车商家通过该接口查看整车订单信息
type TmallcarorderqueryAPIRequest struct {
	model.Params
	// 入参
	_topOrderQuery *TopOrderQuery
}

// NewTmallcarorderqueryRequest 初始化TmallcarorderqueryAPIRequest对象
func NewTmallcarorderqueryRequest() *TmallcarorderqueryAPIRequest {
	return &TmallcarorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarorderqueryAPIRequest) GetApiMethodName() string {
	return "tmall.car.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopOrderQuery is TopOrderQuery Setter
// 入参
func (r *TmallcarorderqueryAPIRequest) SetTopOrderQuery(_topOrderQuery *TopOrderQuery) error {
	r._topOrderQuery = _topOrderQuery
	r.Set("top_order_query", _topOrderQuery)
	return nil
}

// GetTopOrderQuery TopOrderQuery Getter
func (r TmallcarorderqueryAPIRequest) GetTopOrderQuery() *TopOrderQuery {
	return r._topOrderQuery
}
