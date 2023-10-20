package alsc

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
