package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyCommonBindMerchantIdAPIRequest 绑定用户和merchantID API请求
// alitrip.merchant.galaxy.common.bind.merchant.id
//
// 绑定用户和merchantID
type AlitripMerchantGalaxyCommonBindMerchantIdAPIRequest struct {
	model.Params
	// tenant_key
	_tenantKey string
	// 状态
	_bindingMerchantIdUserDto *BindingMerchantIdUserDto
}

// NewAlitripMerchantGalaxyCommonBindMerchantIdRequest 初始化AlitripMerchantGalaxyCommonBindMerchantIdAPIRequest对象
func NewAlitripMerchantGalaxyCommonBindMerchantIdRequest() *AlitripMerchantGalaxyCommonBindMerchantIdAPIRequest {
	return &AlitripMerchantGalaxyCommonBindMerchantIdAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyCommonBindMerchantIdAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.common.bind.merchant.id"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyCommonBindMerchantIdAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyCommonBindMerchantIdAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// tenant_key
func (r *AlitripMerchantGalaxyCommonBindMerchantIdAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyCommonBindMerchantIdAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetBindingMerchantIdUserDto is BindingMerchantIdUserDto Setter
// 状态
func (r *AlitripMerchantGalaxyCommonBindMerchantIdAPIRequest) SetBindingMerchantIdUserDto(_bindingMerchantIdUserDto *BindingMerchantIdUserDto) error {
	r._bindingMerchantIdUserDto = _bindingMerchantIdUserDto
	r.Set("binding_merchant_id_user_dto", _bindingMerchantIdUserDto)
	return nil
}

// GetBindingMerchantIdUserDto BindingMerchantIdUserDto Getter
func (r AlitripMerchantGalaxyCommonBindMerchantIdAPIRequest) GetBindingMerchantIdUserDto() *BindingMerchantIdUserDto {
	return r._bindingMerchantIdUserDto
}
