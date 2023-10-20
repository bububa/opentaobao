package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmcardbindcustomerAPIRequest 卡号绑定顾客 API请求
// alibaba.alsc.crm.card.bindcustomer
//
// 为卡号绑定顾客
type AlibabaalsccrmcardbindcustomerAPIRequest struct {
	model.Params
	// 请求参数
	_paramBindCustomerOpenReq *BindCustomerOpenReq
}

// NewAlibabaalsccrmcardbindcustomerRequest 初始化AlibabaalsccrmcardbindcustomerAPIRequest对象
func NewAlibabaalsccrmcardbindcustomerRequest() *AlibabaalsccrmcardbindcustomerAPIRequest {
	return &AlibabaalsccrmcardbindcustomerAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmcardbindcustomerAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.bindcustomer"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmcardbindcustomerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmcardbindcustomerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBindCustomerOpenReq is ParamBindCustomerOpenReq Setter
// 请求参数
func (r *AlibabaalsccrmcardbindcustomerAPIRequest) SetParamBindCustomerOpenReq(_paramBindCustomerOpenReq *BindCustomerOpenReq) error {
	r._paramBindCustomerOpenReq = _paramBindCustomerOpenReq
	r.Set("param_bind_customer_open_req", _paramBindCustomerOpenReq)
	return nil
}

// GetParamBindCustomerOpenReq ParamBindCustomerOpenReq Getter
func (r AlibabaalsccrmcardbindcustomerAPIRequest) GetParamBindCustomerOpenReq() *BindCustomerOpenReq {
	return r._paramBindCustomerOpenReq
}
