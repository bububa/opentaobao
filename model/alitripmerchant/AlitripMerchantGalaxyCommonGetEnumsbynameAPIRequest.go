package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest 小程序公共枚举类获取接口 API请求
// alitrip.merchant.galaxy.common.get.enumsbyname
//
// 通过枚举名称获取枚举类实例内容
type AlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 枚举名
	_enumClassName string
}

// NewAlitripMerchantGalaxyCommonGetEnumsbynameRequest 初始化AlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest对象
func NewAlitripMerchantGalaxyCommonGetEnumsbynameRequest() *AlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest {
	return &AlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest) Reset() {
	r._tenantKey = ""
	r._enumClassName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.common.get.enumsbyname"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetEnumClassName is EnumClassName Setter
// 枚举名
func (r *AlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest) SetEnumClassName(_enumClassName string) error {
	r._enumClassName = _enumClassName
	r.Set("enum_class_name", _enumClassName)
	return nil
}

// GetEnumClassName EnumClassName Getter
func (r AlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest) GetEnumClassName() string {
	return r._enumClassName
}

var poolAlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyCommonGetEnumsbynameRequest()
	},
}

// GetAlitripMerchantGalaxyCommonGetEnumsbynameRequest 从 sync.Pool 获取 AlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest
func GetAlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest() *AlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest {
	return poolAlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest.Get().(*AlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest)
}

// ReleaseAlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest 将 AlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest(v *AlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyCommonGetEnumsbynameAPIRequest.Put(v)
}
