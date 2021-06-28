package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
查询关键词前五名排价 
alibaba.scbp.ad.keyword.query.keyword.rank.price

查询关键词前五名排价
*/
func AlibabaScbpAdKeywordQueryKeywordRankPrice(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordQueryKeywordRankPriceRequest, session string) (*scbp.AlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponse, error) {
    var resp scbp.AlibabaScbpAdKeywordQueryKeywordRankPriceAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
