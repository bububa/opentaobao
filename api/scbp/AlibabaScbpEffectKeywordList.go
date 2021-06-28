package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
关键词报表 
alibaba.scbp.effect.keyword.list

关键词报表
*/
func AlibabaScbpEffectKeywordList(clt *core.SDKClient, req *scbp.AlibabaScbpEffectKeywordListRequest, session string) (*scbp.AlibabaScbpEffectKeywordListAPIResponse, error) {
    var resp scbp.AlibabaScbpEffectKeywordListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
