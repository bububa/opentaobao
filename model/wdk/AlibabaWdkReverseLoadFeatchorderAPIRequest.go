package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkReverseLoadFeatchorderAPIRequest 取货详情 API请求
// alibaba.wdk.reverse.load.featchorder
//
// 取货详情
type AlibabaWdkReverseLoadFeatchorderAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramLoadFetchReq *LoadFetchReq
}

// NewAlibabaWdkReverseLoadFeatchorderRequest 初始化AlibabaWdkReverseLoadFeatchorderAPIRequest对象
func NewAlibabaWdkReverseLoadFeatchorderRequest() *AlibabaWdkReverseLoadFeatchorderAPIRequest {
	return &AlibabaWdkReverseLoadFeatchorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkReverseLoadFeatchorderAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.reverse.load.featchorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkReverseLoadFeatchorderAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamLoadFetchReq is ParamLoadFetchReq Setter
// 系统自动生成
func (r *AlibabaWdkReverseLoadFeatchorderAPIRequest) SetParamLoadFetchReq(_paramLoadFetchReq *LoadFetchReq) error {
	r._paramLoadFetchReq = _paramLoadFetchReq
	r.Set("param_load_fetch_req", _paramLoadFetchReq)
	return nil
}

// GetParamLoadFetchReq ParamLoadFetchReq Getter
func (r AlibabaWdkReverseLoadFeatchorderAPIRequest) GetParamLoadFetchReq() *LoadFetchReq {
	return r._paramLoadFetchReq
}
