package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询关键词推广账户是否欠款 APIRequest
alibaba.scbp.account.isarrears.get

查询关键词推广账户是否欠款
*/
type AlibabaScbpAccountIsarrearsGetRequest struct {
    model.Params

}

func NewAlibabaScbpAccountIsarrearsGetRequest() *AlibabaScbpAccountIsarrearsGetRequest{
    return &AlibabaScbpAccountIsarrearsGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAccountIsarrearsGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.account.isarrears.get"
}

func (r AlibabaScbpAccountIsarrearsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


