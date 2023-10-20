package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatxcsbrandmarketingcouponqrcodegetAPIRequest 品牌营销导购员券页面二维码获取 API请求
// alibaba.txcs.brandmarketing.coupon.qrcode.get
//
// 构建券页码二维码url
type AlibabatxcsbrandmarketingcouponqrcodegetAPIRequest struct {
	model.Params
	// 请求信息
	_couponQrcodeParamDo *CouponQrcodeParamDo
}

// NewAlibabatxcsbrandmarketingcouponqrcodegetRequest 初始化AlibabatxcsbrandmarketingcouponqrcodegetAPIRequest对象
func NewAlibabatxcsbrandmarketingcouponqrcodegetRequest() *AlibabatxcsbrandmarketingcouponqrcodegetAPIRequest {
	return &AlibabatxcsbrandmarketingcouponqrcodegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatxcsbrandmarketingcouponqrcodegetAPIRequest) GetApiMethodName() string {
	return "alibaba.txcs.brandmarketing.coupon.qrcode.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatxcsbrandmarketingcouponqrcodegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatxcsbrandmarketingcouponqrcodegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCouponQrcodeParamDo is CouponQrcodeParamDo Setter
// 请求信息
func (r *AlibabatxcsbrandmarketingcouponqrcodegetAPIRequest) SetCouponQrcodeParamDo(_couponQrcodeParamDo *CouponQrcodeParamDo) error {
	r._couponQrcodeParamDo = _couponQrcodeParamDo
	r.Set("coupon_qrcode_param_do", _couponQrcodeParamDo)
	return nil
}

// GetCouponQrcodeParamDo CouponQrcodeParamDo Getter
func (r AlibabatxcsbrandmarketingcouponqrcodegetAPIRequest) GetCouponQrcodeParamDo() *CouponQrcodeParamDo {
	return r._couponQrcodeParamDo
}
