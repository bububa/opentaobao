package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbyvouchercardunlimitedchangecallbackAPIRequest 德比无限次券核销通知接口 API请求
// alitrip.merchant.galaxy.derby.voucher.card.unlimited.change.callback
//
// 德比无限次券核销通知接口
type AlitripmerchantgalaxyderbyvouchercardunlimitedchangecallbackAPIRequest struct {
	model.Params
	// 1
	_tenantKey string
	// 方法标识
	_derbyMethod string
	// 德比回调json信息
	_derbyBody string
}

// NewAlitripmerchantgalaxyderbyvouchercardunlimitedchangecallbackRequest 初始化AlitripmerchantgalaxyderbyvouchercardunlimitedchangecallbackAPIRequest对象
func NewAlitripmerchantgalaxyderbyvouchercardunlimitedchangecallbackRequest() *AlitripmerchantgalaxyderbyvouchercardunlimitedchangecallbackAPIRequest {
	return &AlitripmerchantgalaxyderbyvouchercardunlimitedchangecallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyderbyvouchercardunlimitedchangecallbackAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.voucher.card.unlimited.change.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyderbyvouchercardunlimitedchangecallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyderbyvouchercardunlimitedchangecallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 1
func (r *AlitripmerchantgalaxyderbyvouchercardunlimitedchangecallbackAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyderbyvouchercardunlimitedchangecallbackAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetDerbyMethod is DerbyMethod Setter
// 方法标识
func (r *AlitripmerchantgalaxyderbyvouchercardunlimitedchangecallbackAPIRequest) SetDerbyMethod(_derbyMethod string) error {
	r._derbyMethod = _derbyMethod
	r.Set("derby_method", _derbyMethod)
	return nil
}

// GetDerbyMethod DerbyMethod Getter
func (r AlitripmerchantgalaxyderbyvouchercardunlimitedchangecallbackAPIRequest) GetDerbyMethod() string {
	return r._derbyMethod
}

// SetDerbyBody is DerbyBody Setter
// 德比回调json信息
func (r *AlitripmerchantgalaxyderbyvouchercardunlimitedchangecallbackAPIRequest) SetDerbyBody(_derbyBody string) error {
	r._derbyBody = _derbyBody
	r.Set("derby_body", _derbyBody)
	return nil
}

// GetDerbyBody DerbyBody Getter
func (r AlitripmerchantgalaxyderbyvouchercardunlimitedchangecallbackAPIRequest) GetDerbyBody() string {
	return r._derbyBody
}
