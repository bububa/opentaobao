package shop

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/shop"
)

/* 
获取优惠券信息 
alibaba.data.coupon.get

获取优惠券信息，仅作客户端鉴权虚拟api使用
*/
func AlibabaDataCouponGet(clt *core.SDKClient, req *shop.AlibabaDataCouponGetRequest, session string) (*shop.AlibabaDataCouponGetAPIResponse, error) {
    var resp shop.AlibabaDataCouponGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
