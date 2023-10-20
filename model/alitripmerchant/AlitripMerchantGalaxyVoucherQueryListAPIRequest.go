package alitripmerchant

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyVoucherQueryListAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyVoucherQueryListAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.voucher.query.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyVoucherQueryListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyVoucherQueryListAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlitripMerchantGalaxyVoucherQueryListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyVoucherQueryListRequest()
	},
}

// GetAlitripMerchantGalaxyVoucherQueryListRequest 从 sync.Pool 获取 AlitripMerchantGalaxyVoucherQueryListAPIRequest
func GetAlitripMerchantGalaxyVoucherQueryListAPIRequest() *AlitripMerchantGalaxyVoucherQueryListAPIRequest {
	return poolAlitripMerchantGalaxyVoucherQueryListAPIRequest.Get().(*AlitripMerchantGalaxyVoucherQueryListAPIRequest)
}

// ReleaseAlitripMerchantGalaxyVoucherQueryListAPIRequest 将 AlitripMerchantGalaxyVoucherQueryListAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyVoucherQueryListAPIRequest(v *AlitripMerchantGalaxyVoucherQueryListAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyVoucherQueryListAPIRequest.Put(v)
}
