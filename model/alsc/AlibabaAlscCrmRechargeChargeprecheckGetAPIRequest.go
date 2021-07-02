package alsc

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.chargeprecheck.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamChargePreCheckOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest) SetParamChargePreCheckOpenReq(_paramChargePreCheckOpenReq *ChargePreCheckOpenReq) error {
	r._paramChargePreCheckOpenReq = _paramChargePreCheckOpenReq
	r.Set("param_charge_pre_check_open_req", _paramChargePreCheckOpenReq)
	return nil
}

// Get ParamChargePreCheckOpenReq Getter
func (r AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest) GetParamChargePreCheckOpenReq() *ChargePreCheckOpenReq {
	return r._paramChargePreCheckOpenReq
}
