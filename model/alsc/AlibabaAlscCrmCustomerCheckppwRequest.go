package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
校验支付密码 API请求
alibaba.alsc.crm.customer.checkppw

校验支付密码
*/
type AlibabaAlscCrmCustomerCheckppwRequest struct {
    model.Params
    // 请求参数
    checkRequest   *CheckPayPasswdReq
}

// 初始化AlibabaAlscCrmCustomerCheckppwRequest对象
func NewAlibabaAlscCrmCustomerCheckppwRequest() *AlibabaAlscCrmCustomerCheckppwRequest{
    return &AlibabaAlscCrmCustomerCheckppwRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCustomerCheckppwRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.customer.checkppw"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCustomerCheckppwRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CheckRequest Setter
// 请求参数
func (r *AlibabaAlscCrmCustomerCheckppwRequest) SetCheckRequest(checkRequest *CheckPayPasswdReq) error {
    r.checkRequest = checkRequest
    r.Set("check_request", checkRequest)
    return nil
}

// CheckRequest Getter
func (r AlibabaAlscCrmCustomerCheckppwRequest) GetCheckRequest() *CheckPayPasswdReq {
    return r.checkRequest
}
