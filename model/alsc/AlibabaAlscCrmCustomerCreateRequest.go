package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建顾客 APIRequest
alibaba.alsc.crm.customer.create

开放本地生活创建顾客功能
*/
type AlibabaAlscCrmCustomerCreateRequest struct {
    model.Params

    // 创建顾客参数
    paramCustomerCreateOpenReq   *CustomerCreateOpenReq 

}

func NewAlibabaAlscCrmCustomerCreateRequest() *AlibabaAlscCrmCustomerCreateRequest{
    return &AlibabaAlscCrmCustomerCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmCustomerCreateRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.customer.create"
}

func (r AlibabaAlscCrmCustomerCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmCustomerCreateRequest) SetParamCustomerCreateOpenReq(paramCustomerCreateOpenReq *CustomerCreateOpenReq) error {
    r.paramCustomerCreateOpenReq = paramCustomerCreateOpenReq
    r.Set("param_customer_create_open_req", paramCustomerCreateOpenReq)
    return nil
}

func (r AlibabaAlscCrmCustomerCreateRequest) GetParamCustomerCreateOpenReq() *CustomerCreateOpenReq {
    return r.paramCustomerCreateOpenReq
}

