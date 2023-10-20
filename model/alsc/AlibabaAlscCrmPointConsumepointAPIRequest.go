package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmPointConsumepointAPIRequest 积分抵现 API请求
// alibaba.alsc.crm.point.consumepoint
//
// 积分抵现
type AlibabaAlscCrmPointConsumepointAPIRequest struct {
	model.Params
	// 入参
	_paramConsumePointOpenReq *ConsumePointOpenReq
}

// NewAlibabaAlscCrmPointConsumepointRequest 初始化AlibabaAlscCrmPointConsumepointAPIRequest对象
func NewAlibabaAlscCrmPointConsumepointRequest() *AlibabaAlscCrmPointConsumepointAPIRequest {
	return &AlibabaAlscCrmPointConsumepointAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmPointConsumepointAPIRequest) Reset() {
	r._paramConsumePointOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmPointConsumepointAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.point.consumepoint"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmPointConsumepointAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmPointConsumepointAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamConsumePointOpenReq is ParamConsumePointOpenReq Setter
// 入参
func (r *AlibabaAlscCrmPointConsumepointAPIRequest) SetParamConsumePointOpenReq(_paramConsumePointOpenReq *ConsumePointOpenReq) error {
	r._paramConsumePointOpenReq = _paramConsumePointOpenReq
	r.Set("param_consume_point_open_req", _paramConsumePointOpenReq)
	return nil
}

// GetParamConsumePointOpenReq ParamConsumePointOpenReq Getter
func (r AlibabaAlscCrmPointConsumepointAPIRequest) GetParamConsumePointOpenReq() *ConsumePointOpenReq {
	return r._paramConsumePointOpenReq
}

var poolAlibabaAlscCrmPointConsumepointAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmPointConsumepointRequest()
	},
}

// GetAlibabaAlscCrmPointConsumepointRequest 从 sync.Pool 获取 AlibabaAlscCrmPointConsumepointAPIRequest
func GetAlibabaAlscCrmPointConsumepointAPIRequest() *AlibabaAlscCrmPointConsumepointAPIRequest {
	return poolAlibabaAlscCrmPointConsumepointAPIRequest.Get().(*AlibabaAlscCrmPointConsumepointAPIRequest)
}

// ReleaseAlibabaAlscCrmPointConsumepointAPIRequest 将 AlibabaAlscCrmPointConsumepointAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmPointConsumepointAPIRequest(v *AlibabaAlscCrmPointConsumepointAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmPointConsumepointAPIRequest.Put(v)
}
