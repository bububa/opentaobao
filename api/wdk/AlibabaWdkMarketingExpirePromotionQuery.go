package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
短保优惠查询 
alibaba.wdk.marketing.expire.promotion.query

短保优惠查询
*/
func AlibabaWdkMarketingExpirePromotionQuery(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingExpirePromotionQueryRequest, session string) (*wdk.AlibabaWdkMarketingExpirePromotionQueryResponse, error) {
    var resp wdk.AlibabaWdkMarketingExpirePromotionQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
