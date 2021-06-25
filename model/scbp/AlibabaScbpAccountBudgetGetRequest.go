package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询日消耗预算 APIRequest
alibaba.scbp.account.budget.get

查询日消耗预算
*/
type AlibabaScbpAccountBudgetGetRequest struct {
    model.Params

}

func NewAlibabaScbpAccountBudgetGetRequest() *AlibabaScbpAccountBudgetGetRequest{
    return &AlibabaScbpAccountBudgetGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAccountBudgetGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.account.budget.get"
}

func (r AlibabaScbpAccountBudgetGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


