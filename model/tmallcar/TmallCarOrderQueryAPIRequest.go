package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCarOrderQueryAPIRequest 天猫汽车整车订单查询 API请求
// tmall.car.order.query
//
// 天猫汽车商家通过该接口查看整车订单信息
type TmallCarOrderQueryAPIRequest struct {
	model.Params
	// 入参
	_topOrderQuery *TopOrderQuery
}

// NewTmallCarOrderQueryRequest 初始化TmallCarOrderQueryAPIRequest对象
func NewTmallCarOrderQueryRequest() *TmallCarOrderQueryAPIRequest {
	return &TmallCarOrderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarOrderQueryAPIRequest) GetApiMethodName() string {
	return "tmall.car.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarOrderQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTopOrderQuery is TopOrderQuery Setter
// 入参
func (r *TmallCarOrderQueryAPIRequest) SetTopOrderQuery(_topOrderQuery *TopOrderQuery) error {
	r._topOrderQuery = _topOrderQuery
	r.Set("top_order_query", _topOrderQuery)
	return nil
}

// GetTopOrderQuery TopOrderQuery Getter
func (r TmallCarOrderQueryAPIRequest) GetTopOrderQuery() *TopOrderQuery {
	return r._topOrderQuery
}
