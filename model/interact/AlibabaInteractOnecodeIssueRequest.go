package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
onecode代码通用鉴权 API请求
alibaba.interact.onecode.issue

手淘开放鉴权接口，仅用于tida接口鉴权，无输入输出。
*/
type AlibabaInteractOnecodeIssueAPIRequest struct {
    model.Params
}

// 初始化AlibabaInteractOnecodeIssueAPIRequest对象
func NewAlibabaInteractOnecodeIssueRequest() *AlibabaInteractOnecodeIssueAPIRequest{
    return &AlibabaInteractOnecodeIssueAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractOnecodeIssueAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.onecode.issue"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractOnecodeIssueAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
