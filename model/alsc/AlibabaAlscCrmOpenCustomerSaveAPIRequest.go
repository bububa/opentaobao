package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmOpenCustomerSaveAPIRequest 保存和更新顾客 API请求
// alibaba.alsc.crm.open.customer.save
//
// 用来保存顾客，如果已经存在的话，则更新顾客
type AlibabaAlscCrmOpenCustomerSaveAPIRequest struct {
	model.Params
	// 入参
	_paramCustomerSaveOpenReq *CustomerSaveOpenReq
}

// NewAlibabaAlscCrmOpenCustomerSaveRequest 初始化AlibabaAlscCrmOpenCustomerSaveAPIRequest对象
func NewAlibabaAlscCrmOpenCustomerSaveRequest() *AlibabaAlscCrmOpenCustomerSaveAPIRequest {
	return &AlibabaAlscCrmOpenCustomerSaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmOpenCustomerSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.open.customer.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmOpenCustomerSaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmOpenCustomerSaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCustomerSaveOpenReq is ParamCustomerSaveOpenReq Setter
// 入参
func (r *AlibabaAlscCrmOpenCustomerSaveAPIRequest) SetParamCustomerSaveOpenReq(_paramCustomerSaveOpenReq *CustomerSaveOpenReq) error {
	r._paramCustomerSaveOpenReq = _paramCustomerSaveOpenReq
	r.Set("param_customer_save_open_req", _paramCustomerSaveOpenReq)
	return nil
}

// GetParamCustomerSaveOpenReq ParamCustomerSaveOpenReq Getter
func (r AlibabaAlscCrmOpenCustomerSaveAPIRequest) GetParamCustomerSaveOpenReq() *CustomerSaveOpenReq {
	return r._paramCustomerSaveOpenReq
}
