package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
onecode代码通用鉴权 
alibaba.interact.onecode.issue

手淘开放鉴权接口，仅用于tida接口鉴权，无输入输出。
*/
func AlibabaInteractOnecodeIssue(clt *core.SDKClient, req *interact.AlibabaInteractOnecodeIssueAPIRequest, session string) (*interact.AlibabaInteractOnecodeIssueAPIResponse, error) {
    var resp interact.AlibabaInteractOnecodeIssueAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
