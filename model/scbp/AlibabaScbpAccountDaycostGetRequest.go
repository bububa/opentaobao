package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询今日消耗 APIRequest
alibaba.scbp.account.daycost.get

查询今日消耗
*/
type AlibabaScbpAccountDaycostGetRequest struct {
    model.Params

}

func NewAlibabaScbpAccountDaycostGetRequest() *AlibabaScbpAccountDaycostGetRequest{
    return &AlibabaScbpAccountDaycostGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAccountDaycostGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.account.daycost.get"
}

func (r AlibabaScbpAccountDaycostGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


