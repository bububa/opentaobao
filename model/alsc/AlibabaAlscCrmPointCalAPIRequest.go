package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmPointCalAPIRequest 计算积分可以抵扣的金额 API请求
// alibaba.alsc.crm.point.cal
//
// 计算积分可以抵扣的金额
// 积分的抵扣为区间型
// 如抵扣规则为100积分抵扣50元，则输入消费120积分的话，回返回消费100积分抵扣50元
//
//	这里为纯计算逻辑，不会校验用户是否有足够的可用积分进行抵扣
type AlibabaAlscCrmPointCalAPIRequest struct {
	model.Params
	// 入参
	_paramCalculateDeductedMoneyOpenReq *CalculateDeductedMoneyOpenReq
}

// NewAlibabaAlscCrmPointCalRequest 初始化AlibabaAlscCrmPointCalAPIRequest对象
func NewAlibabaAlscCrmPointCalRequest() *AlibabaAlscCrmPointCalAPIRequest {
	return &AlibabaAlscCrmPointCalAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmPointCalAPIRequest) Reset() {
	r._paramCalculateDeductedMoneyOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmPointCalAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.point.cal"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmPointCalAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmPointCalAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCalculateDeductedMoneyOpenReq is ParamCalculateDeductedMoneyOpenReq Setter
// 入参
func (r *AlibabaAlscCrmPointCalAPIRequest) SetParamCalculateDeductedMoneyOpenReq(_paramCalculateDeductedMoneyOpenReq *CalculateDeductedMoneyOpenReq) error {
	r._paramCalculateDeductedMoneyOpenReq = _paramCalculateDeductedMoneyOpenReq
	r.Set("param_calculate_deducted_money_open_req", _paramCalculateDeductedMoneyOpenReq)
	return nil
}

// GetParamCalculateDeductedMoneyOpenReq ParamCalculateDeductedMoneyOpenReq Getter
func (r AlibabaAlscCrmPointCalAPIRequest) GetParamCalculateDeductedMoneyOpenReq() *CalculateDeductedMoneyOpenReq {
	return r._paramCalculateDeductedMoneyOpenReq
}

var poolAlibabaAlscCrmPointCalAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmPointCalRequest()
	},
}

// GetAlibabaAlscCrmPointCalRequest 从 sync.Pool 获取 AlibabaAlscCrmPointCalAPIRequest
func GetAlibabaAlscCrmPointCalAPIRequest() *AlibabaAlscCrmPointCalAPIRequest {
	return poolAlibabaAlscCrmPointCalAPIRequest.Get().(*AlibabaAlscCrmPointCalAPIRequest)
}

// ReleaseAlibabaAlscCrmPointCalAPIRequest 将 AlibabaAlscCrmPointCalAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmPointCalAPIRequest(v *AlibabaAlscCrmPointCalAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmPointCalAPIRequest.Put(v)
}
