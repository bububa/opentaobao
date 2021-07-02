package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest 储值流水详细 API请求
// alibaba.alsc.crm.recharge.account.flowdetail.get
//
// 查询储值流水详细接口
type AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest struct {
	model.Params
	// 入参
	_paramQueryRechargeAccountFlowOpenReq *QueryRechargeAccountFlowOpenReq
}

// NewAlibabaAlscCrmRechargeAccountFlowdetailGetRequest 初始化AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest对象
func NewAlibabaAlscCrmRechargeAccountFlowdetailGetRequest() *AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest {
	return &AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.account.flowdetail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamQueryRechargeAccountFlowOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest) SetParamQueryRechargeAccountFlowOpenReq(_paramQueryRechargeAccountFlowOpenReq *QueryRechargeAccountFlowOpenReq) error {
	r._paramQueryRechargeAccountFlowOpenReq = _paramQueryRechargeAccountFlowOpenReq
	r.Set("param_query_recharge_account_flow_open_req", _paramQueryRechargeAccountFlowOpenReq)
	return nil
}

// Get ParamQueryRechargeAccountFlowOpenReq Getter
func (r AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest) GetParamQueryRechargeAccountFlowOpenReq() *QueryRechargeAccountFlowOpenReq {
	return r._paramQueryRechargeAccountFlowOpenReq
}
