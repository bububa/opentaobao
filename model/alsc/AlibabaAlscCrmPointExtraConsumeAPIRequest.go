package alsc

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmPointExtraConsumeAPIRequest) Reset() {
	r._paramExtraConsumePointOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmPointExtraConsumeAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.point.extra.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmPointExtraConsumeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmPointExtraConsumeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamExtraConsumePointOpenReq is ParamExtraConsumePointOpenReq Setter
// 入参
func (r *AlibabaAlscCrmPointExtraConsumeAPIRequest) SetParamExtraConsumePointOpenReq(_paramExtraConsumePointOpenReq *ExtraConsumePointOpenReq) error {
	r._paramExtraConsumePointOpenReq = _paramExtraConsumePointOpenReq
	r.Set("param_extra_consume_point_open_req", _paramExtraConsumePointOpenReq)
	return nil
}

// GetParamExtraConsumePointOpenReq ParamExtraConsumePointOpenReq Getter
func (r AlibabaAlscCrmPointExtraConsumeAPIRequest) GetParamExtraConsumePointOpenReq() *ExtraConsumePointOpenReq {
	return r._paramExtraConsumePointOpenReq
}

var poolAlibabaAlscCrmPointExtraConsumeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmPointExtraConsumeRequest()
	},
}

// GetAlibabaAlscCrmPointExtraConsumeRequest 从 sync.Pool 获取 AlibabaAlscCrmPointExtraConsumeAPIRequest
func GetAlibabaAlscCrmPointExtraConsumeAPIRequest() *AlibabaAlscCrmPointExtraConsumeAPIRequest {
	return poolAlibabaAlscCrmPointExtraConsumeAPIRequest.Get().(*AlibabaAlscCrmPointExtraConsumeAPIRequest)
}

// ReleaseAlibabaAlscCrmPointExtraConsumeAPIRequest 将 AlibabaAlscCrmPointExtraConsumeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmPointExtraConsumeAPIRequest(v *AlibabaAlscCrmPointExtraConsumeAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmPointExtraConsumeAPIRequest.Put(v)
}
