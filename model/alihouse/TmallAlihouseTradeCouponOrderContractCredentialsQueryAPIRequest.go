package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIRequest 查询用于电商券履约单合同下载的临时访问凭证 API请求
// tmall.alihouse.trade.coupon.order.contract.credentials.query
//
// 获取用于下载合同的临时aksk和安全token
type TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIRequest struct {
	model.Params
}

// NewTmallAlihouseTradeCouponOrderContractCredentialsQueryRequest 初始化TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIRequest对象
func NewTmallAlihouseTradeCouponOrderContractCredentialsQueryRequest() *TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIRequest {
	return &TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIRequest) GetApiMethodName() string {
	return "tmall.alihouse.trade.coupon.order.contract.credentials.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTmallAlihouseTradeCouponOrderContractCredentialsQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallAlihouseTradeCouponOrderContractCredentialsQueryRequest()
	},
}

// GetTmallAlihouseTradeCouponOrderContractCredentialsQueryRequest 从 sync.Pool 获取 TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIRequest
func GetTmallAlihouseTradeCouponOrderContractCredentialsQueryAPIRequest() *TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIRequest {
	return poolTmallAlihouseTradeCouponOrderContractCredentialsQueryAPIRequest.Get().(*TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIRequest)
}

// ReleaseTmallAlihouseTradeCouponOrderContractCredentialsQueryAPIRequest 将 TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIRequest 放入 sync.Pool
func ReleaseTmallAlihouseTradeCouponOrderContractCredentialsQueryAPIRequest(v *TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIRequest) {
	v.Reset()
	poolTmallAlihouseTradeCouponOrderContractCredentialsQueryAPIRequest.Put(v)
}
