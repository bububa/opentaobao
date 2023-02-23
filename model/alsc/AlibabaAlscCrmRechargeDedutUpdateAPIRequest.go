package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRechargeDedutUpdateAPIRequest 储值消费 API请求
// alibaba.alsc.crm.recharge.dedut.update
//
// 增加储值消费接口
type AlibabaAlscCrmRechargeDedutUpdateAPIRequest struct {
	model.Params
	// 入参
	_paramDedutOpenReq *DedutOpenReq
}

// NewAlibabaAlscCrmRechargeDedutUpdateRequest 初始化AlibabaAlscCrmRechargeDedutUpdateAPIRequest对象
func NewAlibabaAlscCrmRechargeDedutUpdateRequest() *AlibabaAlscCrmRechargeDedutUpdateAPIRequest {
	return &AlibabaAlscCrmRechargeDedutUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeDedutUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.dedut.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeDedutUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmRechargeDedutUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamDedutOpenReq is ParamDedutOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeDedutUpdateAPIRequest) SetParamDedutOpenReq(_paramDedutOpenReq *DedutOpenReq) error {
	r._paramDedutOpenReq = _paramDedutOpenReq
	r.Set("param_dedut_open_req", _paramDedutOpenReq)
	return nil
}

// GetParamDedutOpenReq ParamDedutOpenReq Getter
func (r AlibabaAlscCrmRechargeDedutUpdateAPIRequest) GetParamDedutOpenReq() *DedutOpenReq {
	return r._paramDedutOpenReq
}
