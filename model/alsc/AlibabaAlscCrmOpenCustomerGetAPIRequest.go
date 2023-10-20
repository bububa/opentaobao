package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmopencustomergetAPIRequest 查询会员资产 API请求
// alibaba.alsc.crm.open.customer.get
//
// 查询会员身份，资产等
type AlibabaalsccrmopencustomergetAPIRequest struct {
	model.Params
	// 入参
	_paramCustomerGetOpenReq *CustomerGetOpenReq
}

// NewAlibabaalsccrmopencustomergetRequest 初始化AlibabaalsccrmopencustomergetAPIRequest对象
func NewAlibabaalsccrmopencustomergetRequest() *AlibabaalsccrmopencustomergetAPIRequest {
	return &AlibabaalsccrmopencustomergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmopencustomergetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.open.customer.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmopencustomergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmopencustomergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCustomerGetOpenReq is ParamCustomerGetOpenReq Setter
// 入参
func (r *AlibabaalsccrmopencustomergetAPIRequest) SetParamCustomerGetOpenReq(_paramCustomerGetOpenReq *CustomerGetOpenReq) error {
	r._paramCustomerGetOpenReq = _paramCustomerGetOpenReq
	r.Set("param_customer_get_open_req", _paramCustomerGetOpenReq)
	return nil
}

// GetParamCustomerGetOpenReq ParamCustomerGetOpenReq Getter
func (r AlibabaalsccrmopencustomergetAPIRequest) GetParamCustomerGetOpenReq() *CustomerGetOpenReq {
	return r._paramCustomerGetOpenReq
}
