package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRechargeDedutUpdateAPIRequest 储值消费 API请求
// alibaba.alsc.crm.recharge.dedut.update
//
// 增加储值消费接口
type AlibabaAlscCrmRechargeDedutUpdateAPIRequest struct {
	model.Params
	// 入参
	_paramDedutOpenReq *DedutOpenReq
}

// NewAlibabaAlscCrmRechargeDedutUpdateRequest 初始化AlibabaAlscCrmRechargeDedutUpdateAPIRequest对象
func NewAlibabaAlscCrmRechargeDedutUpdateRequest() *AlibabaAlscCrmRechargeDedutUpdateAPIRequest {
	return &AlibabaAlscCrmRechargeDedutUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmRechargeDedutUpdateAPIRequest) Reset() {
	r._paramDedutOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeDedutUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.dedut.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeDedutUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmRechargeDedutUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamDedutOpenReq is ParamDedutOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeDedutUpdateAPIRequest) SetParamDedutOpenReq(_paramDedutOpenReq *DedutOpenReq) error {
	r._paramDedutOpenReq = _paramDedutOpenReq
	r.Set("param_dedut_open_req", _paramDedutOpenReq)
	return nil
}

// GetParamDedutOpenReq ParamDedutOpenReq Getter
func (r AlibabaAlscCrmRechargeDedutUpdateAPIRequest) GetParamDedutOpenReq() *DedutOpenReq {
	return r._paramDedutOpenReq
}

var poolAlibabaAlscCrmRechargeDedutUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmRechargeDedutUpdateRequest()
	},
}

// GetAlibabaAlscCrmRechargeDedutUpdateRequest 从 sync.Pool 获取 AlibabaAlscCrmRechargeDedutUpdateAPIRequest
func GetAlibabaAlscCrmRechargeDedutUpdateAPIRequest() *AlibabaAlscCrmRechargeDedutUpdateAPIRequest {
	return poolAlibabaAlscCrmRechargeDedutUpdateAPIRequest.Get().(*AlibabaAlscCrmRechargeDedutUpdateAPIRequest)
}

// ReleaseAlibabaAlscCrmRechargeDedutUpdateAPIRequest 将 AlibabaAlscCrmRechargeDedutUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmRechargeDedutUpdateAPIRequest(v *AlibabaAlscCrmRechargeDedutUpdateAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmRechargeDedutUpdateAPIRequest.Put(v)
}
