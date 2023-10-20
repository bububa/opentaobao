package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest 查询电商券履约单合同key API请求
// tmall.alihouse.trade.coupon.order.contract.key.query
//
// 查询电商券履约单合同地址
type TmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest struct {
	model.Params
	// 查询参数
	_query *CouponOrderQuery
}

// NewTmallAlihouseTradeCouponOrderContractKeyQueryRequest 初始化TmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest对象
func NewTmallAlihouseTradeCouponOrderContractKeyQueryRequest() *TmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest {
	return &TmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest) Reset() {
	r._query = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest) GetApiMethodName() string {
	return "tmall.alihouse.trade.coupon.order.contract.key.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询参数
func (r *TmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest) SetQuery(_query *CouponOrderQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest) GetQuery() *CouponOrderQuery {
	return r._query
}

var poolTmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallAlihouseTradeCouponOrderContractKeyQueryRequest()
	},
}

// GetTmallAlihouseTradeCouponOrderContractKeyQueryRequest 从 sync.Pool 获取 TmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest
func GetTmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest() *TmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest {
	return poolTmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest.Get().(*TmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest)
}

// ReleaseTmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest 将 TmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest 放入 sync.Pool
func ReleaseTmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest(v *TmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest) {
	v.Reset()
	poolTmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest.Put(v)
}
