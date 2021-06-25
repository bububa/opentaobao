package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询账户级别关键词推广状态 APIRequest
alibaba.scbp.account.status.get

查询账户级别关键词推广状态
*/
type AlibabaScbpAccountStatusGetRequest struct {
    model.Params

}

func NewAlibabaScbpAccountStatusGetRequest() *AlibabaScbpAccountStatusGetRequest{
    return &AlibabaScbpAccountStatusGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAccountStatusGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.account.status.get"
}

func (r AlibabaScbpAccountStatusGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


