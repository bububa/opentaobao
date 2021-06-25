package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询账户余额 APIRequest
alibaba.scbp.ad.account.balance.get

查询推广账户余额
*/
type AlibabaScbpAdAccountBalanceGetRequest struct {
    model.Params

}

func NewAlibabaScbpAdAccountBalanceGetRequest() *AlibabaScbpAdAccountBalanceGetRequest{
    return &AlibabaScbpAdAccountBalanceGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdAccountBalanceGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.account.balance.get"
}

func (r AlibabaScbpAdAccountBalanceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


