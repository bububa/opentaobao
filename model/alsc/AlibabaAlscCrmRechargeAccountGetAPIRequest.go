package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRechargeAccountGetAPIRequest 查询储值账户信息 API请求
// alibaba.alsc.crm.recharge.account.get
//
// 查询储值账户信息接口
type AlibabaAlscCrmRechargeAccountGetAPIRequest struct {
	model.Params
	// 入参
	_paramQueryRechargeAccountOpenReq *QueryRechargeAccountOpenReq
}

// NewAlibabaAlscCrmRechargeAccountGetRequest 初始化AlibabaAlscCrmRechargeAccountGetAPIRequest对象
func NewAlibabaAlscCrmRechargeAccountGetRequest() *AlibabaAlscCrmRechargeAccountGetAPIRequest {
	return &AlibabaAlscCrmRechargeAccountGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmRechargeAccountGetAPIRequest) Reset() {
	r._paramQueryRechargeAccountOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeAccountGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.account.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeAccountGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmRechargeAccountGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQueryRechargeAccountOpenReq is ParamQueryRechargeAccountOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeAccountGetAPIRequest) SetParamQueryRechargeAccountOpenReq(_paramQueryRechargeAccountOpenReq *QueryRechargeAccountOpenReq) error {
	r._paramQueryRechargeAccountOpenReq = _paramQueryRechargeAccountOpenReq
	r.Set("param_query_recharge_account_open_req", _paramQueryRechargeAccountOpenReq)
	return nil
}

// GetParamQueryRechargeAccountOpenReq ParamQueryRechargeAccountOpenReq Getter
func (r AlibabaAlscCrmRechargeAccountGetAPIRequest) GetParamQueryRechargeAccountOpenReq() *QueryRechargeAccountOpenReq {
	return r._paramQueryRechargeAccountOpenReq
}

var poolAlibabaAlscCrmRechargeAccountGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmRechargeAccountGetRequest()
	},
}

// GetAlibabaAlscCrmRechargeAccountGetRequest 从 sync.Pool 获取 AlibabaAlscCrmRechargeAccountGetAPIRequest
func GetAlibabaAlscCrmRechargeAccountGetAPIRequest() *AlibabaAlscCrmRechargeAccountGetAPIRequest {
	return poolAlibabaAlscCrmRechargeAccountGetAPIRequest.Get().(*AlibabaAlscCrmRechargeAccountGetAPIRequest)
}

// ReleaseAlibabaAlscCrmRechargeAccountGetAPIRequest 将 AlibabaAlscCrmRechargeAccountGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmRechargeAccountGetAPIRequest(v *AlibabaAlscCrmRechargeAccountGetAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmRechargeAccountGetAPIRequest.Put(v)
}
