package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCustomerUpdateAPIRequest 更新顾客信息 API请求
// alibaba.alsc.crm.customer.update
//
// 更新顾客信息
type AlibabaAlscCrmCustomerUpdateAPIRequest struct {
	model.Params
	// 修改顾客参数
	_paramCustomerUpdateOpenReq *CustomerUpdateOpenReq
}

// NewAlibabaAlscCrmCustomerUpdateRequest 初始化AlibabaAlscCrmCustomerUpdateAPIRequest对象
func NewAlibabaAlscCrmCustomerUpdateRequest() *AlibabaAlscCrmCustomerUpdateAPIRequest {
	return &AlibabaAlscCrmCustomerUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCustomerUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.customer.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCustomerUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmCustomerUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCustomerUpdateOpenReq is ParamCustomerUpdateOpenReq Setter
// 修改顾客参数
func (r *AlibabaAlscCrmCustomerUpdateAPIRequest) SetParamCustomerUpdateOpenReq(_paramCustomerUpdateOpenReq *CustomerUpdateOpenReq) error {
	r._paramCustomerUpdateOpenReq = _paramCustomerUpdateOpenReq
	r.Set("param_customer_update_open_req", _paramCustomerUpdateOpenReq)
	return nil
}

// GetParamCustomerUpdateOpenReq ParamCustomerUpdateOpenReq Getter
func (r AlibabaAlscCrmCustomerUpdateAPIRequest) GetParamCustomerUpdateOpenReq() *CustomerUpdateOpenReq {
	return r._paramCustomerUpdateOpenReq
}
