package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmPointExtrachargeAPIRequest 积分补录 API请求
// alibaba.alsc.crm.point.extracharge
//
// 积分补录
type AlibabaAlscCrmPointExtrachargeAPIRequest struct {
	model.Params
	// 入参
	_paramExtraChargePointOpenReq *ExtraChargePointOpenReq
}

// NewAlibabaAlscCrmPointExtrachargeRequest 初始化AlibabaAlscCrmPointExtrachargeAPIRequest对象
func NewAlibabaAlscCrmPointExtrachargeRequest() *AlibabaAlscCrmPointExtrachargeAPIRequest {
	return &AlibabaAlscCrmPointExtrachargeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmPointExtrachargeAPIRequest) Reset() {
	r._paramExtraChargePointOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmPointExtrachargeAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.point.extracharge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmPointExtrachargeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmPointExtrachargeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamExtraChargePointOpenReq is ParamExtraChargePointOpenReq Setter
// 入参
func (r *AlibabaAlscCrmPointExtrachargeAPIRequest) SetParamExtraChargePointOpenReq(_paramExtraChargePointOpenReq *ExtraChargePointOpenReq) error {
	r._paramExtraChargePointOpenReq = _paramExtraChargePointOpenReq
	r.Set("param_extra_charge_point_open_req", _paramExtraChargePointOpenReq)
	return nil
}

// GetParamExtraChargePointOpenReq ParamExtraChargePointOpenReq Getter
func (r AlibabaAlscCrmPointExtrachargeAPIRequest) GetParamExtraChargePointOpenReq() *ExtraChargePointOpenReq {
	return r._paramExtraChargePointOpenReq
}

var poolAlibabaAlscCrmPointExtrachargeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmPointExtrachargeRequest()
	},
}

// GetAlibabaAlscCrmPointExtrachargeRequest 从 sync.Pool 获取 AlibabaAlscCrmPointExtrachargeAPIRequest
func GetAlibabaAlscCrmPointExtrachargeAPIRequest() *AlibabaAlscCrmPointExtrachargeAPIRequest {
	return poolAlibabaAlscCrmPointExtrachargeAPIRequest.Get().(*AlibabaAlscCrmPointExtrachargeAPIRequest)
}

// ReleaseAlibabaAlscCrmPointExtrachargeAPIRequest 将 AlibabaAlscCrmPointExtrachargeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmPointExtrachargeAPIRequest(v *AlibabaAlscCrmPointExtrachargeAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmPointExtrachargeAPIRequest.Put(v)
}
