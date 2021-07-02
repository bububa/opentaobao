package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkReverseCreatrefundAPIRequest 逆向提交 API请求
// alibaba.wdk.reverse.creatrefund
//
// 逆向申请提交
type AlibabaWdkReverseCreatrefundAPIRequest struct {
	model.Params
	// CreateReverseReq
	_paramCreateReverseReq *CreateReverseReq
}

// NewAlibabaWdkReverseCreatrefundRequest 初始化AlibabaWdkReverseCreatrefundAPIRequest对象
func NewAlibabaWdkReverseCreatrefundRequest() *AlibabaWdkReverseCreatrefundAPIRequest {
	return &AlibabaWdkReverseCreatrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkReverseCreatrefundAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.reverse.creatrefund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkReverseCreatrefundAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamCreateReverseReq is ParamCreateReverseReq Setter
// CreateReverseReq
func (r *AlibabaWdkReverseCreatrefundAPIRequest) SetParamCreateReverseReq(_paramCreateReverseReq *CreateReverseReq) error {
	r._paramCreateReverseReq = _paramCreateReverseReq
	r.Set("param_create_reverse_req", _paramCreateReverseReq)
	return nil
}

// GetParamCreateReverseReq ParamCreateReverseReq Getter
func (r AlibabaWdkReverseCreatrefundAPIRequest) GetParamCreateReverseReq() *CreateReverseReq {
	return r._paramCreateReverseReq
}
