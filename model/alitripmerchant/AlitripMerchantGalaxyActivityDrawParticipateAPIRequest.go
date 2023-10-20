package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyactivitydrawparticipateAPIRequest 星河-幸运抽奖活动参与 API请求
// alitrip.merchant.galaxy.activity.draw.participate
//
// 雅高小程序幸运抽奖活动页面用户进行抽奖，根据活动规则返回抽奖结果
type AlitripmerchantgalaxyactivitydrawparticipateAPIRequest struct {
	model.Params
	// 商家租户id
	_tenantKey string
	// 用户登录信息
	_token string
	// 活动id
	_offerId int64
}

// NewAlitripmerchantgalaxyactivitydrawparticipateRequest 初始化AlitripmerchantgalaxyactivitydrawparticipateAPIRequest对象
func NewAlitripmerchantgalaxyactivitydrawparticipateRequest() *AlitripmerchantgalaxyactivitydrawparticipateAPIRequest {
	return &AlitripmerchantgalaxyactivitydrawparticipateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyactivitydrawparticipateAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.activity.draw.participate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyactivitydrawparticipateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyactivitydrawparticipateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 商家租户id
func (r *AlitripmerchantgalaxyactivitydrawparticipateAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyactivitydrawparticipateAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录信息
func (r *AlitripmerchantgalaxyactivitydrawparticipateAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyactivitydrawparticipateAPIRequest) GetToken() string {
	return r._token
}

// SetOfferId is OfferId Setter
// 活动id
func (r *AlitripmerchantgalaxyactivitydrawparticipateAPIRequest) SetOfferId(_offerId int64) error {
	r._offerId = _offerId
	r.Set("offer_id", _offerId)
	return nil
}

// GetOfferId OfferId Getter
func (r AlitripmerchantgalaxyactivitydrawparticipateAPIRequest) GetOfferId() int64 {
	return r._offerId
}
