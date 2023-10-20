package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyWechatCardParmQueryAPIRequest 微信会员卡添加 API请求
// alitrip.merchant.galaxy.wechat.card.parm.query
//
// 微信会员卡添加参数获取
type AlitripMerchantGalaxyWechatCardParmQueryAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 用户登录认证
	_token string
	// 用户在公众号下的标识信息
	_code string
}

// NewAlitripMerchantGalaxyWechatCardParmQueryRequest 初始化AlitripMerchantGalaxyWechatCardParmQueryAPIRequest对象
func NewAlitripMerchantGalaxyWechatCardParmQueryRequest() *AlitripMerchantGalaxyWechatCardParmQueryAPIRequest {
	return &AlitripMerchantGalaxyWechatCardParmQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyWechatCardParmQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.wechat.card.parm.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyWechatCardParmQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyWechatCardParmQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripMerchantGalaxyWechatCardParmQueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyWechatCardParmQueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录认证
func (r *AlitripMerchantGalaxyWechatCardParmQueryAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyWechatCardParmQueryAPIRequest) GetToken() string {
	return r._token
}

// SetCode is Code Setter
// 用户在公众号下的标识信息
func (r *AlitripMerchantGalaxyWechatCardParmQueryAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlitripMerchantGalaxyWechatCardParmQueryAPIRequest) GetCode() string {
	return r._code
}
