package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmcustomercreateAPIRequest 创建顾客 API请求
// alibaba.alsc.crm.customer.create
//
// 开放本地生活创建顾客功能
type AlibabaalsccrmcustomercreateAPIRequest struct {
	model.Params
	// 创建顾客参数
	_paramCustomerCreateOpenReq *CustomerCreateOpenReq
}

// NewAlibabaalsccrmcustomercreateRequest 初始化AlibabaalsccrmcustomercreateAPIRequest对象
func NewAlibabaalsccrmcustomercreateRequest() *AlibabaalsccrmcustomercreateAPIRequest {
	return &AlibabaalsccrmcustomercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmcustomercreateAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.customer.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmcustomercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmcustomercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCustomerCreateOpenReq is ParamCustomerCreateOpenReq Setter
// 创建顾客参数
func (r *AlibabaalsccrmcustomercreateAPIRequest) SetParamCustomerCreateOpenReq(_paramCustomerCreateOpenReq *CustomerCreateOpenReq) error {
	r._paramCustomerCreateOpenReq = _paramCustomerCreateOpenReq
	r.Set("param_customer_create_open_req", _paramCustomerCreateOpenReq)
	return nil
}

// GetParamCustomerCreateOpenReq ParamCustomerCreateOpenReq Getter
func (r AlibabaalsccrmcustomercreateAPIRequest) GetParamCustomerCreateOpenReq() *CustomerCreateOpenReq {
	return r._paramCustomerCreateOpenReq
}
