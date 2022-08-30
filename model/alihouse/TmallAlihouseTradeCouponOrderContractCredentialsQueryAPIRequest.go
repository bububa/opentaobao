package alihouse

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIRequest) GetApiMethodName() string {
	return "tmall.alihouse.trade.coupon.order.contract.credentials.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAlihouseTradeCouponOrderContractCredentialsQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
