package tmallcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoTradeCarOrderGetAPIRequest 整车订单详情查询 API请求
// tmall.aliauto.trade.car.order.get
//
// 整车订单详情查询接口
type TmallAliautoTradeCarOrderGetAPIRequest struct {
	model.Params
	// 入参
	_query *SingleOrderDetailQuery
}

// NewTmallAliautoTradeCarOrderGetRequest 初始化TmallAliautoTradeCarOrderGetAPIRequest对象
func NewTmallAliautoTradeCarOrderGetRequest() *TmallAliautoTradeCarOrderGetAPIRequest {
	return &TmallAliautoTradeCarOrderGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallAliautoTradeCarOrderGetAPIRequest) Reset() {
	r._query = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoTradeCarOrderGetAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.trade.car.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoTradeCarOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAliautoTradeCarOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 入参
func (r *TmallAliautoTradeCarOrderGetAPIRequest) SetQuery(_query *SingleOrderDetailQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TmallAliautoTradeCarOrderGetAPIRequest) GetQuery() *SingleOrderDetailQuery {
	return r._query
}

var poolTmallAliautoTradeCarOrderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallAliautoTradeCarOrderGetRequest()
	},
}

// GetTmallAliautoTradeCarOrderGetRequest 从 sync.Pool 获取 TmallAliautoTradeCarOrderGetAPIRequest
func GetTmallAliautoTradeCarOrderGetAPIRequest() *TmallAliautoTradeCarOrderGetAPIRequest {
	return poolTmallAliautoTradeCarOrderGetAPIRequest.Get().(*TmallAliautoTradeCarOrderGetAPIRequest)
}

// ReleaseTmallAliautoTradeCarOrderGetAPIRequest 将 TmallAliautoTradeCarOrderGetAPIRequest 放入 sync.Pool
func ReleaseTmallAliautoTradeCarOrderGetAPIRequest(v *TmallAliautoTradeCarOrderGetAPIRequest) {
	v.Reset()
	poolTmallAliautoTradeCarOrderGetAPIRequest.Put(v)
}
