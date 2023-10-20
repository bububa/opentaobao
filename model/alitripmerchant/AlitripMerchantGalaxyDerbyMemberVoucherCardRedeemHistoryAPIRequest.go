package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervouchercardredeemhistoryAPIRequest 查询会员兑换臻享卡历史记录 API请求
// alitrip.merchant.galaxy.derby.member.voucher.card.redeem.history
//
// 查询会员兑换臻享卡历史记录
type AlitripmerchantgalaxyderbymembervouchercardredeemhistoryAPIRequest struct {
	model.Params
	// 租户ID
	_tenantKey string
	// 用户标识
	_token string
}

// NewAlitripmerchantgalaxyderbymembervouchercardredeemhistoryRequest 初始化AlitripmerchantgalaxyderbymembervouchercardredeemhistoryAPIRequest对象
func NewAlitripmerchantgalaxyderbymembervouchercardredeemhistoryRequest() *AlitripmerchantgalaxyderbymembervouchercardredeemhistoryAPIRequest {
	return &AlitripmerchantgalaxyderbymembervouchercardredeemhistoryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyderbymembervouchercardredeemhistoryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.card.redeem.history"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyderbymembervouchercardredeemhistoryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyderbymembervouchercardredeemhistoryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户ID
func (r *AlitripmerchantgalaxyderbymembervouchercardredeemhistoryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyderbymembervouchercardredeemhistoryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户标识
func (r *AlitripmerchantgalaxyderbymembervouchercardredeemhistoryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyderbymembervouchercardredeemhistoryAPIRequest) GetToken() string {
	return r._token
}
