package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest 星河-记录用户微信优惠券领取记录 API请求
// alitrip.merchant.galaxy.wechat.add.coupon.record
//
// 雅高小程序添加优惠券实例
type AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 用户token
	_token string
	// 卡类型
	_cardType string
	// 实体类
	_sendCouponResult string
}

// NewAlitripMerchantGalaxyWechatAddCouponRecordRequest 初始化AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest对象
func NewAlitripMerchantGalaxyWechatAddCouponRecordRequest() *AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest {
	return &AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._cardType = ""
	r._sendCouponResult = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.wechat.add.coupon.record"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest) GetToken() string {
	return r._token
}

// SetCardType is CardType Setter
// 卡类型
func (r *AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest) SetCardType(_cardType string) error {
	r._cardType = _cardType
	r.Set("card_type", _cardType)
	return nil
}

// GetCardType CardType Getter
func (r AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest) GetCardType() string {
	return r._cardType
}

// SetSendCouponResult is SendCouponResult Setter
// 实体类
func (r *AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest) SetSendCouponResult(_sendCouponResult string) error {
	r._sendCouponResult = _sendCouponResult
	r.Set("send_coupon_result", _sendCouponResult)
	return nil
}

// GetSendCouponResult SendCouponResult Getter
func (r AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest) GetSendCouponResult() string {
	return r._sendCouponResult
}

var poolAlitripMerchantGalaxyWechatAddCouponRecordAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyWechatAddCouponRecordRequest()
	},
}

// GetAlitripMerchantGalaxyWechatAddCouponRecordRequest 从 sync.Pool 获取 AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest
func GetAlitripMerchantGalaxyWechatAddCouponRecordAPIRequest() *AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest {
	return poolAlitripMerchantGalaxyWechatAddCouponRecordAPIRequest.Get().(*AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest)
}

// ReleaseAlitripMerchantGalaxyWechatAddCouponRecordAPIRequest 将 AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyWechatAddCouponRecordAPIRequest(v *AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyWechatAddCouponRecordAPIRequest.Put(v)
}
