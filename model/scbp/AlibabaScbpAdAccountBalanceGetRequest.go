package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询账户余额 API请求
alibaba.scbp.ad.account.balance.get

查询推广账户余额
*/
type AlibabaScbpAdAccountBalanceGetRequest struct {
    model.Params
}

// 初始化AlibabaScbpAdAccountBalanceGetRequest对象
func NewAlibabaScbpAdAccountBalanceGetRequest() *AlibabaScbpAdAccountBalanceGetRequest{
    return &AlibabaScbpAdAccountBalanceGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdAccountBalanceGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.account.balance.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdAccountBalanceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
