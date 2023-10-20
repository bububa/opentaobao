package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxywechataddcouponrecordAPIRequest 星河-记录用户微信优惠券领取记录 API请求
// alitrip.merchant.galaxy.wechat.add.coupon.record
//
// 雅高小程序添加优惠券实例
type AlitripmerchantgalaxywechataddcouponrecordAPIRequest struct {
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

// NewAlitripmerchantgalaxywechataddcouponrecordRequest 初始化AlitripmerchantgalaxywechataddcouponrecordAPIRequest对象
func NewAlitripmerchantgalaxywechataddcouponrecordRequest() *AlitripmerchantgalaxywechataddcouponrecordAPIRequest {
	return &AlitripmerchantgalaxywechataddcouponrecordAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxywechataddcouponrecordAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.wechat.add.coupon.record"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxywechataddcouponrecordAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxywechataddcouponrecordAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripmerchantgalaxywechataddcouponrecordAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxywechataddcouponrecordAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripmerchantgalaxywechataddcouponrecordAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxywechataddcouponrecordAPIRequest) GetToken() string {
	return r._token
}

// SetCardType is CardType Setter
// 卡类型
func (r *AlitripmerchantgalaxywechataddcouponrecordAPIRequest) SetCardType(_cardType string) error {
	r._cardType = _cardType
	r.Set("card_type", _cardType)
	return nil
}

// GetCardType CardType Getter
func (r AlitripmerchantgalaxywechataddcouponrecordAPIRequest) GetCardType() string {
	return r._cardType
}

// SetSendCouponResult is SendCouponResult Setter
// 实体类
func (r *AlitripmerchantgalaxywechataddcouponrecordAPIRequest) SetSendCouponResult(_sendCouponResult string) error {
	r._sendCouponResult = _sendCouponResult
	r.Set("send_coupon_result", _sendCouponResult)
	return nil
}

// GetSendCouponResult SendCouponResult Getter
func (r AlitripmerchantgalaxywechataddcouponrecordAPIRequest) GetSendCouponResult() string {
	return r._sendCouponResult
}
