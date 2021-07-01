package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
短保优惠删除 
alibaba.wdk.marketing.expire.promotion.delete

短保优惠删除
*/
func AlibabaWdkMarketingExpirePromotionDelete(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingExpirePromotionDeleteAPIRequest, session string) (*wdk.AlibabaWdkMarketingExpirePromotionDeleteAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingExpirePromotionDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
