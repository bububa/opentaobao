package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkreverseloadfeatchorderAPIRequest 取货详情 API请求
// alibaba.wdk.reverse.load.featchorder
//
// 取货详情
type AlibabawdkreverseloadfeatchorderAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramLoadFetchReq *LoadFetchReq
}

// NewAlibabawdkreverseloadfeatchorderRequest 初始化AlibabawdkreverseloadfeatchorderAPIRequest对象
func NewAlibabawdkreverseloadfeatchorderRequest() *AlibabawdkreverseloadfeatchorderAPIRequest {
	return &AlibabawdkreverseloadfeatchorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkreverseloadfeatchorderAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.reverse.load.featchorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkreverseloadfeatchorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkreverseloadfeatchorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamLoadFetchReq is ParamLoadFetchReq Setter
// 系统自动生成
func (r *AlibabawdkreverseloadfeatchorderAPIRequest) SetParamLoadFetchReq(_paramLoadFetchReq *LoadFetchReq) error {
	r._paramLoadFetchReq = _paramLoadFetchReq
	r.Set("param_load_fetch_req", _paramLoadFetchReq)
	return nil
}

// GetParamLoadFetchReq ParamLoadFetchReq Getter
func (r AlibabawdkreverseloadfeatchorderAPIRequest) GetParamLoadFetchReq() *LoadFetchReq {
	return r._paramLoadFetchReq
}
