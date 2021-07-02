package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallPromotionCouponUserAPIRequest 用户信息查询接口 API请求
// tmall.promotion.coupon.user
//
// 开发给外部合作商（例如：苏宁），通过会员付款码获得会员nick
type TmallPromotionCouponUserAPIRequest struct {
	model.Params
	// 例如：suning
	_bizType string
	// 会员付款码
	_payCode string
	// 扩展字段
	_extra string
}

// NewTmallPromotionCouponUserRequest 初始化TmallPromotionCouponUserAPIRequest对象
func NewTmallPromotionCouponUserRequest() *TmallPromotionCouponUserAPIRequest {
	return &TmallPromotionCouponUserAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallPromotionCouponUserAPIRequest) GetApiMethodName() string {
	return "tmall.promotion.coupon.user"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallPromotionCouponUserAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBizType is BizType Setter
// 例如：suning
func (r *TmallPromotionCouponUserAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TmallPromotionCouponUserAPIRequest) GetBizType() string {
	return r._bizType
}

// SetPayCode is PayCode Setter
// 会员付款码
func (r *TmallPromotionCouponUserAPIRequest) SetPayCode(_payCode string) error {
	r._payCode = _payCode
	r.Set("pay_code", _payCode)
	return nil
}

// GetPayCode PayCode Getter
func (r TmallPromotionCouponUserAPIRequest) GetPayCode() string {
	return r._payCode
}

// SetExtra is Extra Setter
// 扩展字段
func (r *TmallPromotionCouponUserAPIRequest) SetExtra(_extra string) error {
	r._extra = _extra
	r.Set("extra", _extra)
	return nil
}

// GetExtra Extra Getter
func (r TmallPromotionCouponUserAPIRequest) GetExtra() string {
	return r._extra
}
