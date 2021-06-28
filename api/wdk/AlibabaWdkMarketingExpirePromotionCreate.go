package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
短保优惠创建 
alibaba.wdk.marketing.expire.promotion.create

过期优惠优惠信息录入
*/
func AlibabaWdkMarketingExpirePromotionCreate(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingExpirePromotionCreateRequest, session string) (*wdk.AlibabaWdkMarketingExpirePromotionCreateAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingExpirePromotionCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
