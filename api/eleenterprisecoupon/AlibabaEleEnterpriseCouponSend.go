package eleenterprisecoupon

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eleenterprisecoupon"
)

/* 
发放优惠券 
alibaba.ele.enterprise.coupon.send

发放优惠券
*/
func AlibabaEleEnterpriseCouponSend(clt *core.SDKClient, req *eleenterprisecoupon.AlibabaEleEnterpriseCouponSendAPIRequest, session string) (*eleenterprisecoupon.AlibabaEleEnterpriseCouponSendAPIResponse, error) {
    var resp eleenterprisecoupon.AlibabaEleEnterpriseCouponSendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
