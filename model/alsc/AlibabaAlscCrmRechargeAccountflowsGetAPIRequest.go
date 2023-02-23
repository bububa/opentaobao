package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRechargeAccountflowsGetAPIRequest 分页查询储值流水 API请求
// alibaba.alsc.crm.recharge.accountflows.get
//
// 增加分页查询储值流水接口
type AlibabaAlscCrmRechargeAccountflowsGetAPIRequest struct {
	model.Params
	// 入参
	_paramPageQueryAccountFlowsOpenReq *PageQueryAccountFlowsOpenReq
}

// NewAlibabaAlscCrmRechargeAccountflowsGetRequest 初始化AlibabaAlscCrmRechargeAccountflowsGetAPIRequest对象
func NewAlibabaAlscCrmRechargeAccountflowsGetRequest() *AlibabaAlscCrmRechargeAccountflowsGetAPIRequest {
	return &AlibabaAlscCrmRechargeAccountflowsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeAccountflowsGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.accountflows.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeAccountflowsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmRechargeAccountflowsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPageQueryAccountFlowsOpenReq is ParamPageQueryAccountFlowsOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeAccountflowsGetAPIRequest) SetParamPageQueryAccountFlowsOpenReq(_paramPageQueryAccountFlowsOpenReq *PageQueryAccountFlowsOpenReq) error {
	r._paramPageQueryAccountFlowsOpenReq = _paramPageQueryAccountFlowsOpenReq
	r.Set("param_page_query_account_flows_open_req", _paramPageQueryAccountFlowsOpenReq)
	return nil
}

// GetParamPageQueryAccountFlowsOpenReq ParamPageQueryAccountFlowsOpenReq Getter
func (r AlibabaAlscCrmRechargeAccountflowsGetAPIRequest) GetParamPageQueryAccountFlowsOpenReq() *PageQueryAccountFlowsOpenReq {
	return r._paramPageQueryAccountFlowsOpenReq
}
