package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
查询账户级别关键词推广状态 
alibaba.scbp.account.status.get

查询账户级别关键词推广状态
*/
func AlibabaScbpAccountStatusGet(clt *core.SDKClient, req *scbp.AlibabaScbpAccountStatusGetRequest, session string) (*scbp.AlibabaScbpAccountStatusGetResponse, error) {
    var resp scbp.AlibabaScbpAccountStatusGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
