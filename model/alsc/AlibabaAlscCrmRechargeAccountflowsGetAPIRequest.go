package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmrechargeaccountflowsgetAPIRequest 分页查询储值流水 API请求
// alibaba.alsc.crm.recharge.accountflows.get
//
// 增加分页查询储值流水接口
type AlibabaalsccrmrechargeaccountflowsgetAPIRequest struct {
	model.Params
	// 入参
	_paramPageQueryAccountFlowsOpenReq *PageQueryAccountFlowsOpenReq
}

// NewAlibabaalsccrmrechargeaccountflowsgetRequest 初始化AlibabaalsccrmrechargeaccountflowsgetAPIRequest对象
func NewAlibabaalsccrmrechargeaccountflowsgetRequest() *AlibabaalsccrmrechargeaccountflowsgetAPIRequest {
	return &AlibabaalsccrmrechargeaccountflowsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmrechargeaccountflowsgetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.accountflows.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmrechargeaccountflowsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmrechargeaccountflowsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPageQueryAccountFlowsOpenReq is ParamPageQueryAccountFlowsOpenReq Setter
// 入参
func (r *AlibabaalsccrmrechargeaccountflowsgetAPIRequest) SetParamPageQueryAccountFlowsOpenReq(_paramPageQueryAccountFlowsOpenReq *PageQueryAccountFlowsOpenReq) error {
	r._paramPageQueryAccountFlowsOpenReq = _paramPageQueryAccountFlowsOpenReq
	r.Set("param_page_query_account_flows_open_req", _paramPageQueryAccountFlowsOpenReq)
	return nil
}

// GetParamPageQueryAccountFlowsOpenReq ParamPageQueryAccountFlowsOpenReq Getter
func (r AlibabaalsccrmrechargeaccountflowsgetAPIRequest) GetParamPageQueryAccountFlowsOpenReq() *PageQueryAccountFlowsOpenReq {
	return r._paramPageQueryAccountFlowsOpenReq
}
