package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmPointExtraConsumeAPIRequest 积分补扣 API请求
// alibaba.alsc.crm.point.extra.consume
//
// 积分补扣
type AlibabaAlscCrmPointExtraConsumeAPIRequest struct {
	model.Params
	// 入参
	_paramExtraConsumePointOpenReq *ExtraConsumePointOpenReq
}

// NewAlibabaAlscCrmPointExtraConsumeRequest 初始化AlibabaAlscCrmPointExtraConsumeAPIRequest对象
func NewAlibabaAlscCrmPointExtraConsumeRequest() *AlibabaAlscCrmPointExtraConsumeAPIRequest {
	return &AlibabaAlscCrmPointExtraConsumeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmPointExtraConsumeAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.point.extra.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmPointExtraConsumeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamExtraConsumePointOpenReq Setter
// 入参
func (r *AlibabaAlscCrmPointExtraConsumeAPIRequest) SetParamExtraConsumePointOpenReq(_paramExtraConsumePointOpenReq *ExtraConsumePointOpenReq) error {
	r._paramExtraConsumePointOpenReq = _paramExtraConsumePointOpenReq
	r.Set("param_extra_consume_point_open_req", _paramExtraConsumePointOpenReq)
	return nil
}

// Get ParamExtraConsumePointOpenReq Getter
func (r AlibabaAlscCrmPointExtraConsumeAPIRequest) GetParamExtraConsumePointOpenReq() *ExtraConsumePointOpenReq {
	return r._paramExtraConsumePointOpenReq
}
