package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCustomerCreateAPIRequest
创建顾客 API请求
alibaba.alsc.crm.customer.create

开放本地生活创建顾客功能 */
type AlibabaAlscCrmCustomerCreateAPIRequest struct {
	model.Params
	// 创建顾客参数
	_paramCustomerCreateOpenReq *CustomerCreateOpenReq
}

// NewAlibabaAlscCrmCustomerCreateRequest 初始化AlibabaAlscCrmCustomerCreateAPIRequest对象
func NewAlibabaAlscCrmCustomerCreateRequest() *AlibabaAlscCrmCustomerCreateAPIRequest {
	return &AlibabaAlscCrmCustomerCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCustomerCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.customer.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCustomerCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamCustomerCreateOpenReq Setter
// 创建顾客参数
func (r *AlibabaAlscCrmCustomerCreateAPIRequest) SetParamCustomerCreateOpenReq(_paramCustomerCreateOpenReq *CustomerCreateOpenReq) error {
	r._paramCustomerCreateOpenReq = _paramCustomerCreateOpenReq
	r.Set("param_customer_create_open_req", _paramCustomerCreateOpenReq)
	return nil
}

// Get ParamCustomerCreateOpenReq Getter
func (r AlibabaAlscCrmCustomerCreateAPIRequest) GetParamCustomerCreateOpenReq() *CustomerCreateOpenReq {
	return r._paramCustomerCreateOpenReq
}
