package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyQueryParticipateNumberAPIRequest 星河-抽奖活动次数查询 API请求
// alitrip.merchant.galaxy.query.participate.number
//
// 雅高小程序抽奖活动，查询用户抽奖次数
type AlitripMerchantGalaxyQueryParticipateNumberAPIRequest struct {
	model.Params
	// 商家租户id
	_tenantKey string
	// 用户登录信息
	_token string
	// 活动id
	_offerId int64
}

// NewAlitripMerchantGalaxyQueryParticipateNumberRequest 初始化AlitripMerchantGalaxyQueryParticipateNumberAPIRequest对象
func NewAlitripMerchantGalaxyQueryParticipateNumberRequest() *AlitripMerchantGalaxyQueryParticipateNumberAPIRequest {
	return &AlitripMerchantGalaxyQueryParticipateNumberAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyQueryParticipateNumberAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._offerId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyQueryParticipateNumberAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.query.participate.number"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyQueryParticipateNumberAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyQueryParticipateNumberAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 商家租户id
func (r *AlitripMerchantGalaxyQueryParticipateNumberAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyQueryParticipateNumberAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录信息
func (r *AlitripMerchantGalaxyQueryParticipateNumberAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyQueryParticipateNumberAPIRequest) GetToken() string {
	return r._token
}

// SetOfferId is OfferId Setter
// 活动id
func (r *AlitripMerchantGalaxyQueryParticipateNumberAPIRequest) SetOfferId(_offerId int64) error {
	r._offerId = _offerId
	r.Set("offer_id", _offerId)
	return nil
}

// GetOfferId OfferId Getter
func (r AlitripMerchantGalaxyQueryParticipateNumberAPIRequest) GetOfferId() int64 {
	return r._offerId
}

var poolAlitripMerchantGalaxyQueryParticipateNumberAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyQueryParticipateNumberRequest()
	},
}

// GetAlitripMerchantGalaxyQueryParticipateNumberRequest 从 sync.Pool 获取 AlitripMerchantGalaxyQueryParticipateNumberAPIRequest
func GetAlitripMerchantGalaxyQueryParticipateNumberAPIRequest() *AlitripMerchantGalaxyQueryParticipateNumberAPIRequest {
	return poolAlitripMerchantGalaxyQueryParticipateNumberAPIRequest.Get().(*AlitripMerchantGalaxyQueryParticipateNumberAPIRequest)
}

// ReleaseAlitripMerchantGalaxyQueryParticipateNumberAPIRequest 将 AlitripMerchantGalaxyQueryParticipateNumberAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyQueryParticipateNumberAPIRequest(v *AlitripMerchantGalaxyQueryParticipateNumberAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyQueryParticipateNumberAPIRequest.Put(v)
}
