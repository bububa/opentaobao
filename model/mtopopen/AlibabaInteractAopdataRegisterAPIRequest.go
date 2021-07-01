package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractAopdataRegisterAPIRequest
资源位数据推送接口 API请求
alibaba.interact.aopdata.register

提供给isv，查询以及推送浮层资源位的三方活动数据 */
type AlibabaInteractAopdataRegisterAPIRequest struct {
	model.Params
	// 入参
	_paramTopIsvDecorateParam *TopIsvDecorateParam
}

// NewAlibabaInteractAopdataRegisterRequest 初始化AlibabaInteractAopdataRegisterAPIRequest对象
func NewAlibabaInteractAopdataRegisterRequest() *AlibabaInteractAopdataRegisterAPIRequest {
	return &AlibabaInteractAopdataRegisterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractAopdataRegisterAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.aopdata.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractAopdataRegisterAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamTopIsvDecorateParam Setter
// 入参
func (r *AlibabaInteractAopdataRegisterAPIRequest) SetParamTopIsvDecorateParam(_paramTopIsvDecorateParam *TopIsvDecorateParam) error {
	r._paramTopIsvDecorateParam = _paramTopIsvDecorateParam
	r.Set("param_top_isv_decorate_param", _paramTopIsvDecorateParam)
	return nil
}

// Get ParamTopIsvDecorateParam Getter
func (r AlibabaInteractAopdataRegisterAPIRequest) GetParamTopIsvDecorateParam() *TopIsvDecorateParam {
	return r._paramTopIsvDecorateParam
}
