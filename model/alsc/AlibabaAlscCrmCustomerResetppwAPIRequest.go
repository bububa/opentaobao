package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCustomerResetppwAPIRequest
重置支付密码 API请求
alibaba.alsc.crm.customer.resetppw

重置支付密码 */
type AlibabaAlscCrmCustomerResetppwAPIRequest struct {
	model.Params
	// 系统自动生成
	_resetPayPwdRequest *ResetPayPasswdOpenReq
}

// NewAlibabaAlscCrmCustomerResetppwRequest 初始化AlibabaAlscCrmCustomerResetppwAPIRequest对象
func NewAlibabaAlscCrmCustomerResetppwRequest() *AlibabaAlscCrmCustomerResetppwAPIRequest {
	return &AlibabaAlscCrmCustomerResetppwAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCustomerResetppwAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.customer.resetppw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCustomerResetppwAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ResetPayPwdRequest Setter
// 系统自动生成
func (r *AlibabaAlscCrmCustomerResetppwAPIRequest) SetResetPayPwdRequest(_resetPayPwdRequest *ResetPayPasswdOpenReq) error {
	r._resetPayPwdRequest = _resetPayPwdRequest
	r.Set("reset_pay_pwd_request", _resetPayPwdRequest)
	return nil
}

// Get ResetPayPwdRequest Getter
func (r AlibabaAlscCrmCustomerResetppwAPIRequest) GetResetPayPwdRequest() *ResetPayPasswdOpenReq {
	return r._resetPayPwdRequest
}
