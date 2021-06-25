package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
校验支付密码 APIRequest
alibaba.alsc.crm.customer.checkppw

校验支付密码
*/
type AlibabaAlscCrmCustomerCheckppwRequest struct {
    model.Params

    // 请求参数
    checkRequest   *CheckPayPasswdReq 

}

func NewAlibabaAlscCrmCustomerCheckppwRequest() *AlibabaAlscCrmCustomerCheckppwRequest{
    return &AlibabaAlscCrmCustomerCheckppwRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmCustomerCheckppwRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.customer.checkppw"
}

func (r AlibabaAlscCrmCustomerCheckppwRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmCustomerCheckppwRequest) SetCheckRequest(checkRequest *CheckPayPasswdReq) error {
    r.checkRequest = checkRequest
    r.Set("check_request", checkRequest)
    return nil
}

func (r AlibabaAlscCrmCustomerCheckppwRequest) GetCheckRequest() *CheckPayPasswdReq {
    return r.checkRequest
}

