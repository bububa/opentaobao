package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxycommonbindmerchantidAPIRequest 绑定用户和merchantID API请求
// alitrip.merchant.galaxy.common.bind.merchant.id
//
// 绑定用户和merchantID
type AlitripmerchantgalaxycommonbindmerchantidAPIRequest struct {
	model.Params
	// tenant_key
	_tenantKey string
	// 状态
	_bindingMerchantIdUserDto *BindingMerchantIdUserDto
}

// NewAlitripmerchantgalaxycommonbindmerchantidRequest 初始化AlitripmerchantgalaxycommonbindmerchantidAPIRequest对象
func NewAlitripmerchantgalaxycommonbindmerchantidRequest() *AlitripmerchantgalaxycommonbindmerchantidAPIRequest {
	return &AlitripmerchantgalaxycommonbindmerchantidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxycommonbindmerchantidAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.common.bind.merchant.id"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxycommonbindmerchantidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxycommonbindmerchantidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// tenant_key
func (r *AlitripmerchantgalaxycommonbindmerchantidAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxycommonbindmerchantidAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetBindingMerchantIdUserDto is BindingMerchantIdUserDto Setter
// 状态
func (r *AlitripmerchantgalaxycommonbindmerchantidAPIRequest) SetBindingMerchantIdUserDto(_bindingMerchantIdUserDto *BindingMerchantIdUserDto) error {
	r._bindingMerchantIdUserDto = _bindingMerchantIdUserDto
	r.Set("binding_merchant_id_user_dto", _bindingMerchantIdUserDto)
	return nil
}

// GetBindingMerchantIdUserDto BindingMerchantIdUserDto Getter
func (r AlitripmerchantgalaxycommonbindmerchantidAPIRequest) GetBindingMerchantIdUserDto() *BindingMerchantIdUserDto {
	return r._bindingMerchantIdUserDto
}
