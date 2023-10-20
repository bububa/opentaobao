package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest 充值退款 API请求
// alibaba.alsc.crm.recharge.uncharge.update
//
// 充值退款
type AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest struct {
	model.Params
	// 入参
	_paramUnchargeOpenReq *UnchargeOpenReq
}

// NewAlibabaAlscCrmRechargeUnchargeUpdateRequest 初始化AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest对象
func NewAlibabaAlscCrmRechargeUnchargeUpdateRequest() *AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest {
	return &AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest) Reset() {
	r._paramUnchargeOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.uncharge.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamUnchargeOpenReq is ParamUnchargeOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest) SetParamUnchargeOpenReq(_paramUnchargeOpenReq *UnchargeOpenReq) error {
	r._paramUnchargeOpenReq = _paramUnchargeOpenReq
	r.Set("param_uncharge_open_req", _paramUnchargeOpenReq)
	return nil
}

// GetParamUnchargeOpenReq ParamUnchargeOpenReq Getter
func (r AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest) GetParamUnchargeOpenReq() *UnchargeOpenReq {
	return r._paramUnchargeOpenReq
}

var poolAlibabaAlscCrmRechargeUnchargeUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmRechargeUnchargeUpdateRequest()
	},
}

// GetAlibabaAlscCrmRechargeUnchargeUpdateRequest 从 sync.Pool 获取 AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest
func GetAlibabaAlscCrmRechargeUnchargeUpdateAPIRequest() *AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest {
	return poolAlibabaAlscCrmRechargeUnchargeUpdateAPIRequest.Get().(*AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest)
}

// ReleaseAlibabaAlscCrmRechargeUnchargeUpdateAPIRequest 将 AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmRechargeUnchargeUpdateAPIRequest(v *AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmRechargeUnchargeUpdateAPIRequest.Put(v)
}
