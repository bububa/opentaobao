package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmrechargeaccountgetAPIRequest 查询储值账户信息 API请求
// alibaba.alsc.crm.recharge.account.get
//
// 查询储值账户信息接口
type AlibabaalsccrmrechargeaccountgetAPIRequest struct {
	model.Params
	// 入参
	_paramQueryRechargeAccountOpenReq *QueryRechargeAccountOpenReq
}

// NewAlibabaalsccrmrechargeaccountgetRequest 初始化AlibabaalsccrmrechargeaccountgetAPIRequest对象
func NewAlibabaalsccrmrechargeaccountgetRequest() *AlibabaalsccrmrechargeaccountgetAPIRequest {
	return &AlibabaalsccrmrechargeaccountgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmrechargeaccountgetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.account.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmrechargeaccountgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmrechargeaccountgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQueryRechargeAccountOpenReq is ParamQueryRechargeAccountOpenReq Setter
// 入参
func (r *AlibabaalsccrmrechargeaccountgetAPIRequest) SetParamQueryRechargeAccountOpenReq(_paramQueryRechargeAccountOpenReq *QueryRechargeAccountOpenReq) error {
	r._paramQueryRechargeAccountOpenReq = _paramQueryRechargeAccountOpenReq
	r.Set("param_query_recharge_account_open_req", _paramQueryRechargeAccountOpenReq)
	return nil
}

// GetParamQueryRechargeAccountOpenReq ParamQueryRechargeAccountOpenReq Getter
func (r AlibabaalsccrmrechargeaccountgetAPIRequest) GetParamQueryRechargeAccountOpenReq() *QueryRechargeAccountOpenReq {
	return r._paramQueryRechargeAccountOpenReq
}
