package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyqueryparticipatenumberAPIRequest 星河-抽奖活动次数查询 API请求
// alitrip.merchant.galaxy.query.participate.number
//
// 雅高小程序抽奖活动，查询用户抽奖次数
type AlitripmerchantgalaxyqueryparticipatenumberAPIRequest struct {
	model.Params
	// 商家租户id
	_tenantKey string
	// 用户登录信息
	_token string
	// 活动id
	_offerId int64
}

// NewAlitripmerchantgalaxyqueryparticipatenumberRequest 初始化AlitripmerchantgalaxyqueryparticipatenumberAPIRequest对象
func NewAlitripmerchantgalaxyqueryparticipatenumberRequest() *AlitripmerchantgalaxyqueryparticipatenumberAPIRequest {
	return &AlitripmerchantgalaxyqueryparticipatenumberAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyqueryparticipatenumberAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.query.participate.number"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyqueryparticipatenumberAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyqueryparticipatenumberAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 商家租户id
func (r *AlitripmerchantgalaxyqueryparticipatenumberAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyqueryparticipatenumberAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录信息
func (r *AlitripmerchantgalaxyqueryparticipatenumberAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyqueryparticipatenumberAPIRequest) GetToken() string {
	return r._token
}

// SetOfferId is OfferId Setter
// 活动id
func (r *AlitripmerchantgalaxyqueryparticipatenumberAPIRequest) SetOfferId(_offerId int64) error {
	r._offerId = _offerId
	r.Set("offer_id", _offerId)
	return nil
}

// GetOfferId OfferId Getter
func (r AlitripmerchantgalaxyqueryparticipatenumberAPIRequest) GetOfferId() int64 {
	return r._offerId
}
