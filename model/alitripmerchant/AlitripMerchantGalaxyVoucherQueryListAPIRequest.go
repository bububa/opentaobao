package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyVoucherQueryListAPIRequest 查询代金券列表 API请求
// alitrip.merchant.galaxy.voucher.query.list
//
// 查询代金券列表
type AlitripMerchantGalaxyVoucherQueryListAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 用户token
	_token string
}

// NewAlitripMerchantGalaxyVoucherQueryListRequest 初始化AlitripMerchantGalaxyVoucherQueryListAPIRequest对象
func NewAlitripMerchantGalaxyVoucherQueryListRequest() *AlitripMerchantGalaxyVoucherQueryListAPIRequest {
	return &AlitripMerchantGalaxyVoucherQueryListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyVoucherQueryListAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.voucher.query.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyVoucherQueryListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripMerchantGalaxyVoucherQueryListAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyVoucherQueryListAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripMerchantGalaxyVoucherQueryListAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyVoucherQueryListAPIRequest) GetToken() string {
	return r._token
}
