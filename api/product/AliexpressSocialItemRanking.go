package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
社交排行榜 
aliexpress.social.item.ranking

社交商品成交排行榜
*/
func AliexpressSocialItemRanking(clt *core.SDKClient, req *product.AliexpressSocialItemRankingAPIRequest, session string) (*product.AliexpressSocialItemRankingAPIResponse, error) {
    var resp product.AliexpressSocialItemRankingAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
