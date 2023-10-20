package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyActivityFatigueAPIRequest 营销抽奖-弹窗疲劳度控制 API请求
// alitrip.merchant.galaxy.activity.fatigue
//
// 星河产品-营销抽奖首页弹窗疲劳度控制服务
type AlitripMerchantGalaxyActivityFatigueAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// 用户登录标识
	_token string
	// 活动ID
	_offerId int64
}

// NewAlitripMerchantGalaxyActivityFatigueRequest 初始化AlitripMerchantGalaxyActivityFatigueAPIRequest对象
func NewAlitripMerchantGalaxyActivityFatigueRequest() *AlitripMerchantGalaxyActivityFatigueAPIRequest {
	return &AlitripMerchantGalaxyActivityFatigueAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyActivityFatigueAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._offerId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyActivityFatigueAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.activity.fatigue"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyActivityFatigueAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyActivityFatigueAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户身份信息
func (r *AlitripMerchantGalaxyActivityFatigueAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyActivityFatigueAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录标识
func (r *AlitripMerchantGalaxyActivityFatigueAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyActivityFatigueAPIRequest) GetToken() string {
	return r._token
}

// SetOfferId is OfferId Setter
// 活动ID
func (r *AlitripMerchantGalaxyActivityFatigueAPIRequest) SetOfferId(_offerId int64) error {
	r._offerId = _offerId
	r.Set("offer_id", _offerId)
	return nil
}

// GetOfferId OfferId Getter
func (r AlitripMerchantGalaxyActivityFatigueAPIRequest) GetOfferId() int64 {
	return r._offerId
}

var poolAlitripMerchantGalaxyActivityFatigueAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyActivityFatigueRequest()
	},
}

// GetAlitripMerchantGalaxyActivityFatigueRequest 从 sync.Pool 获取 AlitripMerchantGalaxyActivityFatigueAPIRequest
func GetAlitripMerchantGalaxyActivityFatigueAPIRequest() *AlitripMerchantGalaxyActivityFatigueAPIRequest {
	return poolAlitripMerchantGalaxyActivityFatigueAPIRequest.Get().(*AlitripMerchantGalaxyActivityFatigueAPIRequest)
}

// ReleaseAlitripMerchantGalaxyActivityFatigueAPIRequest 将 AlitripMerchantGalaxyActivityFatigueAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyActivityFatigueAPIRequest(v *AlitripMerchantGalaxyActivityFatigueAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyActivityFatigueAPIRequest.Put(v)
}
