package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyActivityPopupControlAPIRequest 营销弹屏疲劳度控制 API请求
// alitrip.merchant.galaxy.activity.popup.control
//
// 星河=营销弹屏疲劳控制
type AlitripMerchantGalaxyActivityPopupControlAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 用户token
	_token string
	// 活动类型
	_type string
	// 活动id
	_offerId int64
}

// NewAlitripMerchantGalaxyActivityPopupControlRequest 初始化AlitripMerchantGalaxyActivityPopupControlAPIRequest对象
func NewAlitripMerchantGalaxyActivityPopupControlRequest() *AlitripMerchantGalaxyActivityPopupControlAPIRequest {
	return &AlitripMerchantGalaxyActivityPopupControlAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyActivityPopupControlAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._type = ""
	r._offerId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyActivityPopupControlAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.activity.popup.control"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyActivityPopupControlAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyActivityPopupControlAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripMerchantGalaxyActivityPopupControlAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyActivityPopupControlAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripMerchantGalaxyActivityPopupControlAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyActivityPopupControlAPIRequest) GetToken() string {
	return r._token
}

// SetType is Type Setter
// 活动类型
func (r *AlitripMerchantGalaxyActivityPopupControlAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlitripMerchantGalaxyActivityPopupControlAPIRequest) GetType() string {
	return r._type
}

// SetOfferId is OfferId Setter
// 活动id
func (r *AlitripMerchantGalaxyActivityPopupControlAPIRequest) SetOfferId(_offerId int64) error {
	r._offerId = _offerId
	r.Set("offer_id", _offerId)
	return nil
}

// GetOfferId OfferId Getter
func (r AlitripMerchantGalaxyActivityPopupControlAPIRequest) GetOfferId() int64 {
	return r._offerId
}

var poolAlitripMerchantGalaxyActivityPopupControlAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyActivityPopupControlRequest()
	},
}

// GetAlitripMerchantGalaxyActivityPopupControlRequest 从 sync.Pool 获取 AlitripMerchantGalaxyActivityPopupControlAPIRequest
func GetAlitripMerchantGalaxyActivityPopupControlAPIRequest() *AlitripMerchantGalaxyActivityPopupControlAPIRequest {
	return poolAlitripMerchantGalaxyActivityPopupControlAPIRequest.Get().(*AlitripMerchantGalaxyActivityPopupControlAPIRequest)
}

// ReleaseAlitripMerchantGalaxyActivityPopupControlAPIRequest 将 AlitripMerchantGalaxyActivityPopupControlAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyActivityPopupControlAPIRequest(v *AlitripMerchantGalaxyActivityPopupControlAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyActivityPopupControlAPIRequest.Put(v)
}
