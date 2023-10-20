package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmRechargeUndedutUpdateAPIRequest 储值消费退款(逆向) API请求
// alibaba.alsc.crm.recharge.undedut.update
//
// 新增储值消费退款接口
type AlibabaAlscCrmRechargeUndedutUpdateAPIRequest struct {
	model.Params
	// 入参
	_paramUndedutOpenReq *UndedutOpenReq
}

// NewAlibabaAlscCrmRechargeUndedutUpdateRequest 初始化AlibabaAlscCrmRechargeUndedutUpdateAPIRequest对象
func NewAlibabaAlscCrmRechargeUndedutUpdateRequest() *AlibabaAlscCrmRechargeUndedutUpdateAPIRequest {
	return &AlibabaAlscCrmRechargeUndedutUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmRechargeUndedutUpdateAPIRequest) Reset() {
	r._paramUndedutOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeUndedutUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.recharge.undedut.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeUndedutUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmRechargeUndedutUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamUndedutOpenReq is ParamUndedutOpenReq Setter
// 入参
func (r *AlibabaAlscCrmRechargeUndedutUpdateAPIRequest) SetParamUndedutOpenReq(_paramUndedutOpenReq *UndedutOpenReq) error {
	r._paramUndedutOpenReq = _paramUndedutOpenReq
	r.Set("param_undedut_open_req", _paramUndedutOpenReq)
	return nil
}

// GetParamUndedutOpenReq ParamUndedutOpenReq Getter
func (r AlibabaAlscCrmRechargeUndedutUpdateAPIRequest) GetParamUndedutOpenReq() *UndedutOpenReq {
	return r._paramUndedutOpenReq
}

var poolAlibabaAlscCrmRechargeUndedutUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmRechargeUndedutUpdateRequest()
	},
}

// GetAlibabaAlscCrmRechargeUndedutUpdateRequest 从 sync.Pool 获取 AlibabaAlscCrmRechargeUndedutUpdateAPIRequest
func GetAlibabaAlscCrmRechargeUndedutUpdateAPIRequest() *AlibabaAlscCrmRechargeUndedutUpdateAPIRequest {
	return poolAlibabaAlscCrmRechargeUndedutUpdateAPIRequest.Get().(*AlibabaAlscCrmRechargeUndedutUpdateAPIRequest)
}

// ReleaseAlibabaAlscCrmRechargeUndedutUpdateAPIRequest 将 AlibabaAlscCrmRechargeUndedutUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmRechargeUndedutUpdateAPIRequest(v *AlibabaAlscCrmRechargeUndedutUpdateAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmRechargeUndedutUpdateAPIRequest.Put(v)
}
