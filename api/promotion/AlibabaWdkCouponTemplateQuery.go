package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
优惠券模版查询 
alibaba.wdk.coupon.template.query

优惠券模版查询
*/
func AlibabaWdkCouponTemplateQuery(clt *core.SDKClient, req *promotion.AlibabaWdkCouponTemplateQueryRequest, session string) (*promotion.AlibabaWdkCouponTemplateQueryAPIResponse, error) {
    var resp promotion.AlibabaWdkCouponTemplateQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
