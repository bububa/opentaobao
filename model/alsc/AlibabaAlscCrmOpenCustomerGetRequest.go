package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询会员资产 API请求
alibaba.alsc.crm.open.customer.get

查询会员身份，资产等
*/
type AlibabaAlscCrmOpenCustomerGetRequest struct {
    model.Params
    // 入参
    _paramCustomerGetOpenReq   *CustomerGetOpenReq
}

// 初始化AlibabaAlscCrmOpenCustomerGetRequest对象
func NewAlibabaAlscCrmOpenCustomerGetRequest() *AlibabaAlscCrmOpenCustomerGetRequest{
    return &AlibabaAlscCrmOpenCustomerGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmOpenCustomerGetRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.open.customer.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmOpenCustomerGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamCustomerGetOpenReq Setter
// 入参
func (r *AlibabaAlscCrmOpenCustomerGetRequest) SetParamCustomerGetOpenReq(_paramCustomerGetOpenReq *CustomerGetOpenReq) error {
    r._paramCustomerGetOpenReq = _paramCustomerGetOpenReq
    r.Set("param_customer_get_open_req", _paramCustomerGetOpenReq)
    return nil
}

// ParamCustomerGetOpenReq Getter
func (r AlibabaAlscCrmOpenCustomerGetRequest) GetParamCustomerGetOpenReq() *CustomerGetOpenReq {
    return r._paramCustomerGetOpenReq
}
