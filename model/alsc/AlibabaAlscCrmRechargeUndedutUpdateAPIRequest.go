package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmRechargeUndedutUpdateAPIRequest
储值消费退款(逆向) API请求
alibaba.alsc.crm.recharge.undedut.update

新增储值消费退款接口 */
type AlibabaAlscCrmRechargeUndedutUpdateAPIRequest struct {
	model.Params
	// 入参
	_paramUndedutOpenReq *UndedutOpenReq
}

// NewAlibabaAlscCrmRechargeUndedutUpdateRequest 初始化AlibabaAlscCrmRechargeUndedutUpdateAPIRequest对象
func NewAlibabaAlscCrmRechargeUndedutUpdateRequest() *AlibabaAlscCrmRechargeUndedutUpdateAPIRequest {
	return &AlibabaAlscCrmRechargeUndedutUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeUndedutUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.undedut.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeUndedutUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamUndedutOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeUndedutUpdateAPIRequest) SetParamUndedutOpenReq(_paramUndedutOpenReq *UndedutOpenReq) error {
	r._paramUndedutOpenReq = _paramUndedutOpenReq
	r.Set("param_undedut_open_req", _paramUndedutOpenReq)
	return nil
}

// Get ParamUndedutOpenReq Getter
func (r AlibabaAlscCrmRechargeUndedutUpdateAPIRequest) GetParamUndedutOpenReq() *UndedutOpenReq {
	return r._paramUndedutOpenReq
}
