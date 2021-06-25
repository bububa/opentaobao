package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卡号绑定顾客 APIRequest
alibaba.alsc.crm.card.bindcustomer

为卡号绑定顾客
*/
type AlibabaAlscCrmCardBindcustomerRequest struct {
    model.Params

    // 请求参数
    paramBindCustomerOpenReq   *BindCustomerOpenReq 

}

func NewAlibabaAlscCrmCardBindcustomerRequest() *AlibabaAlscCrmCardBindcustomerRequest{
    return &AlibabaAlscCrmCardBindcustomerRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmCardBindcustomerRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.card.bindcustomer"
}

func (r AlibabaAlscCrmCardBindcustomerRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmCardBindcustomerRequest) SetParamBindCustomerOpenReq(paramBindCustomerOpenReq *BindCustomerOpenReq) error {
    r.paramBindCustomerOpenReq = paramBindCustomerOpenReq
    r.Set("param_bind_customer_open_req", paramBindCustomerOpenReq)
    return nil
}

func (r AlibabaAlscCrmCardBindcustomerRequest) GetParamBindCustomerOpenReq() *BindCustomerOpenReq {
    return r.paramBindCustomerOpenReq
}

