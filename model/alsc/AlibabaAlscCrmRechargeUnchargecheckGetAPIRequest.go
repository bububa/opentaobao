package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest 储值账户退充值校验 API请求
// alibaba.alsc.crm.recharge.unchargecheck.get
//
// 储值账户退充值校验接口
type AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest struct {
	model.Params
	// 入参
	_paramUnchargeCheckOpenReq *UnchargeCheckOpenReq
}

// NewAlibabaAlscCrmRechargeUnchargecheckGetRequest 初始化AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest对象
func NewAlibabaAlscCrmRechargeUnchargecheckGetRequest() *AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest {
	return &AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest) Reset() {
	r._paramUnchargeCheckOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.unchargecheck.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamUnchargeCheckOpenReq is ParamUnchargeCheckOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest) SetParamUnchargeCheckOpenReq(_paramUnchargeCheckOpenReq *UnchargeCheckOpenReq) error {
	r._paramUnchargeCheckOpenReq = _paramUnchargeCheckOpenReq
	r.Set("param_uncharge_check_open_req", _paramUnchargeCheckOpenReq)
	return nil
}

// GetParamUnchargeCheckOpenReq ParamUnchargeCheckOpenReq Getter
func (r AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest) GetParamUnchargeCheckOpenReq() *UnchargeCheckOpenReq {
	return r._paramUnchargeCheckOpenReq
}

var poolAlibabaAlscCrmRechargeUnchargecheckGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmRechargeUnchargecheckGetRequest()
	},
}

// GetAlibabaAlscCrmRechargeUnchargecheckGetRequest 从 sync.Pool 获取 AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest
func GetAlibabaAlscCrmRechargeUnchargecheckGetAPIRequest() *AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest {
	return poolAlibabaAlscCrmRechargeUnchargecheckGetAPIRequest.Get().(*AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest)
}

// ReleaseAlibabaAlscCrmRechargeUnchargecheckGetAPIRequest 将 AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmRechargeUnchargecheckGetAPIRequest(v *AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmRechargeUnchargecheckGetAPIRequest.Put(v)
}
