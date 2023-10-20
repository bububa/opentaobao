package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAlihouseTradeCouponOrderStatusQueryAPIRequest 查询电商券履约单状态 API请求
// tmall.alihouse.trade.coupon.order.status.query
//
// 查询电商券履约单状态
type TmallAlihouseTradeCouponOrderStatusQueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_query *CouponOrderQuery
}

// NewTmallAlihouseTradeCouponOrderStatusQueryRequest 初始化TmallAlihouseTradeCouponOrderStatusQueryAPIRequest对象
func NewTmallAlihouseTradeCouponOrderStatusQueryRequest() *TmallAlihouseTradeCouponOrderStatusQueryAPIRequest {
	return &TmallAlihouseTradeCouponOrderStatusQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallAlihouseTradeCouponOrderStatusQueryAPIRequest) Reset() {
	r._query = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAlihouseTradeCouponOrderStatusQueryAPIRequest) GetApiMethodName() string {
	return "tmall.alihouse.trade.coupon.order.status.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAlihouseTradeCouponOrderStatusQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAlihouseTradeCouponOrderStatusQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 系统自动生成
func (r *TmallAlihouseTradeCouponOrderStatusQueryAPIRequest) SetQuery(_query *CouponOrderQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TmallAlihouseTradeCouponOrderStatusQueryAPIRequest) GetQuery() *CouponOrderQuery {
	return r._query
}

var poolTmallAlihouseTradeCouponOrderStatusQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallAlihouseTradeCouponOrderStatusQueryRequest()
	},
}

// GetTmallAlihouseTradeCouponOrderStatusQueryRequest 从 sync.Pool 获取 TmallAlihouseTradeCouponOrderStatusQueryAPIRequest
func GetTmallAlihouseTradeCouponOrderStatusQueryAPIRequest() *TmallAlihouseTradeCouponOrderStatusQueryAPIRequest {
	return poolTmallAlihouseTradeCouponOrderStatusQueryAPIRequest.Get().(*TmallAlihouseTradeCouponOrderStatusQueryAPIRequest)
}

// ReleaseTmallAlihouseTradeCouponOrderStatusQueryAPIRequest 将 TmallAlihouseTradeCouponOrderStatusQueryAPIRequest 放入 sync.Pool
func ReleaseTmallAlihouseTradeCouponOrderStatusQueryAPIRequest(v *TmallAlihouseTradeCouponOrderStatusQueryAPIRequest) {
	v.Reset()
	poolTmallAlihouseTradeCouponOrderStatusQueryAPIRequest.Put(v)
}
