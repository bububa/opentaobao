package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改账户级别关键词推广状态 API请求
alibaba.scbp.account.status.update

修改账户级别关键词推广状态
*/
type AlibabaScbpAccountStatusUpdateRequest struct {
    model.Params
    // on:开启,off:暂停
    status   string
}

// 初始化AlibabaScbpAccountStatusUpdateRequest对象
func NewAlibabaScbpAccountStatusUpdateRequest() *AlibabaScbpAccountStatusUpdateRequest{
    return &AlibabaScbpAccountStatusUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAccountStatusUpdateRequest) GetApiMethodName() string {
    return "alibaba.scbp.account.status.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAccountStatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Status Setter
// on:开启,off:暂停
func (r *AlibabaScbpAccountStatusUpdateRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r AlibabaScbpAccountStatusUpdateRequest) GetStatus() string {
    return r.status
}
