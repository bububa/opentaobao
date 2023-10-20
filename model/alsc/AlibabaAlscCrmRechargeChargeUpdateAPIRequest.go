package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRechargeChargeUpdateAPIRequest 储值充值 API请求
// alibaba.alsc.crm.recharge.charge.update
//
// 顾客储值账户充值
type AlibabaAlscCrmRechargeChargeUpdateAPIRequest struct {
	model.Params
	// 入参
	_paramRechargeOpenReq *RechargeOpenReq
}

// NewAlibabaAlscCrmRechargeChargeUpdateRequest 初始化AlibabaAlscCrmRechargeChargeUpdateAPIRequest对象
func NewAlibabaAlscCrmRechargeChargeUpdateRequest() *AlibabaAlscCrmRechargeChargeUpdateAPIRequest {
	return &AlibabaAlscCrmRechargeChargeUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmRechargeChargeUpdateAPIRequest) Reset() {
	r._paramRechargeOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeChargeUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.charge.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeChargeUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmRechargeChargeUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamRechargeOpenReq is ParamRechargeOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeChargeUpdateAPIRequest) SetParamRechargeOpenReq(_paramRechargeOpenReq *RechargeOpenReq) error {
	r._paramRechargeOpenReq = _paramRechargeOpenReq
	r.Set("param_recharge_open_req", _paramRechargeOpenReq)
	return nil
}

// GetParamRechargeOpenReq ParamRechargeOpenReq Getter
func (r AlibabaAlscCrmRechargeChargeUpdateAPIRequest) GetParamRechargeOpenReq() *RechargeOpenReq {
	return r._paramRechargeOpenReq
}

var poolAlibabaAlscCrmRechargeChargeUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmRechargeChargeUpdateRequest()
	},
}

// GetAlibabaAlscCrmRechargeChargeUpdateRequest 从 sync.Pool 获取 AlibabaAlscCrmRechargeChargeUpdateAPIRequest
func GetAlibabaAlscCrmRechargeChargeUpdateAPIRequest() *AlibabaAlscCrmRechargeChargeUpdateAPIRequest {
	return poolAlibabaAlscCrmRechargeChargeUpdateAPIRequest.Get().(*AlibabaAlscCrmRechargeChargeUpdateAPIRequest)
}

// ReleaseAlibabaAlscCrmRechargeChargeUpdateAPIRequest 将 AlibabaAlscCrmRechargeChargeUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmRechargeChargeUpdateAPIRequest(v *AlibabaAlscCrmRechargeChargeUpdateAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmRechargeChargeUpdateAPIRequest.Put(v)
}
