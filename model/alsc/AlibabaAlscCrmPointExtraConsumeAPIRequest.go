package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmpointextraconsumeAPIRequest 积分补扣 API请求
// alibaba.alsc.crm.point.extra.consume
//
// 积分补扣
type AlibabaalsccrmpointextraconsumeAPIRequest struct {
	model.Params
	// 入参
	_paramExtraConsumePointOpenReq *ExtraConsumePointOpenReq
}

// NewAlibabaalsccrmpointextraconsumeRequest 初始化AlibabaalsccrmpointextraconsumeAPIRequest对象
func NewAlibabaalsccrmpointextraconsumeRequest() *AlibabaalsccrmpointextraconsumeAPIRequest {
	return &AlibabaalsccrmpointextraconsumeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmpointextraconsumeAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.point.extra.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmpointextraconsumeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmpointextraconsumeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamExtraConsumePointOpenReq is ParamExtraConsumePointOpenReq Setter
// 入参
func (r *AlibabaalsccrmpointextraconsumeAPIRequest) SetParamExtraConsumePointOpenReq(_paramExtraConsumePointOpenReq *ExtraConsumePointOpenReq) error {
	r._paramExtraConsumePointOpenReq = _paramExtraConsumePointOpenReq
	r.Set("param_extra_consume_point_open_req", _paramExtraConsumePointOpenReq)
	return nil
}

// GetParamExtraConsumePointOpenReq ParamExtraConsumePointOpenReq Getter
func (r AlibabaalsccrmpointextraconsumeAPIRequest) GetParamExtraConsumePointOpenReq() *ExtraConsumePointOpenReq {
	return r._paramExtraConsumePointOpenReq
}
