package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌营销导购员券页面二维码获取 APIRequest
alibaba.txcs.brandmarketing.coupon.qrcode.get

构建券页码二维码url
*/
type AlibabaTxcsBrandmarketingCouponQrcodeGetRequest struct {
    model.Params

    // 请求信息
    couponQrcodeParamDo   *CouponQrcodeParamDO 

}

func NewAlibabaTxcsBrandmarketingCouponQrcodeGetRequest() *AlibabaTxcsBrandmarketingCouponQrcodeGetRequest{
    return &AlibabaTxcsBrandmarketingCouponQrcodeGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTxcsBrandmarketingCouponQrcodeGetRequest) GetApiMethodName() string {
    return "alibaba.txcs.brandmarketing.coupon.qrcode.get"
}

func (r AlibabaTxcsBrandmarketingCouponQrcodeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTxcsBrandmarketingCouponQrcodeGetRequest) SetCouponQrcodeParamDo(couponQrcodeParamDo *CouponQrcodeParamDO) error {
    r.couponQrcodeParamDo = couponQrcodeParamDo
    r.Set("coupon_qrcode_param_do", couponQrcodeParamDo)
    return nil
}

func (r AlibabaTxcsBrandmarketingCouponQrcodeGetRequest) GetCouponQrcodeParamDo() *CouponQrcodeParamDO {
    return r.couponQrcodeParamDo
}

