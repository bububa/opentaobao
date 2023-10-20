package alihouse

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallAlihouseTradeCouponRefundOrderQueryAPIRequest) Reset() {
	r._query = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAlihouseTradeCouponRefundOrderQueryAPIRequest) GetApiMethodName() string {
	return "tmall.alihouse.trade.coupon.refund.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAlihouseTradeCouponRefundOrderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAlihouseTradeCouponRefundOrderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTmallAlihouseTradeCouponRefundOrderQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallAlihouseTradeCouponRefundOrderQueryRequest()
	},
}

// GetTmallAlihouseTradeCouponRefundOrderQueryRequest 从 sync.Pool 获取 TmallAlihouseTradeCouponRefundOrderQueryAPIRequest
func GetTmallAlihouseTradeCouponRefundOrderQueryAPIRequest() *TmallAlihouseTradeCouponRefundOrderQueryAPIRequest {
	return poolTmallAlihouseTradeCouponRefundOrderQueryAPIRequest.Get().(*TmallAlihouseTradeCouponRefundOrderQueryAPIRequest)
}

// ReleaseTmallAlihouseTradeCouponRefundOrderQueryAPIRequest 将 TmallAlihouseTradeCouponRefundOrderQueryAPIRequest 放入 sync.Pool
func ReleaseTmallAlihouseTradeCouponRefundOrderQueryAPIRequest(v *TmallAlihouseTradeCouponRefundOrderQueryAPIRequest) {
	v.Reset()
	poolTmallAlihouseTradeCouponRefundOrderQueryAPIRequest.Put(v)
}
