package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmrechargeaccountflowdetailgetAPIRequest 储值流水详细 API请求
// alibaba.alsc.crm.recharge.account.flowdetail.get
//
// 查询储值流水详细接口
type AlibabaalsccrmrechargeaccountflowdetailgetAPIRequest struct {
	model.Params
	// 入参
	_paramQueryRechargeAccountFlowOpenReq *QueryRechargeAccountFlowOpenReq
}

// NewAlibabaalsccrmrechargeaccountflowdetailgetRequest 初始化AlibabaalsccrmrechargeaccountflowdetailgetAPIRequest对象
func NewAlibabaalsccrmrechargeaccountflowdetailgetRequest() *AlibabaalsccrmrechargeaccountflowdetailgetAPIRequest {
	return &AlibabaalsccrmrechargeaccountflowdetailgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmrechargeaccountflowdetailgetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.account.flowdetail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmrechargeaccountflowdetailgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmrechargeaccountflowdetailgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQueryRechargeAccountFlowOpenReq is ParamQueryRechargeAccountFlowOpenReq Setter
// 入参
func (r *AlibabaalsccrmrechargeaccountflowdetailgetAPIRequest) SetParamQueryRechargeAccountFlowOpenReq(_paramQueryRechargeAccountFlowOpenReq *QueryRechargeAccountFlowOpenReq) error {
	r._paramQueryRechargeAccountFlowOpenReq = _paramQueryRechargeAccountFlowOpenReq
	r.Set("param_query_recharge_account_flow_open_req", _paramQueryRechargeAccountFlowOpenReq)
	return nil
}

// GetParamQueryRechargeAccountFlowOpenReq ParamQueryRechargeAccountFlowOpenReq Getter
func (r AlibabaalsccrmrechargeaccountflowdetailgetAPIRequest) GetParamQueryRechargeAccountFlowOpenReq() *QueryRechargeAccountFlowOpenReq {
	return r._paramQueryRechargeAccountFlowOpenReq
}
