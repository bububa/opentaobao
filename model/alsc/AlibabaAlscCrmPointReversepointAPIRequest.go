package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmpointreversepointAPIRequest 积分消费回退 API请求
// alibaba.alsc.crm.point.reversepoint
//
// 积分消费回退
type AlibabaalsccrmpointreversepointAPIRequest struct {
	model.Params
	// 入参
	_paramReverseConsumePointOpenReq *ReverseConsumePointOpenReq
}

// NewAlibabaalsccrmpointreversepointRequest 初始化AlibabaalsccrmpointreversepointAPIRequest对象
func NewAlibabaalsccrmpointreversepointRequest() *AlibabaalsccrmpointreversepointAPIRequest {
	return &AlibabaalsccrmpointreversepointAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmpointreversepointAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.point.reversepoint"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmpointreversepointAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmpointreversepointAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamReverseConsumePointOpenReq is ParamReverseConsumePointOpenReq Setter
// 入参
func (r *AlibabaalsccrmpointreversepointAPIRequest) SetParamReverseConsumePointOpenReq(_paramReverseConsumePointOpenReq *ReverseConsumePointOpenReq) error {
	r._paramReverseConsumePointOpenReq = _paramReverseConsumePointOpenReq
	r.Set("param_reverse_consume_point_open_req", _paramReverseConsumePointOpenReq)
	return nil
}

// GetParamReverseConsumePointOpenReq ParamReverseConsumePointOpenReq Getter
func (r AlibabaalsccrmpointreversepointAPIRequest) GetParamReverseConsumePointOpenReq() *ReverseConsumePointOpenReq {
	return r._paramReverseConsumePointOpenReq
}
