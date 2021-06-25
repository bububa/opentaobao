package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询顾客详情 APIRequest
alibaba.alsc.crm.customer.get

查询顾客详情
*/
type AlibabaAlscCrmCustomerGetRequest struct {
    model.Params

    // 顾客详情查询条件
    paramCustomerIdQueryOpenReq   *CustomerIdQueryOpenReq 

}

func NewAlibabaAlscCrmCustomerGetRequest() *AlibabaAlscCrmCustomerGetRequest{
    return &AlibabaAlscCrmCustomerGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmCustomerGetRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.customer.get"
}

func (r AlibabaAlscCrmCustomerGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmCustomerGetRequest) SetParamCustomerIdQueryOpenReq(paramCustomerIdQueryOpenReq *CustomerIdQueryOpenReq) error {
    r.paramCustomerIdQueryOpenReq = paramCustomerIdQueryOpenReq
    r.Set("param_customer_id_query_open_req", paramCustomerIdQueryOpenReq)
    return nil
}

func (r AlibabaAlscCrmCustomerGetRequest) GetParamCustomerIdQueryOpenReq() *CustomerIdQueryOpenReq {
    return r.paramCustomerIdQueryOpenReq
}

