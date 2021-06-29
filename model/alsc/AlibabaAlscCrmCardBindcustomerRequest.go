package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卡号绑定顾客 API请求
alibaba.alsc.crm.card.bindcustomer

为卡号绑定顾客
*/
type AlibabaAlscCrmCardBindcustomerRequest struct {
    model.Params
    // 请求参数
    _paramBindCustomerOpenReq   *BindCustomerOpenReq
}

// 初始化AlibabaAlscCrmCardBindcustomerRequest对象
func NewAlibabaAlscCrmCardBindcustomerRequest() *AlibabaAlscCrmCardBindcustomerRequest{
    return &AlibabaAlscCrmCardBindcustomerRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardBindcustomerRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.card.bindcustomer"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardBindcustomerRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamBindCustomerOpenReq Setter
// 请求参数
func (r *AlibabaAlscCrmCardBindcustomerRequest) SetParamBindCustomerOpenReq(_paramBindCustomerOpenReq *BindCustomerOpenReq) error {
    r._paramBindCustomerOpenReq = _paramBindCustomerOpenReq
    r.Set("param_bind_customer_open_req", _paramBindCustomerOpenReq)
    return nil
}

// ParamBindCustomerOpenReq Getter
func (r AlibabaAlscCrmCardBindcustomerRequest) GetParamBindCustomerOpenReq() *BindCustomerOpenReq {
    return r._paramBindCustomerOpenReq
}
