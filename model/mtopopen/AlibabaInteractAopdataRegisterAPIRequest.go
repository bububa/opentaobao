package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractaopdataregisterAPIRequest 资源位数据推送接口 API请求
// alibaba.interact.aopdata.register
//
// 提供给isv，查询以及推送浮层资源位的三方活动数据
type AlibabainteractaopdataregisterAPIRequest struct {
	model.Params
	// 入参
	_paramTopIsvDecorateParam *TopIsvDecorateParam
}

// NewAlibabainteractaopdataregisterRequest 初始化AlibabainteractaopdataregisterAPIRequest对象
func NewAlibabainteractaopdataregisterRequest() *AlibabainteractaopdataregisterAPIRequest {
	return &AlibabainteractaopdataregisterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractaopdataregisterAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.aopdata.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractaopdataregisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractaopdataregisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTopIsvDecorateParam is ParamTopIsvDecorateParam Setter
// 入参
func (r *AlibabainteractaopdataregisterAPIRequest) SetParamTopIsvDecorateParam(_paramTopIsvDecorateParam *TopIsvDecorateParam) error {
	r._paramTopIsvDecorateParam = _paramTopIsvDecorateParam
	r.Set("param_top_isv_decorate_param", _paramTopIsvDecorateParam)
	return nil
}

// GetParamTopIsvDecorateParam ParamTopIsvDecorateParam Getter
func (r AlibabainteractaopdataregisterAPIRequest) GetParamTopIsvDecorateParam() *TopIsvDecorateParam {
	return r._paramTopIsvDecorateParam
}
