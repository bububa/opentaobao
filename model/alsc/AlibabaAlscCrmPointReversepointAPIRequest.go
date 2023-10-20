package alsc

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmPointReversepointAPIRequest) Reset() {
	r._paramReverseConsumePointOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmPointReversepointAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.point.reversepoint"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmPointReversepointAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmPointReversepointAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamReverseConsumePointOpenReq is ParamReverseConsumePointOpenReq Setter
// 入参
func (r *AlibabaAlscCrmPointReversepointAPIRequest) SetParamReverseConsumePointOpenReq(_paramReverseConsumePointOpenReq *ReverseConsumePointOpenReq) error {
	r._paramReverseConsumePointOpenReq = _paramReverseConsumePointOpenReq
	r.Set("param_reverse_consume_point_open_req", _paramReverseConsumePointOpenReq)
	return nil
}

// GetParamReverseConsumePointOpenReq ParamReverseConsumePointOpenReq Getter
func (r AlibabaAlscCrmPointReversepointAPIRequest) GetParamReverseConsumePointOpenReq() *ReverseConsumePointOpenReq {
	return r._paramReverseConsumePointOpenReq
}

var poolAlibabaAlscCrmPointReversepointAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmPointReversepointRequest()
	},
}

// GetAlibabaAlscCrmPointReversepointRequest 从 sync.Pool 获取 AlibabaAlscCrmPointReversepointAPIRequest
func GetAlibabaAlscCrmPointReversepointAPIRequest() *AlibabaAlscCrmPointReversepointAPIRequest {
	return poolAlibabaAlscCrmPointReversepointAPIRequest.Get().(*AlibabaAlscCrmPointReversepointAPIRequest)
}

// ReleaseAlibabaAlscCrmPointReversepointAPIRequest 将 AlibabaAlscCrmPointReversepointAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmPointReversepointAPIRequest(v *AlibabaAlscCrmPointReversepointAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmPointReversepointAPIRequest.Put(v)
}
