package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmPointConsumepointAPIRequest
积分抵现 API请求
alibaba.alsc.crm.point.consumepoint

积分抵现 */
type AlibabaAlscCrmPointConsumepointAPIRequest struct {
	model.Params
	// 入参
	_paramConsumePointOpenReq *ConsumePointOpenReq
}

// NewAlibabaAlscCrmPointConsumepointRequest 初始化AlibabaAlscCrmPointConsumepointAPIRequest对象
func NewAlibabaAlscCrmPointConsumepointRequest() *AlibabaAlscCrmPointConsumepointAPIRequest {
	return &AlibabaAlscCrmPointConsumepointAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmPointConsumepointAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.point.consumepoint"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmPointConsumepointAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamConsumePointOpenReq Setter
// 入参
func (r *AlibabaAlscCrmPointConsumepointAPIRequest) SetParamConsumePointOpenReq(_paramConsumePointOpenReq *ConsumePointOpenReq) error {
	r._paramConsumePointOpenReq = _paramConsumePointOpenReq
	r.Set("param_consume_point_open_req", _paramConsumePointOpenReq)
	return nil
}

// Get ParamConsumePointOpenReq Getter
func (r AlibabaAlscCrmPointConsumepointAPIRequest) GetParamConsumePointOpenReq() *ConsumePointOpenReq {
	return r._paramConsumePointOpenReq
}
