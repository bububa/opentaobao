package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
推荐词-词推词 
alibaba.scbp.reckeyword.search

推荐词-词推词
*/
func AlibabaScbpReckeywordSearch(clt *core.SDKClient, req *scbp.AlibabaScbpReckeywordSearchRequest, session string) (*scbp.AlibabaScbpReckeywordSearchResponse, error) {
    var resp scbp.AlibabaScbpReckeywordSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
