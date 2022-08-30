package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyVoucherQueryAPIRequest 查询单个代金券信息 API请求
// alitrip.merchant.galaxy.voucher.query
//
// 查询单个代金券信息
type AlitripMerchantGalaxyVoucherQueryAPIRequest struct {
	model.Params
	// 11
	_tenantKey string
	// 11
	_token string
	// 11
	_voucherId string
}

// NewAlitripMerchantGalaxyVoucherQueryRequest 初始化AlitripMerchantGalaxyVoucherQueryAPIRequest对象
func NewAlitripMerchantGalaxyVoucherQueryRequest() *AlitripMerchantGalaxyVoucherQueryAPIRequest {
	return &AlitripMerchantGalaxyVoucherQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyVoucherQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.voucher.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyVoucherQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTenantKey is TenantKey Setter
// 11
func (r *AlitripMerchantGalaxyVoucherQueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyVoucherQueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 11
func (r *AlitripMerchantGalaxyVoucherQueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyVoucherQueryAPIRequest) GetToken() string {
	return r._token
}

// SetVoucherId is VoucherId Setter
// 11
func (r *AlitripMerchantGalaxyVoucherQueryAPIRequest) SetVoucherId(_voucherId string) error {
	r._voucherId = _voucherId
	r.Set("voucher_id", _voucherId)
	return nil
}

// GetVoucherId VoucherId Getter
func (r AlitripMerchantGalaxyVoucherQueryAPIRequest) GetVoucherId() string {
	return r._voucherId
}
