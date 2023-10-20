package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCustomerCheckppwAPIRequest 校验支付密码 API请求
// alibaba.alsc.crm.customer.checkppw
//
// 校验支付密码
type AlibabaAlscCrmCustomerCheckppwAPIRequest struct {
	model.Params
	// 请求参数
	_checkRequest *CheckPayPasswdReq
}

// NewAlibabaAlscCrmCustomerCheckppwRequest 初始化AlibabaAlscCrmCustomerCheckppwAPIRequest对象
func NewAlibabaAlscCrmCustomerCheckppwRequest() *AlibabaAlscCrmCustomerCheckppwAPIRequest {
	return &AlibabaAlscCrmCustomerCheckppwAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmCustomerCheckppwAPIRequest) Reset() {
	r._checkRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCustomerCheckppwAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.customer.checkppw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCustomerCheckppwAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmCustomerCheckppwAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCheckRequest is CheckRequest Setter
// 请求参数
func (r *AlibabaAlscCrmCustomerCheckppwAPIRequest) SetCheckRequest(_checkRequest *CheckPayPasswdReq) error {
	r._checkRequest = _checkRequest
	r.Set("check_request", _checkRequest)
	return nil
}

// GetCheckRequest CheckRequest Getter
func (r AlibabaAlscCrmCustomerCheckppwAPIRequest) GetCheckRequest() *CheckPayPasswdReq {
	return r._checkRequest
}

var poolAlibabaAlscCrmCustomerCheckppwAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmCustomerCheckppwRequest()
	},
}

// GetAlibabaAlscCrmCustomerCheckppwRequest 从 sync.Pool 获取 AlibabaAlscCrmCustomerCheckppwAPIRequest
func GetAlibabaAlscCrmCustomerCheckppwAPIRequest() *AlibabaAlscCrmCustomerCheckppwAPIRequest {
	return poolAlibabaAlscCrmCustomerCheckppwAPIRequest.Get().(*AlibabaAlscCrmCustomerCheckppwAPIRequest)
}

// ReleaseAlibabaAlscCrmCustomerCheckppwAPIRequest 将 AlibabaAlscCrmCustomerCheckppwAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmCustomerCheckppwAPIRequest(v *AlibabaAlscCrmCustomerCheckppwAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmCustomerCheckppwAPIRequest.Put(v)
}
