package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新顾客信息 APIRequest
alibaba.alsc.crm.customer.update

更新顾客信息
*/
type AlibabaAlscCrmCustomerUpdateRequest struct {
    model.Params

    // 修改顾客参数
    paramCustomerUpdateOpenReq   *CustomerUpdateOpenReq 

}

func NewAlibabaAlscCrmCustomerUpdateRequest() *AlibabaAlscCrmCustomerUpdateRequest{
    return &AlibabaAlscCrmCustomerUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmCustomerUpdateRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.customer.update"
}

func (r AlibabaAlscCrmCustomerUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmCustomerUpdateRequest) SetParamCustomerUpdateOpenReq(paramCustomerUpdateOpenReq *CustomerUpdateOpenReq) error {
    r.paramCustomerUpdateOpenReq = paramCustomerUpdateOpenReq
    r.Set("param_customer_update_open_req", paramCustomerUpdateOpenReq)
    return nil
}

func (r AlibabaAlscCrmCustomerUpdateRequest) GetParamCustomerUpdateOpenReq() *CustomerUpdateOpenReq {
    return r.paramCustomerUpdateOpenReq
}

