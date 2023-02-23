package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRechargeQryruleAPIRequest 储值规则下行 API请求
// alibaba.alsc.crm.recharge.qryrule
//
// 储值规则下行
type AlibabaAlscCrmRechargeQryruleAPIRequest struct {
	model.Params
	// 请求对象
	_paramPullRechargeRuleByShopReq *PullRechargeRuleByShopReq
}

// NewAlibabaAlscCrmRechargeQryruleRequest 初始化AlibabaAlscCrmRechargeQryruleAPIRequest对象
func NewAlibabaAlscCrmRechargeQryruleRequest() *AlibabaAlscCrmRechargeQryruleAPIRequest {
	return &AlibabaAlscCrmRechargeQryruleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeQryruleAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.qryrule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeQryruleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmRechargeQryruleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPullRechargeRuleByShopReq is ParamPullRechargeRuleByShopReq Setter
// 请求对象
func (r *AlibabaAlscCrmRechargeQryruleAPIRequest) SetParamPullRechargeRuleByShopReq(_paramPullRechargeRuleByShopReq *PullRechargeRuleByShopReq) error {
	r._paramPullRechargeRuleByShopReq = _paramPullRechargeRuleByShopReq
	r.Set("param_pull_recharge_rule_by_shop_req", _paramPullRechargeRuleByShopReq)
	return nil
}

// GetParamPullRechargeRuleByShopReq ParamPullRechargeRuleByShopReq Getter
func (r AlibabaAlscCrmRechargeQryruleAPIRequest) GetParamPullRechargeRuleByShopReq() *PullRechargeRuleByShopReq {
	return r._paramPullRechargeRuleByShopReq
}
