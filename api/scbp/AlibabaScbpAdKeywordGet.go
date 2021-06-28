package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
外贸直通车查询关键词 
alibaba.scbp.ad.keyword.get

外贸直通车查询关键词
*/
func AlibabaScbpAdKeywordGet(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordGetRequest, session string) (*scbp.AlibabaScbpAdKeywordGetAPIResponse, error) {
    var resp scbp.AlibabaScbpAdKeywordGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
