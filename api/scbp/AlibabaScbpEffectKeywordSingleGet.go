package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
单个关键词效果报表 
alibaba.scbp.effect.keyword.single.get

单个关键词效果报表
*/
func AlibabaScbpEffectKeywordSingleGet(clt *core.SDKClient, req *scbp.AlibabaScbpEffectKeywordSingleGetRequest, session string) (*scbp.AlibabaScbpEffectKeywordSingleGetResponse, error) {
    var resp scbp.AlibabaScbpEffectKeywordSingleGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
