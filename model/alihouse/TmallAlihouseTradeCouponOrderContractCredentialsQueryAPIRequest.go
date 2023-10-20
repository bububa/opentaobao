package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallalihousetradecouponordercontractcredentialsqueryAPIRequest 查询用于电商券履约单合同下载的临时访问凭证 API请求
// tmall.alihouse.trade.coupon.order.contract.credentials.query
//
// 获取用于下载合同的临时aksk和安全token
type TmallalihousetradecouponordercontractcredentialsqueryAPIRequest struct {
	model.Params
}

// NewTmallalihousetradecouponordercontractcredentialsqueryRequest 初始化TmallalihousetradecouponordercontractcredentialsqueryAPIRequest对象
func NewTmallalihousetradecouponordercontractcredentialsqueryRequest() *TmallalihousetradecouponordercontractcredentialsqueryAPIRequest {
	return &TmallalihousetradecouponordercontractcredentialsqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallalihousetradecouponordercontractcredentialsqueryAPIRequest) GetApiMethodName() string {
	return "tmall.alihouse.trade.coupon.order.contract.credentials.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallalihousetradecouponordercontractcredentialsqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallalihousetradecouponordercontractcredentialsqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}
