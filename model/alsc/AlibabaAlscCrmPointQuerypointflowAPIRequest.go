package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmPointQuerypointflowAPIRequest 分页查询积分流水 API请求
// alibaba.alsc.crm.point.querypointflow
//
// 分页查询积分流水
type AlibabaAlscCrmPointQuerypointflowAPIRequest struct {
	model.Params
	// 入参
	_paramPageQueryPointFlowOpenReq *PageQueryPointFlowOpenReq
}

// NewAlibabaAlscCrmPointQuerypointflowRequest 初始化AlibabaAlscCrmPointQuerypointflowAPIRequest对象
func NewAlibabaAlscCrmPointQuerypointflowRequest() *AlibabaAlscCrmPointQuerypointflowAPIRequest {
	return &AlibabaAlscCrmPointQuerypointflowAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmPointQuerypointflowAPIRequest) Reset() {
	r._paramPageQueryPointFlowOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmPointQuerypointflowAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.point.querypointflow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmPointQuerypointflowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmPointQuerypointflowAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPageQueryPointFlowOpenReq is ParamPageQueryPointFlowOpenReq Setter
// 入参
func (r *AlibabaAlscCrmPointQuerypointflowAPIRequest) SetParamPageQueryPointFlowOpenReq(_paramPageQueryPointFlowOpenReq *PageQueryPointFlowOpenReq) error {
	r._paramPageQueryPointFlowOpenReq = _paramPageQueryPointFlowOpenReq
	r.Set("param_page_query_point_flow_open_req", _paramPageQueryPointFlowOpenReq)
	return nil
}

// GetParamPageQueryPointFlowOpenReq ParamPageQueryPointFlowOpenReq Getter
func (r AlibabaAlscCrmPointQuerypointflowAPIRequest) GetParamPageQueryPointFlowOpenReq() *PageQueryPointFlowOpenReq {
	return r._paramPageQueryPointFlowOpenReq
}

var poolAlibabaAlscCrmPointQuerypointflowAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmPointQuerypointflowRequest()
	},
}

// GetAlibabaAlscCrmPointQuerypointflowRequest 从 sync.Pool 获取 AlibabaAlscCrmPointQuerypointflowAPIRequest
func GetAlibabaAlscCrmPointQuerypointflowAPIRequest() *AlibabaAlscCrmPointQuerypointflowAPIRequest {
	return poolAlibabaAlscCrmPointQuerypointflowAPIRequest.Get().(*AlibabaAlscCrmPointQuerypointflowAPIRequest)
}

// ReleaseAlibabaAlscCrmPointQuerypointflowAPIRequest 将 AlibabaAlscCrmPointQuerypointflowAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmPointQuerypointflowAPIRequest(v *AlibabaAlscCrmPointQuerypointflowAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmPointQuerypointflowAPIRequest.Put(v)
}
