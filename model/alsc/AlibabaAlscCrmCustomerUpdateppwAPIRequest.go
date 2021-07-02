package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCustomerUpdateppwAPIRequest 修改支付密码 API请求
// alibaba.alsc.crm.customer.updateppw
//
// 修改支付密码
type AlibabaAlscCrmCustomerUpdateppwAPIRequest struct {
	model.Params
	// 修改密码
	_updatePayPasswdReq *UpdatePayPasswdReq
}

// NewAlibabaAlscCrmCustomerUpdateppwRequest 初始化AlibabaAlscCrmCustomerUpdateppwAPIRequest对象
func NewAlibabaAlscCrmCustomerUpdateppwRequest() *AlibabaAlscCrmCustomerUpdateppwAPIRequest {
	return &AlibabaAlscCrmCustomerUpdateppwAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCustomerUpdateppwAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.customer.updateppw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCustomerUpdateppwAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is UpdatePayPasswdReq Setter
// 修改密码
func (r *AlibabaAlscCrmCustomerUpdateppwAPIRequest) SetUpdatePayPasswdReq(_updatePayPasswdReq *UpdatePayPasswdReq) error {
	r._updatePayPasswdReq = _updatePayPasswdReq
	r.Set("update_pay_passwd_req", _updatePayPasswdReq)
	return nil
}

// Get UpdatePayPasswdReq Getter
func (r AlibabaAlscCrmCustomerUpdateppwAPIRequest) GetUpdatePayPasswdReq() *UpdatePayPasswdReq {
	return r._updatePayPasswdReq
}
