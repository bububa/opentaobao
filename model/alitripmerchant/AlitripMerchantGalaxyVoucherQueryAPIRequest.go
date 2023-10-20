package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyvoucherqueryAPIRequest 查询单个代金券信息 API请求
// alitrip.merchant.galaxy.voucher.query
//
// 查询单个代金券信息
type AlitripmerchantgalaxyvoucherqueryAPIRequest struct {
	model.Params
	// 11
	_tenantKey string
	// 11
	_token string
	// 11
	_voucherId string
}

// NewAlitripmerchantgalaxyvoucherqueryRequest 初始化AlitripmerchantgalaxyvoucherqueryAPIRequest对象
func NewAlitripmerchantgalaxyvoucherqueryRequest() *AlitripmerchantgalaxyvoucherqueryAPIRequest {
	return &AlitripmerchantgalaxyvoucherqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyvoucherqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.voucher.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyvoucherqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyvoucherqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 11
func (r *AlitripmerchantgalaxyvoucherqueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyvoucherqueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 11
func (r *AlitripmerchantgalaxyvoucherqueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyvoucherqueryAPIRequest) GetToken() string {
	return r._token
}

// SetVoucherId is VoucherId Setter
// 11
func (r *AlitripmerchantgalaxyvoucherqueryAPIRequest) SetVoucherId(_voucherId string) error {
	r._voucherId = _voucherId
	r.Set("voucher_id", _voucherId)
	return nil
}

// GetVoucherId VoucherId Getter
func (r AlitripmerchantgalaxyvoucherqueryAPIRequest) GetVoucherId() string {
	return r._voucherId
}
