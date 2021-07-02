package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmOpenCustomerGetAPIRequest 查询会员资产 API请求
// alibaba.alsc.crm.open.customer.get
//
// 查询会员身份，资产等
type AlibabaAlscCrmOpenCustomerGetAPIRequest struct {
	model.Params
	// 入参
	_paramCustomerGetOpenReq *CustomerGetOpenReq
}

// NewAlibabaAlscCrmOpenCustomerGetRequest 初始化AlibabaAlscCrmOpenCustomerGetAPIRequest对象
func NewAlibabaAlscCrmOpenCustomerGetRequest() *AlibabaAlscCrmOpenCustomerGetAPIRequest {
	return &AlibabaAlscCrmOpenCustomerGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmOpenCustomerGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.open.customer.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmOpenCustomerGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamCustomerGetOpenReq Setter
// 入参
func (r *AlibabaAlscCrmOpenCustomerGetAPIRequest) SetParamCustomerGetOpenReq(_paramCustomerGetOpenReq *CustomerGetOpenReq) error {
	r._paramCustomerGetOpenReq = _paramCustomerGetOpenReq
	r.Set("param_customer_get_open_req", _paramCustomerGetOpenReq)
	return nil
}

// Get ParamCustomerGetOpenReq Getter
func (r AlibabaAlscCrmOpenCustomerGetAPIRequest) GetParamCustomerGetOpenReq() *CustomerGetOpenReq {
	return r._paramCustomerGetOpenReq
}
