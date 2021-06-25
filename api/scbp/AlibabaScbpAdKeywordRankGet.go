package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
获取外贸直通车关键词预估排名 
alibaba.scbp.ad.keyword.rank.get

获取外贸直通车关键词预估排名
*/
func AlibabaScbpAdKeywordRankGet(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordRankGetRequest, session string) (*scbp.AlibabaScbpAdKeywordRankGetResponse, error) {
    var resp scbp.AlibabaScbpAdKeywordRankGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
