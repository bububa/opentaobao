package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxywechatcardqueryrecordAPIRequest 微信会员卡领取记录查询 API请求
// alitrip.merchant.galaxy.wechat.card.query.record
//
// 微信会员卡领取记录查询
type AlitripmerchantgalaxywechatcardqueryrecordAPIRequest struct {
	model.Params
	// 租户Id
	_tenantKey string
	// 用户登录凭证
	_token string
}

// NewAlitripmerchantgalaxywechatcardqueryrecordRequest 初始化AlitripmerchantgalaxywechatcardqueryrecordAPIRequest对象
func NewAlitripmerchantgalaxywechatcardqueryrecordRequest() *AlitripmerchantgalaxywechatcardqueryrecordAPIRequest {
	return &AlitripmerchantgalaxywechatcardqueryrecordAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxywechatcardqueryrecordAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.wechat.card.query.record"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxywechatcardqueryrecordAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxywechatcardqueryrecordAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户Id
func (r *AlitripmerchantgalaxywechatcardqueryrecordAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxywechatcardqueryrecordAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录凭证
func (r *AlitripmerchantgalaxywechatcardqueryrecordAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxywechatcardqueryrecordAPIRequest) GetToken() string {
	return r._token
}
