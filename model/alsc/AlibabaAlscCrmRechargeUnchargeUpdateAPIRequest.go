package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest 充值退款 API请求
// alibaba.alsc.crm.recharge.uncharge.update
//
// 充值退款
type AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest struct {
	model.Params
	// 入参
	_paramUnchargeOpenReq *UnchargeOpenReq
}

// NewAlibabaAlscCrmRechargeUnchargeUpdateRequest 初始化AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest对象
func NewAlibabaAlscCrmRechargeUnchargeUpdateRequest() *AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest {
	return &AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.uncharge.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamUnchargeOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest) SetParamUnchargeOpenReq(_paramUnchargeOpenReq *UnchargeOpenReq) error {
	r._paramUnchargeOpenReq = _paramUnchargeOpenReq
	r.Set("param_uncharge_open_req", _paramUnchargeOpenReq)
	return nil
}

// Get ParamUnchargeOpenReq Getter
func (r AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest) GetParamUnchargeOpenReq() *UnchargeOpenReq {
	return r._paramUnchargeOpenReq
}
