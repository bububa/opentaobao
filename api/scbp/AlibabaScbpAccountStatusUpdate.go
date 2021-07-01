package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
修改账户级别关键词推广状态 
alibaba.scbp.account.status.update

修改账户级别关键词推广状态
*/
func AlibabaScbpAccountStatusUpdate(clt *core.SDKClient, req *scbp.AlibabaScbpAccountStatusUpdateAPIRequest, session string) (*scbp.AlibabaScbpAccountStatusUpdateAPIResponse, error) {
    var resp scbp.AlibabaScbpAccountStatusUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
