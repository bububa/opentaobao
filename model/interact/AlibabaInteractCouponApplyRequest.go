package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优惠券领取鉴权接口 APIRequest
alibaba.interact.coupon.apply

鉴权接口，为coupon.apply接口鉴权
*/
type AlibabaInteractCouponApplyRequest struct {
    model.Params

}

func NewAlibabaInteractCouponApplyRequest() *AlibabaInteractCouponApplyRequest{
    return &AlibabaInteractCouponApplyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractCouponApplyRequest) GetApiMethodName() string {
    return "alibaba.interact.coupon.apply"
}

func (r AlibabaInteractCouponApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


