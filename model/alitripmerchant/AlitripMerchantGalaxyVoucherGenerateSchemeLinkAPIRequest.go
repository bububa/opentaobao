package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIRequest 生成短信链接 API请求
// alitrip.merchant.galaxy.voucher.generate.scheme.link
//
// 生成微信跳转链接scheme_link
type AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIRequest struct {
	model.Params
	// tenant_key
	_tenantKey string
	// 1111
	_token string
	// 加密merchantId
	_encryptionMerchantId string
}

// NewAlitripMerchantGalaxyVoucherGenerateSchemeLinkRequest 初始化AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIRequest对象
func NewAlitripMerchantGalaxyVoucherGenerateSchemeLinkRequest() *AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIRequest {
	return &AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.voucher.generate.scheme.link"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// tenant_key
func (r *AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 1111
func (r *AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIRequest) GetToken() string {
	return r._token
}

// SetEncryptionMerchantId is EncryptionMerchantId Setter
// 加密merchantId
func (r *AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIRequest) SetEncryptionMerchantId(_encryptionMerchantId string) error {
	r._encryptionMerchantId = _encryptionMerchantId
	r.Set("encryption_merchant_id", _encryptionMerchantId)
	return nil
}

// GetEncryptionMerchantId EncryptionMerchantId Getter
func (r AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIRequest) GetEncryptionMerchantId() string {
	return r._encryptionMerchantId
}
