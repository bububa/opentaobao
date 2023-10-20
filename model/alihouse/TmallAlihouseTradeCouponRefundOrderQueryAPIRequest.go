package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallalihousetradecouponrefundorderqueryAPIRequest 查询电商券履约退款单 API请求
// tmall.alihouse.trade.coupon.refund.order.query
//
// 查询电商券履约退款单
type TmallalihousetradecouponrefundorderqueryAPIRequest struct {
	model.Params
	// 查询参数
	_query *CouponOrderQuery
}

// NewTmallalihousetradecouponrefundorderqueryRequest 初始化TmallalihousetradecouponrefundorderqueryAPIRequest对象
func NewTmallalihousetradecouponrefundorderqueryRequest() *TmallalihousetradecouponrefundorderqueryAPIRequest {
	return &TmallalihousetradecouponrefundorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallalihousetradecouponrefundorderqueryAPIRequest) GetApiMethodName() string {
	return "tmall.alihouse.trade.coupon.refund.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallalihousetradecouponrefundorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallalihousetradecouponrefundorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询参数
func (r *TmallalihousetradecouponrefundorderqueryAPIRequest) SetQuery(_query *CouponOrderQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TmallalihousetradecouponrefundorderqueryAPIRequest) GetQuery() *CouponOrderQuery {
	return r._query
}
