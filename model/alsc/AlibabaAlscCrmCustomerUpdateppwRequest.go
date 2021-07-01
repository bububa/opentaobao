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
type AlibabaAlscCrmCustomerUpdateppwAPIRequest struct {
    model.Params
    // 修改密码
    _updatePayPasswdReq   *UpdatePayPasswdReq
}

// 初始化AlibabaAlscCrmCustomerUpdateppwAPIRequest对象
func NewAlibabaAlscCrmCustomerUpdateppwRequest() *AlibabaAlscCrmCustomerUpdateppwAPIRequest{
    return &AlibabaAlscCrmCustomerUpdateppwAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCustomerUpdateppwAPIRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.customer.updateppw"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCustomerUpdateppwAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UpdatePayPasswdReq Setter
// 修改密码
func (r *AlibabaAlscCrmCustomerUpdateppwAPIRequest) SetUpdatePayPasswdReq(_updatePayPasswdReq *UpdatePayPasswdReq) error {
    r._updatePayPasswdReq = _updatePayPasswdReq
    r.Set("update_pay_passwd_req", _updatePayPasswdReq)
    return nil
}

// UpdatePayPasswdReq Getter
func (r AlibabaAlscCrmCustomerUpdateppwAPIRequest) GetUpdatePayPasswdReq() *UpdatePayPasswdReq {
    return r._updatePayPasswdReq
}
