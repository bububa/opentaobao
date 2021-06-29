package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建顾客 API请求
alibaba.alsc.crm.customer.create

开放本地生活创建顾客功能
*/
type AlibabaAlscCrmCustomerCreateRequest struct {
    model.Params
    // 创建顾客参数
    paramCustomerCreateOpenReq   *CustomerCreateOpenReq
}

// 初始化AlibabaAlscCrmCustomerCreateRequest对象
func NewAlibabaAlscCrmCustomerCreateRequest() *AlibabaAlscCrmCustomerCreateRequest{
    return &AlibabaAlscCrmCustomerCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCustomerCreateRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.customer.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCustomerCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamCustomerCreateOpenReq Setter
// 创建顾客参数
func (r *AlibabaAlscCrmCustomerCreateRequest) SetParamCustomerCreateOpenReq(paramCustomerCreateOpenReq *CustomerCreateOpenReq) error {
    r.paramCustomerCreateOpenReq = paramCustomerCreateOpenReq
    r.Set("param_customer_create_open_req", paramCustomerCreateOpenReq)
    return nil
}

// ParamCustomerCreateOpenReq Getter
func (r AlibabaAlscCrmCustomerCreateRequest) GetParamCustomerCreateOpenReq() *CustomerCreateOpenReq {
    return r.paramCustomerCreateOpenReq
}
