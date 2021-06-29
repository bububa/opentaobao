package aecreatives

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aecreatives"
)

/* 
联盟物料智能推荐api 
aliexpress.affiliate.product.smartmatch

联盟物料算法智能推荐
*/
func AliexpressAffiliateProductSmartmatch(clt *core.SDKClient, req *aecreatives.AliexpressAffiliateProductSmartmatchRequest, session string) (*aecreatives.AliexpressAffiliateProductSmartmatchAPIResponse, error) {
    var resp aecreatives.AliexpressAffiliateProductSmartmatchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
