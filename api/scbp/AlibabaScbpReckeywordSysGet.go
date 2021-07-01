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
func AlibabaScbpReckeywordSysGet(clt *core.SDKClient, req *scbp.AlibabaScbpReckeywordSysGetAPIRequest, session string) (*scbp.AlibabaScbpReckeywordSysGetAPIResponse, error) {
    var resp scbp.AlibabaScbpReckeywordSysGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
