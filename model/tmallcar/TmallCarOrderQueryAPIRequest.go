package tmallcar

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallCarOrderQueryAPIRequest) Reset() {
	r._topOrderQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarOrderQueryAPIRequest) GetApiMethodName() string {
	return "tmall.car.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarOrderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCarOrderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTmallCarOrderQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallCarOrderQueryRequest()
	},
}

// GetTmallCarOrderQueryRequest 从 sync.Pool 获取 TmallCarOrderQueryAPIRequest
func GetTmallCarOrderQueryAPIRequest() *TmallCarOrderQueryAPIRequest {
	return poolTmallCarOrderQueryAPIRequest.Get().(*TmallCarOrderQueryAPIRequest)
}

// ReleaseTmallCarOrderQueryAPIRequest 将 TmallCarOrderQueryAPIRequest 放入 sync.Pool
func ReleaseTmallCarOrderQueryAPIRequest(v *TmallCarOrderQueryAPIRequest) {
	v.Reset()
	poolTmallCarOrderQueryAPIRequest.Put(v)
}
