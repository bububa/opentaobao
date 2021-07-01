package icbuseller

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbuseller"
)

/* 
优惠券校验 
alibaba.seller.coupon.auth.verify

优惠券校验
*/
func AlibabaSellerCouponAuthVerify(clt *core.SDKClient, req *icbuseller.AlibabaSellerCouponAuthVerifyAPIRequest, session string) (*icbuseller.AlibabaSellerCouponAuthVerifyAPIResponse, error) {
    var resp icbuseller.AlibabaSellerCouponAuthVerifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
