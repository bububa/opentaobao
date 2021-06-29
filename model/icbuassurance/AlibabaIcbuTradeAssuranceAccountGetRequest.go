package icbuassurance

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
icbu信保账户信息 APIRequest
alibaba.icbu.trade.assurance.account.get

icbu交易信用保障开通状态&额度信息查询
*/
type AlibabaIcbuTradeAssuranceAccountGetRequest struct {
    model.Params

}

func NewAlibabaIcbuTradeAssuranceAccountGetRequest() *AlibabaIcbuTradeAssuranceAccountGetRequest{
    return &AlibabaIcbuTradeAssuranceAccountGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuTradeAssuranceAccountGetRequest) GetApiMethodName() string {
    return "alibaba.icbu.trade.assurance.account.get"
}

func (r AlibabaIcbuTradeAssuranceAccountGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


