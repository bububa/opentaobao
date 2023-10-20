package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyActivityDrawParticipateAPIRequest 星河-幸运抽奖活动参与 API请求
// alitrip.merchant.galaxy.activity.draw.participate
//
// 雅高小程序幸运抽奖活动页面用户进行抽奖，根据活动规则返回抽奖结果
type AlitripMerchantGalaxyActivityDrawParticipateAPIRequest struct {
	model.Params
	// 商家租户id
	_tenantKey string
	// 用户登录信息
	_token string
	// 活动id
	_offerId int64
}

// NewAlitripMerchantGalaxyActivityDrawParticipateRequest 初始化AlitripMerchantGalaxyActivityDrawParticipateAPIRequest对象
func NewAlitripMerchantGalaxyActivityDrawParticipateRequest() *AlitripMerchantGalaxyActivityDrawParticipateAPIRequest {
	return &AlitripMerchantGalaxyActivityDrawParticipateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyActivityDrawParticipateAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._offerId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyActivityDrawParticipateAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.activity.draw.participate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyActivityDrawParticipateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyActivityDrawParticipateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 商家租户id
func (r *AlitripMerchantGalaxyActivityDrawParticipateAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyActivityDrawParticipateAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录信息
func (r *AlitripMerchantGalaxyActivityDrawParticipateAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyActivityDrawParticipateAPIRequest) GetToken() string {
	return r._token
}

// SetOfferId is OfferId Setter
// 活动id
func (r *AlitripMerchantGalaxyActivityDrawParticipateAPIRequest) SetOfferId(_offerId int64) error {
	r._offerId = _offerId
	r.Set("offer_id", _offerId)
	return nil
}

// GetOfferId OfferId Getter
func (r AlitripMerchantGalaxyActivityDrawParticipateAPIRequest) GetOfferId() int64 {
	return r._offerId
}

var poolAlitripMerchantGalaxyActivityDrawParticipateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyActivityDrawParticipateRequest()
	},
}

// GetAlitripMerchantGalaxyActivityDrawParticipateRequest 从 sync.Pool 获取 AlitripMerchantGalaxyActivityDrawParticipateAPIRequest
func GetAlitripMerchantGalaxyActivityDrawParticipateAPIRequest() *AlitripMerchantGalaxyActivityDrawParticipateAPIRequest {
	return poolAlitripMerchantGalaxyActivityDrawParticipateAPIRequest.Get().(*AlitripMerchantGalaxyActivityDrawParticipateAPIRequest)
}

// ReleaseAlitripMerchantGalaxyActivityDrawParticipateAPIRequest 将 AlitripMerchantGalaxyActivityDrawParticipateAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyActivityDrawParticipateAPIRequest(v *AlitripMerchantGalaxyActivityDrawParticipateAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyActivityDrawParticipateAPIRequest.Put(v)
}
