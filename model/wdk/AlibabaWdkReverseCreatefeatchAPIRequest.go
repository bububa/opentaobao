package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkreversecreatefeatchAPIRequest 逆向取货 API请求
// alibaba.wdk.reverse.createfeatch
//
// 发起逆向取货
type AlibabawdkreversecreatefeatchAPIRequest struct {
	model.Params
	// 入参
	_paramCreateFetchReq *CreateFetchReq
}

// NewAlibabawdkreversecreatefeatchRequest 初始化AlibabawdkreversecreatefeatchAPIRequest对象
func NewAlibabawdkreversecreatefeatchRequest() *AlibabawdkreversecreatefeatchAPIRequest {
	return &AlibabawdkreversecreatefeatchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkreversecreatefeatchAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.reverse.createfeatch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkreversecreatefeatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkreversecreatefeatchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCreateFetchReq is ParamCreateFetchReq Setter
// 入参
func (r *AlibabawdkreversecreatefeatchAPIRequest) SetParamCreateFetchReq(_paramCreateFetchReq *CreateFetchReq) error {
	r._paramCreateFetchReq = _paramCreateFetchReq
	r.Set("param_create_fetch_req", _paramCreateFetchReq)
	return nil
}

// GetParamCreateFetchReq ParamCreateFetchReq Getter
func (r AlibabawdkreversecreatefeatchAPIRequest) GetParamCreateFetchReq() *CreateFetchReq {
	return r._paramCreateFetchReq
}
