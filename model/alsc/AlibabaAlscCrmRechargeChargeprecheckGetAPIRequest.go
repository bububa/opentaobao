package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmrechargechargeprecheckgetAPIRequest 储值账户充值前校验 API请求
// alibaba.alsc.crm.recharge.chargeprecheck.get
//
// 储值账户充值前校验接口
type AlibabaalsccrmrechargechargeprecheckgetAPIRequest struct {
	model.Params
	// 入参
	_paramChargePreCheckOpenReq *ChargePreCheckOpenReq
}

// NewAlibabaalsccrmrechargechargeprecheckgetRequest 初始化AlibabaalsccrmrechargechargeprecheckgetAPIRequest对象
func NewAlibabaalsccrmrechargechargeprecheckgetRequest() *AlibabaalsccrmrechargechargeprecheckgetAPIRequest {
	return &AlibabaalsccrmrechargechargeprecheckgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmrechargechargeprecheckgetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.chargeprecheck.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmrechargechargeprecheckgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmrechargechargeprecheckgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamChargePreCheckOpenReq is ParamChargePreCheckOpenReq Setter
// 入参
func (r *AlibabaalsccrmrechargechargeprecheckgetAPIRequest) SetParamChargePreCheckOpenReq(_paramChargePreCheckOpenReq *ChargePreCheckOpenReq) error {
	r._paramChargePreCheckOpenReq = _paramChargePreCheckOpenReq
	r.Set("param_charge_pre_check_open_req", _paramChargePreCheckOpenReq)
	return nil
}

// GetParamChargePreCheckOpenReq ParamChargePreCheckOpenReq Getter
func (r AlibabaalsccrmrechargechargeprecheckgetAPIRequest) GetParamChargePreCheckOpenReq() *ChargePreCheckOpenReq {
	return r._paramChargePreCheckOpenReq
}
