package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询账户级别关键词推广状态 API请求
alibaba.scbp.account.status.get

查询账户级别关键词推广状态
*/
type AlibabaScbpAccountStatusGetRequest struct {
    model.Params
}

// 初始化AlibabaScbpAccountStatusGetRequest对象
func NewAlibabaScbpAccountStatusGetRequest() *AlibabaScbpAccountStatusGetRequest{
    return &AlibabaScbpAccountStatusGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAccountStatusGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.account.status.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAccountStatusGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
