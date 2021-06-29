package alitripmerchant

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alitripmerchant"
)

/* 
星河-offer查询 
alitrip.merchant.galaxy.offer.query

根据offer的ID查询offer信息
*/
func AlitripMerchantGalaxyOfferQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyOfferQueryRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyOfferQueryAPIResponse, error) {
    var resp alitripmerchant.AlitripMerchantGalaxyOfferQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
