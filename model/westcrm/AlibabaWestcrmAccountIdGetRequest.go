package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据支付宝id查询IB的id APIRequest
alibaba.westcrm.account.id.get

根据支付宝id查询IB的id
*/
type AlibabaWestcrmAccountIdGetRequest struct {
    model.Params

    // 支付宝id
    alipayId   string 

}

func NewAlibabaWestcrmAccountIdGetRequest() *AlibabaWestcrmAccountIdGetRequest{
    return &AlibabaWestcrmAccountIdGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWestcrmAccountIdGetRequest) GetApiMethodName() string {
    return "alibaba.westcrm.account.id.get"
}

func (r AlibabaWestcrmAccountIdGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWestcrmAccountIdGetRequest) SetAlipayId(alipayId string) error {
    r.alipayId = alipayId
    r.Set("alipay_id", alipayId)
    return nil
}

func (r AlibabaWestcrmAccountIdGetRequest) GetAlipayId() string {
    return r.alipayId
}

