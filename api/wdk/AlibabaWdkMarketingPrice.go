package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
促销价签服务 
alibaba.wdk.marketing.price

获取营销-促销商品中的实时价格
*/
func AlibabaWdkMarketingPrice(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingPriceRequest, session string) (*wdk.AlibabaWdkMarketingPriceAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingPriceAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
