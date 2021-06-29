package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询日消耗预算 API请求
alibaba.scbp.account.budget.get

查询日消耗预算
*/
type AlibabaScbpAccountBudgetGetRequest struct {
    model.Params
}

// 初始化AlibabaScbpAccountBudgetGetRequest对象
func NewAlibabaScbpAccountBudgetGetRequest() *AlibabaScbpAccountBudgetGetRequest{
    return &AlibabaScbpAccountBudgetGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAccountBudgetGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.account.budget.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAccountBudgetGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
