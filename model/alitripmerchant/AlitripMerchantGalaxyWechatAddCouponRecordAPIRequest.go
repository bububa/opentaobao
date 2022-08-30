package alitripmerchant

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.wechat.add.coupon.record"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyWechatAddCouponRecordAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
