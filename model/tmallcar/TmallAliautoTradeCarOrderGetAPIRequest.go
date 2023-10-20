package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallaliautotradecarordergetAPIRequest 整车订单详情查询 API请求
// tmall.aliauto.trade.car.order.get
//
// 整车订单详情查询接口
type TmallaliautotradecarordergetAPIRequest struct {
	model.Params
	// 入参
	_query *SingleOrderDetailQuery
}

// NewTmallaliautotradecarordergetRequest 初始化TmallaliautotradecarordergetAPIRequest对象
func NewTmallaliautotradecarordergetRequest() *TmallaliautotradecarordergetAPIRequest {
	return &TmallaliautotradecarordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallaliautotradecarordergetAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.trade.car.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallaliautotradecarordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallaliautotradecarordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 入参
func (r *TmallaliautotradecarordergetAPIRequest) SetQuery(_query *SingleOrderDetailQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TmallaliautotradecarordergetAPIRequest) GetQuery() *SingleOrderDetailQuery {
	return r._query
}
