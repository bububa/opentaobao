package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
修改关键词价格 
alibaba.scbp.ad.keyword.update.keyword.price.batch

修改关键词价格
*/
func AlibabaScbpAdKeywordUpdateKeywordPriceBatch(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordUpdateKeywordPriceBatchRequest, session string) (*scbp.AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponse, error) {
    var resp scbp.AlibabaScbpAdKeywordUpdateKeywordPriceBatchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
