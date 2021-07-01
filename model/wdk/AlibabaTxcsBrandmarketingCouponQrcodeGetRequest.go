package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌营销导购员券页面二维码获取 API请求
alibaba.txcs.brandmarketing.coupon.qrcode.get

构建券页码二维码url
*/
type AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest struct {
    model.Params
    // 请求信息
    _couponQrcodeParamDo   *CouponQrcodeParamDO
}

// 初始化AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest对象
func NewAlibabaTxcsBrandmarketingCouponQrcodeGetRequest() *AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest{
    return &AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest) GetApiMethodName() string {
    return "alibaba.txcs.brandmarketing.coupon.qrcode.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CouponQrcodeParamDo Setter
// 请求信息
func (r *AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest) SetCouponQrcodeParamDo(_couponQrcodeParamDo *CouponQrcodeParamDO) error {
    r._couponQrcodeParamDo = _couponQrcodeParamDo
    r.Set("coupon_qrcode_param_do", _couponQrcodeParamDo)
    return nil
}

// CouponQrcodeParamDo Getter
func (r AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest) GetCouponQrcodeParamDo() *CouponQrcodeParamDO {
    return r._couponQrcodeParamDo
}
