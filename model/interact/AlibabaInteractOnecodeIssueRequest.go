package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
onecode代码通用鉴权 APIRequest
alibaba.interact.onecode.issue

手淘开放鉴权接口，仅用于tida接口鉴权，无输入输出。
*/
type AlibabaInteractOnecodeIssueRequest struct {
    model.Params

}

func NewAlibabaInteractOnecodeIssueRequest() *AlibabaInteractOnecodeIssueRequest{
    return &AlibabaInteractOnecodeIssueRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractOnecodeIssueRequest) GetApiMethodName() string {
    return "alibaba.interact.onecode.issue"
}

func (r AlibabaInteractOnecodeIssueRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


