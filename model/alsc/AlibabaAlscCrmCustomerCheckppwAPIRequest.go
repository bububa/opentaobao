package alsc

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCustomerCheckppwAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.customer.checkppw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCustomerCheckppwAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CheckRequest Setter
// 请求参数
func (r *AlibabaAlscCrmCustomerCheckppwAPIRequest) SetCheckRequest(_checkRequest *CheckPayPasswdReq) error {
	r._checkRequest = _checkRequest
	r.Set("check_request", _checkRequest)
	return nil
}

// Get CheckRequest Getter
func (r AlibabaAlscCrmCustomerCheckppwAPIRequest) GetCheckRequest() *CheckPayPasswdReq {
	return r._checkRequest
}
