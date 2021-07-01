package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
保存和更新顾客 API请求
alibaba.alsc.crm.open.customer.save

用来保存顾客，如果已经存在的话，则更新顾客
*/
type AlibabaAlscCrmOpenCustomerSaveAPIRequest struct {
    model.Params
    // 入参
    _paramCustomerSaveOpenReq   *CustomerSaveOpenReq
}

// 初始化AlibabaAlscCrmOpenCustomerSaveAPIRequest对象
func NewAlibabaAlscCrmOpenCustomerSaveRequest() *AlibabaAlscCrmOpenCustomerSaveAPIRequest{
    return &AlibabaAlscCrmOpenCustomerSaveAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmOpenCustomerSaveAPIRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.open.customer.save"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmOpenCustomerSaveAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamCustomerSaveOpenReq Setter
// 入参
func (r *AlibabaAlscCrmOpenCustomerSaveAPIRequest) SetParamCustomerSaveOpenReq(_paramCustomerSaveOpenReq *CustomerSaveOpenReq) error {
    r._paramCustomerSaveOpenReq = _paramCustomerSaveOpenReq
    r.Set("param_customer_save_open_req", _paramCustomerSaveOpenReq)
    return nil
}

// ParamCustomerSaveOpenReq Getter
func (r AlibabaAlscCrmOpenCustomerSaveAPIRequest) GetParamCustomerSaveOpenReq() *CustomerSaveOpenReq {
    return r._paramCustomerSaveOpenReq
}
