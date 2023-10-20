package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmpointconsumepointAPIRequest 积分抵现 API请求
// alibaba.alsc.crm.point.consumepoint
//
// 积分抵现
type AlibabaalsccrmpointconsumepointAPIRequest struct {
	model.Params
	// 入参
	_paramConsumePointOpenReq *ConsumePointOpenReq
}

// NewAlibabaalsccrmpointconsumepointRequest 初始化AlibabaalsccrmpointconsumepointAPIRequest对象
func NewAlibabaalsccrmpointconsumepointRequest() *AlibabaalsccrmpointconsumepointAPIRequest {
	return &AlibabaalsccrmpointconsumepointAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmpointconsumepointAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.point.consumepoint"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmpointconsumepointAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmpointconsumepointAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamConsumePointOpenReq is ParamConsumePointOpenReq Setter
// 入参
func (r *AlibabaalsccrmpointconsumepointAPIRequest) SetParamConsumePointOpenReq(_paramConsumePointOpenReq *ConsumePointOpenReq) error {
	r._paramConsumePointOpenReq = _paramConsumePointOpenReq
	r.Set("param_consume_point_open_req", _paramConsumePointOpenReq)
	return nil
}

// GetParamConsumePointOpenReq ParamConsumePointOpenReq Getter
func (r AlibabaalsccrmpointconsumepointAPIRequest) GetParamConsumePointOpenReq() *ConsumePointOpenReq {
	return r._paramConsumePointOpenReq
}
