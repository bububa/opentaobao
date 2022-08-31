package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallAlihouseTradeCouponRefundOrderQueryAPIRequest 查询电商券履约退款单 API请求
// tmall.alihouse.trade.coupon.refund.order.query
//
// 查询电商券履约退款单
type TmallAlihouseTradeCouponRefundOrderQueryAPIRequest struct {
	model.Params
	// 查询参数
	_query *CouponOrderQuery
}

// NewTmallAlihouseTradeCouponRefundOrderQueryRequest 初始化TmallAlihouseTradeCouponRefundOrderQueryAPIRequest对象
func NewTmallAlihouseTradeCouponRefundOrderQueryRequest() *TmallAlihouseTradeCouponRefundOrderQueryAPIRequest {
	return &TmallAlihouseTradeCouponRefundOrderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAlihouseTradeCouponRefundOrderQueryAPIRequest) GetApiMethodName() string {
	return "tmall.alihouse.trade.coupon.refund.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAlihouseTradeCouponRefundOrderQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetQuery is Query Setter
// 查询参数
func (r *TmallAlihouseTradeCouponRefundOrderQueryAPIRequest) SetQuery(_query *CouponOrderQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TmallAlihouseTradeCouponRefundOrderQueryAPIRequest) GetQuery() *CouponOrderQuery {
	return r._query
}
