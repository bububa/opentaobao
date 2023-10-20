package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmpointquerypointflowAPIRequest 分页查询积分流水 API请求
// alibaba.alsc.crm.point.querypointflow
//
// 分页查询积分流水
type AlibabaalsccrmpointquerypointflowAPIRequest struct {
	model.Params
	// 入参
	_paramPageQueryPointFlowOpenReq *PageQueryPointFlowOpenReq
}

// NewAlibabaalsccrmpointquerypointflowRequest 初始化AlibabaalsccrmpointquerypointflowAPIRequest对象
func NewAlibabaalsccrmpointquerypointflowRequest() *AlibabaalsccrmpointquerypointflowAPIRequest {
	return &AlibabaalsccrmpointquerypointflowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmpointquerypointflowAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.point.querypointflow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmpointquerypointflowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmpointquerypointflowAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPageQueryPointFlowOpenReq is ParamPageQueryPointFlowOpenReq Setter
// 入参
func (r *AlibabaalsccrmpointquerypointflowAPIRequest) SetParamPageQueryPointFlowOpenReq(_paramPageQueryPointFlowOpenReq *PageQueryPointFlowOpenReq) error {
	r._paramPageQueryPointFlowOpenReq = _paramPageQueryPointFlowOpenReq
	r.Set("param_page_query_point_flow_open_req", _paramPageQueryPointFlowOpenReq)
	return nil
}

// GetParamPageQueryPointFlowOpenReq ParamPageQueryPointFlowOpenReq Getter
func (r AlibabaalsccrmpointquerypointflowAPIRequest) GetParamPageQueryPointFlowOpenReq() *PageQueryPointFlowOpenReq {
	return r._paramPageQueryPointFlowOpenReq
}
