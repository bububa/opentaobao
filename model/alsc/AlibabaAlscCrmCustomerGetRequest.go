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
type AlibabaAlscCrmCustomerGetAPIRequest struct {
    model.Params
    // 顾客详情查询条件
    _paramCustomerIdQueryOpenReq   *CustomerIdQueryOpenReq
}

// 初始化AlibabaAlscCrmCustomerGetAPIRequest对象
func NewAlibabaAlscCrmCustomerGetRequest() *AlibabaAlscCrmCustomerGetAPIRequest{
    return &AlibabaAlscCrmCustomerGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCustomerGetAPIRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.customer.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCustomerGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamCustomerIdQueryOpenReq Setter
// 顾客详情查询条件
func (r *AlibabaAlscCrmCustomerGetAPIRequest) SetParamCustomerIdQueryOpenReq(_paramCustomerIdQueryOpenReq *CustomerIdQueryOpenReq) error {
    r._paramCustomerIdQueryOpenReq = _paramCustomerIdQueryOpenReq
    r.Set("param_customer_id_query_open_req", _paramCustomerIdQueryOpenReq)
    return nil
}

// ParamCustomerIdQueryOpenReq Getter
func (r AlibabaAlscCrmCustomerGetAPIRequest) GetParamCustomerIdQueryOpenReq() *CustomerIdQueryOpenReq {
    return r._paramCustomerIdQueryOpenReq
}
