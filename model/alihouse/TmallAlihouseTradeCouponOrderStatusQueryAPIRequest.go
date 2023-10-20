package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallalihousetradecouponorderstatusqueryAPIRequest 查询电商券履约单状态 API请求
// tmall.alihouse.trade.coupon.order.status.query
//
// 查询电商券履约单状态
type TmallalihousetradecouponorderstatusqueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_query *CouponOrderQuery
}

// NewTmallalihousetradecouponorderstatusqueryRequest 初始化TmallalihousetradecouponorderstatusqueryAPIRequest对象
func NewTmallalihousetradecouponorderstatusqueryRequest() *TmallalihousetradecouponorderstatusqueryAPIRequest {
	return &TmallalihousetradecouponorderstatusqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallalihousetradecouponorderstatusqueryAPIRequest) GetApiMethodName() string {
	return "tmall.alihouse.trade.coupon.order.status.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallalihousetradecouponorderstatusqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallalihousetradecouponorderstatusqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 系统自动生成
func (r *TmallalihousetradecouponorderstatusqueryAPIRequest) SetQuery(_query *CouponOrderQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TmallalihousetradecouponorderstatusqueryAPIRequest) GetQuery() *CouponOrderQuery {
	return r._query
}
