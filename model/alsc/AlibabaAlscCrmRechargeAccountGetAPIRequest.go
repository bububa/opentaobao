package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmRechargeAccountGetAPIRequest
查询储值账户信息 API请求
alibaba.alsc.crm.recharge.account.get

查询储值账户信息接口 */
type AlibabaAlscCrmRechargeAccountGetAPIRequest struct {
	model.Params
	// 入参
	_paramQueryRechargeAccountOpenReq *QueryRechargeAccountOpenReq
}

// NewAlibabaAlscCrmRechargeAccountGetRequest 初始化AlibabaAlscCrmRechargeAccountGetAPIRequest对象
func NewAlibabaAlscCrmRechargeAccountGetRequest() *AlibabaAlscCrmRechargeAccountGetAPIRequest {
	return &AlibabaAlscCrmRechargeAccountGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeAccountGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.account.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeAccountGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamQueryRechargeAccountOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeAccountGetAPIRequest) SetParamQueryRechargeAccountOpenReq(_paramQueryRechargeAccountOpenReq *QueryRechargeAccountOpenReq) error {
	r._paramQueryRechargeAccountOpenReq = _paramQueryRechargeAccountOpenReq
	r.Set("param_query_recharge_account_open_req", _paramQueryRechargeAccountOpenReq)
	return nil
}

// Get ParamQueryRechargeAccountOpenReq Getter
func (r AlibabaAlscCrmRechargeAccountGetAPIRequest) GetParamQueryRechargeAccountOpenReq() *QueryRechargeAccountOpenReq {
	return r._paramQueryRechargeAccountOpenReq
}
