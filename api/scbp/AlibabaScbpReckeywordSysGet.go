package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
系统推荐 
alibaba.scbp.reckeyword.sys.get

查询系统推荐词
*/
func AlibabaScbpReckeywordSysGet(clt *core.SDKClient, req *scbp.AlibabaScbpReckeywordSysGetRequest, session string) (*scbp.AlibabaScbpReckeywordSysGetResponse, error) {
    var resp scbp.AlibabaScbpReckeywordSysGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
