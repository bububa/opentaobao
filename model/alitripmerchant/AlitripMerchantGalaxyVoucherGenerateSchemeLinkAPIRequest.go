package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyvouchergenerateschemelinkAPIRequest 生成短信链接 API请求
// alitrip.merchant.galaxy.voucher.generate.scheme.link
//
// 生成微信跳转链接scheme_link
type AlitripmerchantgalaxyvouchergenerateschemelinkAPIRequest struct {
	model.Params
	// tenant_key
	_tenantKey string
	// 1111
	_token string
	// 加密merchantId
	_encryptionMerchantId string
}

// NewAlitripmerchantgalaxyvouchergenerateschemelinkRequest 初始化AlitripmerchantgalaxyvouchergenerateschemelinkAPIRequest对象
func NewAlitripmerchantgalaxyvouchergenerateschemelinkRequest() *AlitripmerchantgalaxyvouchergenerateschemelinkAPIRequest {
	return &AlitripmerchantgalaxyvouchergenerateschemelinkAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyvouchergenerateschemelinkAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.voucher.generate.scheme.link"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyvouchergenerateschemelinkAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyvouchergenerateschemelinkAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// tenant_key
func (r *AlitripmerchantgalaxyvouchergenerateschemelinkAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyvouchergenerateschemelinkAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 1111
func (r *AlitripmerchantgalaxyvouchergenerateschemelinkAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyvouchergenerateschemelinkAPIRequest) GetToken() string {
	return r._token
}

// SetEncryptionMerchantId is EncryptionMerchantId Setter
// 加密merchantId
func (r *AlitripmerchantgalaxyvouchergenerateschemelinkAPIRequest) SetEncryptionMerchantId(_encryptionMerchantId string) error {
	r._encryptionMerchantId = _encryptionMerchantId
	r.Set("encryption_merchant_id", _encryptionMerchantId)
	return nil
}

// GetEncryptionMerchantId EncryptionMerchantId Getter
func (r AlitripmerchantgalaxyvouchergenerateschemelinkAPIRequest) GetEncryptionMerchantId() string {
	return r._encryptionMerchantId
}
