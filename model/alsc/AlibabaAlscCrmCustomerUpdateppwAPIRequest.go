package alsc

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmCustomerUpdateppwAPIRequest) Reset() {
	r._updatePayPasswdReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCustomerUpdateppwAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.customer.updateppw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCustomerUpdateppwAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmCustomerUpdateppwAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdatePayPasswdReq is UpdatePayPasswdReq Setter
// 修改密码
func (r *AlibabaAlscCrmCustomerUpdateppwAPIRequest) SetUpdatePayPasswdReq(_updatePayPasswdReq *UpdatePayPasswdReq) error {
	r._updatePayPasswdReq = _updatePayPasswdReq
	r.Set("update_pay_passwd_req", _updatePayPasswdReq)
	return nil
}

// GetUpdatePayPasswdReq UpdatePayPasswdReq Getter
func (r AlibabaAlscCrmCustomerUpdateppwAPIRequest) GetUpdatePayPasswdReq() *UpdatePayPasswdReq {
	return r._updatePayPasswdReq
}

var poolAlibabaAlscCrmCustomerUpdateppwAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmCustomerUpdateppwRequest()
	},
}

// GetAlibabaAlscCrmCustomerUpdateppwRequest 从 sync.Pool 获取 AlibabaAlscCrmCustomerUpdateppwAPIRequest
func GetAlibabaAlscCrmCustomerUpdateppwAPIRequest() *AlibabaAlscCrmCustomerUpdateppwAPIRequest {
	return poolAlibabaAlscCrmCustomerUpdateppwAPIRequest.Get().(*AlibabaAlscCrmCustomerUpdateppwAPIRequest)
}

// ReleaseAlibabaAlscCrmCustomerUpdateppwAPIRequest 将 AlibabaAlscCrmCustomerUpdateppwAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmCustomerUpdateppwAPIRequest(v *AlibabaAlscCrmCustomerUpdateppwAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmCustomerUpdateppwAPIRequest.Put(v)
}
