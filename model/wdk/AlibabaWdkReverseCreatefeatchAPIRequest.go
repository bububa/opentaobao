package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkReverseCreatefeatchAPIRequest 逆向取货 API请求
// alibaba.wdk.reverse.createfeatch
//
// 发起逆向取货
type AlibabaWdkReverseCreatefeatchAPIRequest struct {
	model.Params
	// 入参
	_paramCreateFetchReq *CreateFetchReq
}

// NewAlibabaWdkReverseCreatefeatchRequest 初始化AlibabaWdkReverseCreatefeatchAPIRequest对象
func NewAlibabaWdkReverseCreatefeatchRequest() *AlibabaWdkReverseCreatefeatchAPIRequest {
	return &AlibabaWdkReverseCreatefeatchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkReverseCreatefeatchAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.reverse.createfeatch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkReverseCreatefeatchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamCreateFetchReq is ParamCreateFetchReq Setter
// 入参
func (r *AlibabaWdkReverseCreatefeatchAPIRequest) SetParamCreateFetchReq(_paramCreateFetchReq *CreateFetchReq) error {
	r._paramCreateFetchReq = _paramCreateFetchReq
	r.Set("param_create_fetch_req", _paramCreateFetchReq)
	return nil
}

// GetParamCreateFetchReq ParamCreateFetchReq Getter
func (r AlibabaWdkReverseCreatefeatchAPIRequest) GetParamCreateFetchReq() *CreateFetchReq {
	return r._paramCreateFetchReq
}
