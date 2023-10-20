package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmOpenRechargeOperateAPIRequest 储值操作接口 API请求
// alibaba.alsc.crm.open.recharge.operate
//
// 储值操作接口
type AlibabaAlscCrmOpenRechargeOperateAPIRequest struct {
	model.Params
	// 储值操作参数
	_paramRechargeOperateOpenReq *RechargeOperateOpenReq
}

// NewAlibabaAlscCrmOpenRechargeOperateRequest 初始化AlibabaAlscCrmOpenRechargeOperateAPIRequest对象
func NewAlibabaAlscCrmOpenRechargeOperateRequest() *AlibabaAlscCrmOpenRechargeOperateAPIRequest {
	return &AlibabaAlscCrmOpenRechargeOperateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmOpenRechargeOperateAPIRequest) Reset() {
	r._paramRechargeOperateOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmOpenRechargeOperateAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.open.recharge.operate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmOpenRechargeOperateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmOpenRechargeOperateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamRechargeOperateOpenReq is ParamRechargeOperateOpenReq Setter
// 储值操作参数
func (r *AlibabaAlscCrmOpenRechargeOperateAPIRequest) SetParamRechargeOperateOpenReq(_paramRechargeOperateOpenReq *RechargeOperateOpenReq) error {
	r._paramRechargeOperateOpenReq = _paramRechargeOperateOpenReq
	r.Set("param_recharge_operate_open_req", _paramRechargeOperateOpenReq)
	return nil
}

// GetParamRechargeOperateOpenReq ParamRechargeOperateOpenReq Getter
func (r AlibabaAlscCrmOpenRechargeOperateAPIRequest) GetParamRechargeOperateOpenReq() *RechargeOperateOpenReq {
	return r._paramRechargeOperateOpenReq
}

var poolAlibabaAlscCrmOpenRechargeOperateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmOpenRechargeOperateRequest()
	},
}

// GetAlibabaAlscCrmOpenRechargeOperateRequest 从 sync.Pool 获取 AlibabaAlscCrmOpenRechargeOperateAPIRequest
func GetAlibabaAlscCrmOpenRechargeOperateAPIRequest() *AlibabaAlscCrmOpenRechargeOperateAPIRequest {
	return poolAlibabaAlscCrmOpenRechargeOperateAPIRequest.Get().(*AlibabaAlscCrmOpenRechargeOperateAPIRequest)
}

// ReleaseAlibabaAlscCrmOpenRechargeOperateAPIRequest 将 AlibabaAlscCrmOpenRechargeOperateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmOpenRechargeOperateAPIRequest(v *AlibabaAlscCrmOpenRechargeOperateAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmOpenRechargeOperateAPIRequest.Put(v)
}
