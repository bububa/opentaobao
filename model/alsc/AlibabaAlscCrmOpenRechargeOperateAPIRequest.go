package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmopenrechargeoperateAPIRequest 储值操作接口 API请求
// alibaba.alsc.crm.open.recharge.operate
//
// 储值操作接口
type AlibabaalsccrmopenrechargeoperateAPIRequest struct {
	model.Params
	// 储值操作参数
	_paramRechargeOperateOpenReq *RechargeOperateOpenReq
}

// NewAlibabaalsccrmopenrechargeoperateRequest 初始化AlibabaalsccrmopenrechargeoperateAPIRequest对象
func NewAlibabaalsccrmopenrechargeoperateRequest() *AlibabaalsccrmopenrechargeoperateAPIRequest {
	return &AlibabaalsccrmopenrechargeoperateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmopenrechargeoperateAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.open.recharge.operate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmopenrechargeoperateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmopenrechargeoperateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamRechargeOperateOpenReq is ParamRechargeOperateOpenReq Setter
// 储值操作参数
func (r *AlibabaalsccrmopenrechargeoperateAPIRequest) SetParamRechargeOperateOpenReq(_paramRechargeOperateOpenReq *RechargeOperateOpenReq) error {
	r._paramRechargeOperateOpenReq = _paramRechargeOperateOpenReq
	r.Set("param_recharge_operate_open_req", _paramRechargeOperateOpenReq)
	return nil
}

// GetParamRechargeOperateOpenReq ParamRechargeOperateOpenReq Getter
func (r AlibabaalsccrmopenrechargeoperateAPIRequest) GetParamRechargeOperateOpenReq() *RechargeOperateOpenReq {
	return r._paramRechargeOperateOpenReq
}
