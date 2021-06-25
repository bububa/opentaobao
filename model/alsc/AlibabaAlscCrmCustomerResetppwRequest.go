package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
重置支付密码 APIRequest
alibaba.alsc.crm.customer.resetppw

重置支付密码
*/
type AlibabaAlscCrmCustomerResetppwRequest struct {
    model.Params

    // 系统自动生成
    resetPayPwdRequest   *ResetPayPasswdOpenReq 

}

func NewAlibabaAlscCrmCustomerResetppwRequest() *AlibabaAlscCrmCustomerResetppwRequest{
    return &AlibabaAlscCrmCustomerResetppwRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmCustomerResetppwRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.customer.resetppw"
}

func (r AlibabaAlscCrmCustomerResetppwRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmCustomerResetppwRequest) SetResetPayPwdRequest(resetPayPwdRequest *ResetPayPasswdOpenReq) error {
    r.resetPayPwdRequest = resetPayPwdRequest
    r.Set("reset_pay_pwd_request", resetPayPwdRequest)
    return nil
}

func (r AlibabaAlscCrmCustomerResetppwRequest) GetResetPayPwdRequest() *ResetPayPasswdOpenReq {
    return r.resetPayPwdRequest
}

