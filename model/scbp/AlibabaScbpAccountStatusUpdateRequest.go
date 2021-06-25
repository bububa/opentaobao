package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改账户级别关键词推广状态 APIRequest
alibaba.scbp.account.status.update

修改账户级别关键词推广状态
*/
type AlibabaScbpAccountStatusUpdateRequest struct {
    model.Params

    // on:开启,off:暂停
    status   string 

}

func NewAlibabaScbpAccountStatusUpdateRequest() *AlibabaScbpAccountStatusUpdateRequest{
    return &AlibabaScbpAccountStatusUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAccountStatusUpdateRequest) GetApiMethodName() string {
    return "alibaba.scbp.account.status.update"
}

func (r AlibabaScbpAccountStatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAccountStatusUpdateRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaScbpAccountStatusUpdateRequest) GetStatus() string {
    return r.status
}

