package aecreatives

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aecreatives"
)

/* 
查询联盟爆品数据 
aliexpress.affiliate.hotproduct.query

查询联盟爆品API
*/
func AliexpressAffiliateHotproductQuery(clt *core.SDKClient, req *aecreatives.AliexpressAffiliateHotproductQueryAPIRequest, session string) (*aecreatives.AliexpressAffiliateHotproductQueryAPIResponse, error) {
    var resp aecreatives.AliexpressAffiliateHotproductQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
