package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询推广账户等级 APIRequest
alibaba.scbp.ad.account.level.get

查询推广账户等级
*/
type AlibabaScbpAdAccountLevelGetRequest struct {
    model.Params

}

func NewAlibabaScbpAdAccountLevelGetRequest() *AlibabaScbpAdAccountLevelGetRequest{
    return &AlibabaScbpAdAccountLevelGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdAccountLevelGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.account.level.get"
}

func (r AlibabaScbpAdAccountLevelGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


