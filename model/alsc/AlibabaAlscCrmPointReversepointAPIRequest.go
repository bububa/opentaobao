package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmPointReversepointAPIRequest 积分消费回退 API请求
// alibaba.alsc.crm.point.reversepoint
//
// 积分消费回退
type AlibabaAlscCrmPointReversepointAPIRequest struct {
	model.Params
	// 入参
	_paramReverseConsumePointOpenReq *ReverseConsumePointOpenReq
}

// NewAlibabaAlscCrmPointReversepointRequest 初始化AlibabaAlscCrmPointReversepointAPIRequest对象
func NewAlibabaAlscCrmPointReversepointRequest() *AlibabaAlscCrmPointReversepointAPIRequest {
	return &AlibabaAlscCrmPointReversepointAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmPointReversepointAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.point.reversepoint"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmPointReversepointAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamReverseConsumePointOpenReq Setter
// 入参
func (r *AlibabaAlscCrmPointReversepointAPIRequest) SetParamReverseConsumePointOpenReq(_paramReverseConsumePointOpenReq *ReverseConsumePointOpenReq) error {
	r._paramReverseConsumePointOpenReq = _paramReverseConsumePointOpenReq
	r.Set("param_reverse_consume_point_open_req", _paramReverseConsumePointOpenReq)
	return nil
}

// Get ParamReverseConsumePointOpenReq Getter
func (r AlibabaAlscCrmPointReversepointAPIRequest) GetParamReverseConsumePointOpenReq() *ReverseConsumePointOpenReq {
	return r._paramReverseConsumePointOpenReq
}
