package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
优惠券模版创建 
alibaba.wdk.coupon.template.create

开放给外部商家创建优惠券模版
*/
func AlibabaWdkCouponTemplateCreate(clt *core.SDKClient, req *promotion.AlibabaWdkCouponTemplateCreateRequest, session string) (*promotion.AlibabaWdkCouponTemplateCreateAPIResponse, error) {
    var resp promotion.AlibabaWdkCouponTemplateCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
