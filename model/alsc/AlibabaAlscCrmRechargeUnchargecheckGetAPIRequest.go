package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest 储值账户退充值校验 API请求
// alibaba.alsc.crm.recharge.unchargecheck.get
//
// 储值账户退充值校验接口
type AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest struct {
	model.Params
	// 入参
	_paramUnchargeCheckOpenReq *UnchargeCheckOpenReq
}

// NewAlibabaAlscCrmRechargeUnchargecheckGetRequest 初始化AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest对象
func NewAlibabaAlscCrmRechargeUnchargecheckGetRequest() *AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest {
	return &AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.unchargecheck.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamUnchargeCheckOpenReq is ParamUnchargeCheckOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest) SetParamUnchargeCheckOpenReq(_paramUnchargeCheckOpenReq *UnchargeCheckOpenReq) error {
	r._paramUnchargeCheckOpenReq = _paramUnchargeCheckOpenReq
	r.Set("param_uncharge_check_open_req", _paramUnchargeCheckOpenReq)
	return nil
}

// GetParamUnchargeCheckOpenReq ParamUnchargeCheckOpenReq Getter
func (r AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest) GetParamUnchargeCheckOpenReq() *UnchargeCheckOpenReq {
	return r._paramUnchargeCheckOpenReq
}
