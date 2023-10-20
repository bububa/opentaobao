package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest 品牌营销导购员券页面二维码获取 API请求
// alibaba.txcs.brandmarketing.coupon.qrcode.get
//
// 构建券页码二维码url
type AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest struct {
	model.Params
	// 请求信息
	_couponQrcodeParamDo *CouponQrcodeParamDo
}

// NewAlibabaTxcsBrandmarketingCouponQrcodeGetRequest 初始化AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest对象
func NewAlibabaTxcsBrandmarketingCouponQrcodeGetRequest() *AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest {
	return &AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest) Reset() {
	r._couponQrcodeParamDo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest) GetApiMethodName() string {
	return "alibaba.txcs.brandmarketing.coupon.qrcode.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCouponQrcodeParamDo is CouponQrcodeParamDo Setter
// 请求信息
func (r *AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest) SetCouponQrcodeParamDo(_couponQrcodeParamDo *CouponQrcodeParamDo) error {
	r._couponQrcodeParamDo = _couponQrcodeParamDo
	r.Set("coupon_qrcode_param_do", _couponQrcodeParamDo)
	return nil
}

// GetCouponQrcodeParamDo CouponQrcodeParamDo Getter
func (r AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest) GetCouponQrcodeParamDo() *CouponQrcodeParamDo {
	return r._couponQrcodeParamDo
}

var poolAlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTxcsBrandmarketingCouponQrcodeGetRequest()
	},
}

// GetAlibabaTxcsBrandmarketingCouponQrcodeGetRequest 从 sync.Pool 获取 AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest
func GetAlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest() *AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest {
	return poolAlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest.Get().(*AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest)
}

// ReleaseAlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest 将 AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest(v *AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest) {
	v.Reset()
	poolAlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest.Put(v)
}
