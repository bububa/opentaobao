package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmcustomergetAPIRequest 查询顾客详情 API请求
// alibaba.alsc.crm.customer.get
//
// 查询顾客详情
type AlibabaalsccrmcustomergetAPIRequest struct {
	model.Params
	// 顾客详情查询条件
	_paramCustomerIdQueryOpenReq *CustomerIdQueryOpenReq
}

// NewAlibabaalsccrmcustomergetRequest 初始化AlibabaalsccrmcustomergetAPIRequest对象
func NewAlibabaalsccrmcustomergetRequest() *AlibabaalsccrmcustomergetAPIRequest {
	return &AlibabaalsccrmcustomergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmcustomergetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.customer.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmcustomergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmcustomergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCustomerIdQueryOpenReq is ParamCustomerIdQueryOpenReq Setter
// 顾客详情查询条件
func (r *AlibabaalsccrmcustomergetAPIRequest) SetParamCustomerIdQueryOpenReq(_paramCustomerIdQueryOpenReq *CustomerIdQueryOpenReq) error {
	r._paramCustomerIdQueryOpenReq = _paramCustomerIdQueryOpenReq
	r.Set("param_customer_id_query_open_req", _paramCustomerIdQueryOpenReq)
	return nil
}

// GetParamCustomerIdQueryOpenReq ParamCustomerIdQueryOpenReq Getter
func (r AlibabaalsccrmcustomergetAPIRequest) GetParamCustomerIdQueryOpenReq() *CustomerIdQueryOpenReq {
	return r._paramCustomerIdQueryOpenReq
}
