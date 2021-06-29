package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询顾客详情 API请求
alibaba.alsc.crm.customer.get

查询顾客详情
*/
type AlibabaAlscCrmCustomerGetRequest struct {
    model.Params
    // 顾客详情查询条件
    paramCustomerIdQueryOpenReq   *CustomerIdQueryOpenReq
}

// 初始化AlibabaAlscCrmCustomerGetRequest对象
func NewAlibabaAlscCrmCustomerGetRequest() *AlibabaAlscCrmCustomerGetRequest{
    return &AlibabaAlscCrmCustomerGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCustomerGetRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.customer.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCustomerGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamCustomerIdQueryOpenReq Setter
// 顾客详情查询条件
func (r *AlibabaAlscCrmCustomerGetRequest) SetParamCustomerIdQueryOpenReq(paramCustomerIdQueryOpenReq *CustomerIdQueryOpenReq) error {
    r.paramCustomerIdQueryOpenReq = paramCustomerIdQueryOpenReq
    r.Set("param_customer_id_query_open_req", paramCustomerIdQueryOpenReq)
    return nil
}

// ParamCustomerIdQueryOpenReq Getter
func (r AlibabaAlscCrmCustomerGetRequest) GetParamCustomerIdQueryOpenReq() *CustomerIdQueryOpenReq {
    return r.paramCustomerIdQueryOpenReq
}
