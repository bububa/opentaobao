package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyFavoriteListAPIRequest 用户收藏列表查询 API请求
// alitrip.merchant.galaxy.favorite.list
//
// 让用户可以查询自己收藏的酒店列表
type AlitripMerchantGalaxyFavoriteListAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 用户登录标识
	_token string
}

// NewAlitripMerchantGalaxyFavoriteListRequest 初始化AlitripMerchantGalaxyFavoriteListAPIRequest对象
func NewAlitripMerchantGalaxyFavoriteListRequest() *AlitripMerchantGalaxyFavoriteListAPIRequest {
	return &AlitripMerchantGalaxyFavoriteListAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyFavoriteListAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyFavoriteListAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.favorite.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyFavoriteListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyFavoriteListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripMerchantGalaxyFavoriteListAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyFavoriteListAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录标识
func (r *AlitripMerchantGalaxyFavoriteListAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyFavoriteListAPIRequest) GetToken() string {
	return r._token
}

var poolAlitripMerchantGalaxyFavoriteListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyFavoriteListRequest()
	},
}

// GetAlitripMerchantGalaxyFavoriteListRequest 从 sync.Pool 获取 AlitripMerchantGalaxyFavoriteListAPIRequest
func GetAlitripMerchantGalaxyFavoriteListAPIRequest() *AlitripMerchantGalaxyFavoriteListAPIRequest {
	return poolAlitripMerchantGalaxyFavoriteListAPIRequest.Get().(*AlitripMerchantGalaxyFavoriteListAPIRequest)
}

// ReleaseAlitripMerchantGalaxyFavoriteListAPIRequest 将 AlitripMerchantGalaxyFavoriteListAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyFavoriteListAPIRequest(v *AlitripMerchantGalaxyFavoriteListAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyFavoriteListAPIRequest.Put(v)
}
