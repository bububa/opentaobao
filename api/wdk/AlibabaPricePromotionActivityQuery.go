package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
查询盒马帮档期活动详情 
alibaba.price.promotion.activity.query

查询盒马帮档期活动详情
*/
func AlibabaPricePromotionActivityQuery(clt *core.SDKClient, req *wdk.AlibabaPricePromotionActivityQueryRequest, session string) (*wdk.AlibabaPricePromotionActivityQueryResponse, error) {
    var resp wdk.AlibabaPricePromotionActivityQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
