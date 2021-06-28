package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
外贸直通车关键词前五名排价 
alibaba.scbp.ad.keyword.rank.price.get

外贸直通车关键词前五名排价
*/
func AlibabaScbpAdKeywordRankPriceGet(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordRankPriceGetRequest, session string) (*scbp.AlibabaScbpAdKeywordRankPriceGetAPIResponse, error) {
    var resp scbp.AlibabaScbpAdKeywordRankPriceGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
