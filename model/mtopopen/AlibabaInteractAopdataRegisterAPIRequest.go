package mtopopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractAopdataRegisterAPIRequest 资源位数据推送接口 API请求
// alibaba.interact.aopdata.register
//
// 提供给isv，查询以及推送浮层资源位的三方活动数据
type AlibabaInteractAopdataRegisterAPIRequest struct {
	model.Params
	// 入参
	_paramTopIsvDecorateParam *TopIsvDecorateParam
}

// NewAlibabaInteractAopdataRegisterRequest 初始化AlibabaInteractAopdataRegisterAPIRequest对象
func NewAlibabaInteractAopdataRegisterRequest() *AlibabaInteractAopdataRegisterAPIRequest {
	return &AlibabaInteractAopdataRegisterAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractAopdataRegisterAPIRequest) Reset() {
	r._paramTopIsvDecorateParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractAopdataRegisterAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.aopdata.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractAopdataRegisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractAopdataRegisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTopIsvDecorateParam is ParamTopIsvDecorateParam Setter
// 入参
func (r *AlibabaInteractAopdataRegisterAPIRequest) SetParamTopIsvDecorateParam(_paramTopIsvDecorateParam *TopIsvDecorateParam) error {
	r._paramTopIsvDecorateParam = _paramTopIsvDecorateParam
	r.Set("param_top_isv_decorate_param", _paramTopIsvDecorateParam)
	return nil
}

// GetParamTopIsvDecorateParam ParamTopIsvDecorateParam Getter
func (r AlibabaInteractAopdataRegisterAPIRequest) GetParamTopIsvDecorateParam() *TopIsvDecorateParam {
	return r._paramTopIsvDecorateParam
}

var poolAlibabaInteractAopdataRegisterAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractAopdataRegisterRequest()
	},
}

// GetAlibabaInteractAopdataRegisterRequest 从 sync.Pool 获取 AlibabaInteractAopdataRegisterAPIRequest
func GetAlibabaInteractAopdataRegisterAPIRequest() *AlibabaInteractAopdataRegisterAPIRequest {
	return poolAlibabaInteractAopdataRegisterAPIRequest.Get().(*AlibabaInteractAopdataRegisterAPIRequest)
}

// ReleaseAlibabaInteractAopdataRegisterAPIRequest 将 AlibabaInteractAopdataRegisterAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractAopdataRegisterAPIRequest(v *AlibabaInteractAopdataRegisterAPIRequest) {
	v.Reset()
	poolAlibabaInteractAopdataRegisterAPIRequest.Put(v)
}
