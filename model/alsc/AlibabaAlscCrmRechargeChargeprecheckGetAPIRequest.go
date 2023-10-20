package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest 储值账户充值前校验 API请求
// alibaba.alsc.crm.recharge.chargeprecheck.get
//
// 储值账户充值前校验接口
type AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest struct {
	model.Params
	// 入参
	_paramChargePreCheckOpenReq *ChargePreCheckOpenReq
}

// NewAlibabaAlscCrmRechargeChargeprecheckGetRequest 初始化AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest对象
func NewAlibabaAlscCrmRechargeChargeprecheckGetRequest() *AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest {
	return &AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest) Reset() {
	r._paramChargePreCheckOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.chargeprecheck.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamChargePreCheckOpenReq is ParamChargePreCheckOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest) SetParamChargePreCheckOpenReq(_paramChargePreCheckOpenReq *ChargePreCheckOpenReq) error {
	r._paramChargePreCheckOpenReq = _paramChargePreCheckOpenReq
	r.Set("param_charge_pre_check_open_req", _paramChargePreCheckOpenReq)
	return nil
}

// GetParamChargePreCheckOpenReq ParamChargePreCheckOpenReq Getter
func (r AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest) GetParamChargePreCheckOpenReq() *ChargePreCheckOpenReq {
	return r._paramChargePreCheckOpenReq
}

var poolAlibabaAlscCrmRechargeChargeprecheckGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmRechargeChargeprecheckGetRequest()
	},
}

// GetAlibabaAlscCrmRechargeChargeprecheckGetRequest 从 sync.Pool 获取 AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest
func GetAlibabaAlscCrmRechargeChargeprecheckGetAPIRequest() *AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest {
	return poolAlibabaAlscCrmRechargeChargeprecheckGetAPIRequest.Get().(*AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest)
}

// ReleaseAlibabaAlscCrmRechargeChargeprecheckGetAPIRequest 将 AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmRechargeChargeprecheckGetAPIRequest(v *AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmRechargeChargeprecheckGetAPIRequest.Put(v)
}
