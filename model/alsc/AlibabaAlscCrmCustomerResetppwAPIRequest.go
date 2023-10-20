package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCustomerResetppwAPIRequest 重置支付密码 API请求
// alibaba.alsc.crm.customer.resetppw
//
// 重置支付密码
type AlibabaAlscCrmCustomerResetppwAPIRequest struct {
	model.Params
	// 系统自动生成
	_resetPayPwdRequest *ResetPayPasswdOpenReq
}

// NewAlibabaAlscCrmCustomerResetppwRequest 初始化AlibabaAlscCrmCustomerResetppwAPIRequest对象
func NewAlibabaAlscCrmCustomerResetppwRequest() *AlibabaAlscCrmCustomerResetppwAPIRequest {
	return &AlibabaAlscCrmCustomerResetppwAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmCustomerResetppwAPIRequest) Reset() {
	r._resetPayPwdRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCustomerResetppwAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.customer.resetppw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCustomerResetppwAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmCustomerResetppwAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetResetPayPwdRequest is ResetPayPwdRequest Setter
// 系统自动生成
func (r *AlibabaAlscCrmCustomerResetppwAPIRequest) SetResetPayPwdRequest(_resetPayPwdRequest *ResetPayPasswdOpenReq) error {
	r._resetPayPwdRequest = _resetPayPwdRequest
	r.Set("reset_pay_pwd_request", _resetPayPwdRequest)
	return nil
}

// GetResetPayPwdRequest ResetPayPwdRequest Getter
func (r AlibabaAlscCrmCustomerResetppwAPIRequest) GetResetPayPwdRequest() *ResetPayPasswdOpenReq {
	return r._resetPayPwdRequest
}

var poolAlibabaAlscCrmCustomerResetppwAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmCustomerResetppwRequest()
	},
}

// GetAlibabaAlscCrmCustomerResetppwRequest 从 sync.Pool 获取 AlibabaAlscCrmCustomerResetppwAPIRequest
func GetAlibabaAlscCrmCustomerResetppwAPIRequest() *AlibabaAlscCrmCustomerResetppwAPIRequest {
	return poolAlibabaAlscCrmCustomerResetppwAPIRequest.Get().(*AlibabaAlscCrmCustomerResetppwAPIRequest)
}

// ReleaseAlibabaAlscCrmCustomerResetppwAPIRequest 将 AlibabaAlscCrmCustomerResetppwAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmCustomerResetppwAPIRequest(v *AlibabaAlscCrmCustomerResetppwAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmCustomerResetppwAPIRequest.Put(v)
}
