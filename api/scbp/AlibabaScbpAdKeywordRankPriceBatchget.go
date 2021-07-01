package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
外贸直通车关键词前五名批量排价 
alibaba.scbp.ad.keyword.rank.price.batchget

外贸直通车关键词前五名批量排价
*/
func AlibabaScbpAdKeywordRankPriceBatchget(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordRankPriceBatchgetAPIRequest, session string) (*scbp.AlibabaScbpAdKeywordRankPriceBatchgetAPIResponse, error) {
    var resp scbp.AlibabaScbpAdKeywordRankPriceBatchgetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
