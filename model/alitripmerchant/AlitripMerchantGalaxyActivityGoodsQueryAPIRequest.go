package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyActivityGoodsQueryAPIRequest 营销抽奖-用户奖品查询 API请求
// alitrip.merchant.galaxy.activity.goods.query
//
// 星河产品-提供营销抽奖奖品查询服务
type AlitripMerchantGalaxyActivityGoodsQueryAPIRequest struct {
	model.Params
	// 商家租户id
	_tenantKey string
	// 用户登录标识
	_token string
}

// NewAlitripMerchantGalaxyActivityGoodsQueryRequest 初始化AlitripMerchantGalaxyActivityGoodsQueryAPIRequest对象
func NewAlitripMerchantGalaxyActivityGoodsQueryRequest() *AlitripMerchantGalaxyActivityGoodsQueryAPIRequest {
	return &AlitripMerchantGalaxyActivityGoodsQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyActivityGoodsQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.activity.goods.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyActivityGoodsQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTenantKey is TenantKey Setter
// 商家租户id
func (r *AlitripMerchantGalaxyActivityGoodsQueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyActivityGoodsQueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录标识
func (r *AlitripMerchantGalaxyActivityGoodsQueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyActivityGoodsQueryAPIRequest) GetToken() string {
	return r._token
}
