package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRechargeDedutprecheckGetAPIRequest 储值核销预先校验 API请求
// alibaba.alsc.crm.recharge.dedutprecheck.get
//
// 储值核销预先校验接口
type AlibabaAlscCrmRechargeDedutprecheckGetAPIRequest struct {
	model.Params
	// 入参
	_paramDeductPreCheckOpenReq *DeductPreCheckOpenReq
}

// NewAlibabaAlscCrmRechargeDedutprecheckGetRequest 初始化AlibabaAlscCrmRechargeDedutprecheckGetAPIRequest对象
func NewAlibabaAlscCrmRechargeDedutprecheckGetRequest() *AlibabaAlscCrmRechargeDedutprecheckGetAPIRequest {
	return &AlibabaAlscCrmRechargeDedutprecheckGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeDedutprecheckGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.dedutprecheck.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeDedutprecheckGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmRechargeDedutprecheckGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamDeductPreCheckOpenReq is ParamDeductPreCheckOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeDedutprecheckGetAPIRequest) SetParamDeductPreCheckOpenReq(_paramDeductPreCheckOpenReq *DeductPreCheckOpenReq) error {
	r._paramDeductPreCheckOpenReq = _paramDeductPreCheckOpenReq
	r.Set("param_deduct_pre_check_open_req", _paramDeductPreCheckOpenReq)
	return nil
}

// GetParamDeductPreCheckOpenReq ParamDeductPreCheckOpenReq Getter
func (r AlibabaAlscCrmRechargeDedutprecheckGetAPIRequest) GetParamDeductPreCheckOpenReq() *DeductPreCheckOpenReq {
	return r._paramDeductPreCheckOpenReq
}
