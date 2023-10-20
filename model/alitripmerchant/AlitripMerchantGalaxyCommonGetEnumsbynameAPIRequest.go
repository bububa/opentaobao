package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxycommongetenumsbynameAPIRequest 小程序公共枚举类获取接口 API请求
// alitrip.merchant.galaxy.common.get.enumsbyname
//
// 通过枚举名称获取枚举类实例内容
type AlitripmerchantgalaxycommongetenumsbynameAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 枚举名
	_enumClassName string
}

// NewAlitripmerchantgalaxycommongetenumsbynameRequest 初始化AlitripmerchantgalaxycommongetenumsbynameAPIRequest对象
func NewAlitripmerchantgalaxycommongetenumsbynameRequest() *AlitripmerchantgalaxycommongetenumsbynameAPIRequest {
	return &AlitripmerchantgalaxycommongetenumsbynameAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxycommongetenumsbynameAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.common.get.enumsbyname"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxycommongetenumsbynameAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxycommongetenumsbynameAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripmerchantgalaxycommongetenumsbynameAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxycommongetenumsbynameAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetEnumClassName is EnumClassName Setter
// 枚举名
func (r *AlitripmerchantgalaxycommongetenumsbynameAPIRequest) SetEnumClassName(_enumClassName string) error {
	r._enumClassName = _enumClassName
	r.Set("enum_class_name", _enumClassName)
	return nil
}

// GetEnumClassName EnumClassName Getter
func (r AlitripmerchantgalaxycommongetenumsbynameAPIRequest) GetEnumClassName() string {
	return r._enumClassName
}
