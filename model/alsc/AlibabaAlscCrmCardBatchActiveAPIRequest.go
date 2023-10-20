package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardBatchActiveAPIRequest 批量激活卡 API请求
// alibaba.alsc.crm.card.batch.active
//
// 批量激活卡
type AlibabaAlscCrmCardBatchActiveAPIRequest struct {
	model.Params
	// 请求对象
	_paramBatchActiveCardOpenReq *BatchActiveCardOpenReq
}

// NewAlibabaAlscCrmCardBatchActiveRequest 初始化AlibabaAlscCrmCardBatchActiveAPIRequest对象
func NewAlibabaAlscCrmCardBatchActiveRequest() *AlibabaAlscCrmCardBatchActiveAPIRequest {
	return &AlibabaAlscCrmCardBatchActiveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmCardBatchActiveAPIRequest) Reset() {
	r._paramBatchActiveCardOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardBatchActiveAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.batch.active"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardBatchActiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmCardBatchActiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBatchActiveCardOpenReq is ParamBatchActiveCardOpenReq Setter
// 请求对象
func (r *AlibabaAlscCrmCardBatchActiveAPIRequest) SetParamBatchActiveCardOpenReq(_paramBatchActiveCardOpenReq *BatchActiveCardOpenReq) error {
	r._paramBatchActiveCardOpenReq = _paramBatchActiveCardOpenReq
	r.Set("param_batch_active_card_open_req", _paramBatchActiveCardOpenReq)
	return nil
}

// GetParamBatchActiveCardOpenReq ParamBatchActiveCardOpenReq Getter
func (r AlibabaAlscCrmCardBatchActiveAPIRequest) GetParamBatchActiveCardOpenReq() *BatchActiveCardOpenReq {
	return r._paramBatchActiveCardOpenReq
}

var poolAlibabaAlscCrmCardBatchActiveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmCardBatchActiveRequest()
	},
}

// GetAlibabaAlscCrmCardBatchActiveRequest 从 sync.Pool 获取 AlibabaAlscCrmCardBatchActiveAPIRequest
func GetAlibabaAlscCrmCardBatchActiveAPIRequest() *AlibabaAlscCrmCardBatchActiveAPIRequest {
	return poolAlibabaAlscCrmCardBatchActiveAPIRequest.Get().(*AlibabaAlscCrmCardBatchActiveAPIRequest)
}

// ReleaseAlibabaAlscCrmCardBatchActiveAPIRequest 将 AlibabaAlscCrmCardBatchActiveAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmCardBatchActiveAPIRequest(v *AlibabaAlscCrmCardBatchActiveAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmCardBatchActiveAPIRequest.Put(v)
}
