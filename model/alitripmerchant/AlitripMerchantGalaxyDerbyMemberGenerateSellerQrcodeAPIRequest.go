package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembergeneratesellerqrcodeAPIRequest 生成臻享卡德比分销二维码 API请求
// alitrip.merchant.galaxy.derby.member.generate.seller.qrcode
//
// 生成臻享卡德比分销二维码
type AlitripmerchantgalaxyderbymembergeneratesellerqrcodeAPIRequest struct {
	model.Params
	// tenant_key
	_tenantKey string
	// token
	_token string
	// 用户code
	_code string
}

// NewAlitripmerchantgalaxyderbymembergeneratesellerqrcodeRequest 初始化AlitripmerchantgalaxyderbymembergeneratesellerqrcodeAPIRequest对象
func NewAlitripmerchantgalaxyderbymembergeneratesellerqrcodeRequest() *AlitripmerchantgalaxyderbymembergeneratesellerqrcodeAPIRequest {
	return &AlitripmerchantgalaxyderbymembergeneratesellerqrcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyderbymembergeneratesellerqrcodeAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.generate.seller.qrcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyderbymembergeneratesellerqrcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyderbymembergeneratesellerqrcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// tenant_key
func (r *AlitripmerchantgalaxyderbymembergeneratesellerqrcodeAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyderbymembergeneratesellerqrcodeAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// token
func (r *AlitripmerchantgalaxyderbymembergeneratesellerqrcodeAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyderbymembergeneratesellerqrcodeAPIRequest) GetToken() string {
	return r._token
}

// SetCode is Code Setter
// 用户code
func (r *AlitripmerchantgalaxyderbymembergeneratesellerqrcodeAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlitripmerchantgalaxyderbymembergeneratesellerqrcodeAPIRequest) GetCode() string {
	return r._code
}
