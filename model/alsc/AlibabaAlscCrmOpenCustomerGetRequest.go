package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询会员资产 APIRequest
alibaba.alsc.crm.open.customer.get

查询会员身份，资产等
*/
type AlibabaAlscCrmOpenCustomerGetRequest struct {
    model.Params

    // 入参
    paramCustomerGetOpenReq   *CustomerGetOpenReq 

}

func NewAlibabaAlscCrmOpenCustomerGetRequest() *AlibabaAlscCrmOpenCustomerGetRequest{
    return &AlibabaAlscCrmOpenCustomerGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmOpenCustomerGetRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.open.customer.get"
}

func (r AlibabaAlscCrmOpenCustomerGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmOpenCustomerGetRequest) SetParamCustomerGetOpenReq(paramCustomerGetOpenReq *CustomerGetOpenReq) error {
    r.paramCustomerGetOpenReq = paramCustomerGetOpenReq
    r.Set("param_customer_get_open_req", paramCustomerGetOpenReq)
    return nil
}

func (r AlibabaAlscCrmOpenCustomerGetRequest) GetParamCustomerGetOpenReq() *CustomerGetOpenReq {
    return r.paramCustomerGetOpenReq
}

