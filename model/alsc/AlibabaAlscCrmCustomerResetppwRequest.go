package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
重置支付密码 API请求
alibaba.alsc.crm.customer.resetppw

重置支付密码
*/
type AlibabaAlscCrmCustomerResetppwRequest struct {
    model.Params
    // 系统自动生成
    resetPayPwdRequest   *ResetPayPasswdOpenReq
}

// 初始化AlibabaAlscCrmCustomerResetppwRequest对象
func NewAlibabaAlscCrmCustomerResetppwRequest() *AlibabaAlscCrmCustomerResetppwRequest{
    return &AlibabaAlscCrmCustomerResetppwRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCustomerResetppwRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.customer.resetppw"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCustomerResetppwRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ResetPayPwdRequest Setter
// 系统自动生成
func (r *AlibabaAlscCrmCustomerResetppwRequest) SetResetPayPwdRequest(resetPayPwdRequest *ResetPayPasswdOpenReq) error {
    r.resetPayPwdRequest = resetPayPwdRequest
    r.Set("reset_pay_pwd_request", resetPayPwdRequest)
    return nil
}

// ResetPayPwdRequest Getter
func (r AlibabaAlscCrmCustomerResetppwRequest) GetResetPayPwdRequest() *ResetPayPasswdOpenReq {
    return r.resetPayPwdRequest
}
