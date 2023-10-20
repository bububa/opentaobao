package alsc

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest) Reset() {
	r._paramQueryRechargeAccountFlowOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.account.flowdetail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQueryRechargeAccountFlowOpenReq is ParamQueryRechargeAccountFlowOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest) SetParamQueryRechargeAccountFlowOpenReq(_paramQueryRechargeAccountFlowOpenReq *QueryRechargeAccountFlowOpenReq) error {
	r._paramQueryRechargeAccountFlowOpenReq = _paramQueryRechargeAccountFlowOpenReq
	r.Set("param_query_recharge_account_flow_open_req", _paramQueryRechargeAccountFlowOpenReq)
	return nil
}

// GetParamQueryRechargeAccountFlowOpenReq ParamQueryRechargeAccountFlowOpenReq Getter
func (r AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest) GetParamQueryRechargeAccountFlowOpenReq() *QueryRechargeAccountFlowOpenReq {
	return r._paramQueryRechargeAccountFlowOpenReq
}

var poolAlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmRechargeAccountFlowdetailGetRequest()
	},
}

// GetAlibabaAlscCrmRechargeAccountFlowdetailGetRequest 从 sync.Pool 获取 AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest
func GetAlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest() *AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest {
	return poolAlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest.Get().(*AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest)
}

// ReleaseAlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest 将 AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest(v *AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest.Put(v)
}
