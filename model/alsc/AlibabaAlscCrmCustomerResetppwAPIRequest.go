package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmcustomerresetppwAPIRequest 重置支付密码 API请求
// alibaba.alsc.crm.customer.resetppw
//
// 重置支付密码
type AlibabaalsccrmcustomerresetppwAPIRequest struct {
	model.Params
	// 系统自动生成
	_resetPayPwdRequest *ResetPayPasswdOpenReq
}

// NewAlibabaalsccrmcustomerresetppwRequest 初始化AlibabaalsccrmcustomerresetppwAPIRequest对象
func NewAlibabaalsccrmcustomerresetppwRequest() *AlibabaalsccrmcustomerresetppwAPIRequest {
	return &AlibabaalsccrmcustomerresetppwAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmcustomerresetppwAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.customer.resetppw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmcustomerresetppwAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmcustomerresetppwAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetResetPayPwdRequest is ResetPayPwdRequest Setter
// 系统自动生成
func (r *AlibabaalsccrmcustomerresetppwAPIRequest) SetResetPayPwdRequest(_resetPayPwdRequest *ResetPayPasswdOpenReq) error {
	r._resetPayPwdRequest = _resetPayPwdRequest
	r.Set("reset_pay_pwd_request", _resetPayPwdRequest)
	return nil
}

// GetResetPayPwdRequest ResetPayPwdRequest Getter
func (r AlibabaalsccrmcustomerresetppwAPIRequest) GetResetPayPwdRequest() *ResetPayPasswdOpenReq {
	return r._resetPayPwdRequest
}
