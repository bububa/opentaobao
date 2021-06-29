package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改支付密码 API请求
alibaba.alsc.crm.customer.updateppw

修改支付密码
*/
type AlibabaAlscCrmCustomerUpdateppwRequest struct {
    model.Params
    // 修改密码
    updatePayPasswdReq   *UpdatePayPasswdReq
}

// 初始化AlibabaAlscCrmCustomerUpdateppwRequest对象
func NewAlibabaAlscCrmCustomerUpdateppwRequest() *AlibabaAlscCrmCustomerUpdateppwRequest{
    return &AlibabaAlscCrmCustomerUpdateppwRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCustomerUpdateppwRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.customer.updateppw"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCustomerUpdateppwRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UpdatePayPasswdReq Setter
// 修改密码
func (r *AlibabaAlscCrmCustomerUpdateppwRequest) SetUpdatePayPasswdReq(updatePayPasswdReq *UpdatePayPasswdReq) error {
    r.updatePayPasswdReq = updatePayPasswdReq
    r.Set("update_pay_passwd_req", updatePayPasswdReq)
    return nil
}

// UpdatePayPasswdReq Getter
func (r AlibabaAlscCrmCustomerUpdateppwRequest) GetUpdatePayPasswdReq() *UpdatePayPasswdReq {
    return r.updatePayPasswdReq
}
