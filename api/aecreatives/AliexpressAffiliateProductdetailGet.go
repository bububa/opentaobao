package aecreatives

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aecreatives"
)

/* 
联盟商品详情获取接口 
aliexpress.affiliate.productdetail.get

联盟推广商品搜索接口，用于搜索联盟推广商品数据
*/
func AliexpressAffiliateProductdetailGet(clt *core.SDKClient, req *aecreatives.AliexpressAffiliateProductdetailGetAPIRequest, session string) (*aecreatives.AliexpressAffiliateProductdetailGetAPIResponse, error) {
    var resp aecreatives.AliexpressAffiliateProductdetailGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
