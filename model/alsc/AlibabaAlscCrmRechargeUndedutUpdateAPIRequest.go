package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmrechargeundedutupdateAPIRequest 储值消费退款(逆向) API请求
// alibaba.alsc.crm.recharge.undedut.update
//
// 新增储值消费退款接口
type AlibabaalsccrmrechargeundedutupdateAPIRequest struct {
	model.Params
	// 入参
	_paramUndedutOpenReq *UndedutOpenReq
}

// NewAlibabaalsccrmrechargeundedutupdateRequest 初始化AlibabaalsccrmrechargeundedutupdateAPIRequest对象
func NewAlibabaalsccrmrechargeundedutupdateRequest() *AlibabaalsccrmrechargeundedutupdateAPIRequest {
	return &AlibabaalsccrmrechargeundedutupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmrechargeundedutupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.undedut.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmrechargeundedutupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmrechargeundedutupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamUndedutOpenReq is ParamUndedutOpenReq Setter
// 入参
func (r *AlibabaalsccrmrechargeundedutupdateAPIRequest) SetParamUndedutOpenReq(_paramUndedutOpenReq *UndedutOpenReq) error {
	r._paramUndedutOpenReq = _paramUndedutOpenReq
	r.Set("param_undedut_open_req", _paramUndedutOpenReq)
	return nil
}

// GetParamUndedutOpenReq ParamUndedutOpenReq Getter
func (r AlibabaalsccrmrechargeundedutupdateAPIRequest) GetParamUndedutOpenReq() *UndedutOpenReq {
	return r._paramUndedutOpenReq
}
