package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardBatchSellAPIRequest 批量开卡（售卡） API请求
// alibaba.alsc.crm.card.batch.sell
//
// 批量开卡（售卡）
type AlibabaAlscCrmCardBatchSellAPIRequest struct {
	model.Params
	// 请求对象
	_paramBatchOpenCardOpenReq *BatchOpenCardOpenReq
}

// NewAlibabaAlscCrmCardBatchSellRequest 初始化AlibabaAlscCrmCardBatchSellAPIRequest对象
func NewAlibabaAlscCrmCardBatchSellRequest() *AlibabaAlscCrmCardBatchSellAPIRequest {
	return &AlibabaAlscCrmCardBatchSellAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmCardBatchSellAPIRequest) Reset() {
	r._paramBatchOpenCardOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardBatchSellAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.batch.sell"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardBatchSellAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmCardBatchSellAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBatchOpenCardOpenReq is ParamBatchOpenCardOpenReq Setter
// 请求对象
func (r *AlibabaAlscCrmCardBatchSellAPIRequest) SetParamBatchOpenCardOpenReq(_paramBatchOpenCardOpenReq *BatchOpenCardOpenReq) error {
	r._paramBatchOpenCardOpenReq = _paramBatchOpenCardOpenReq
	r.Set("param_batch_open_card_open_req", _paramBatchOpenCardOpenReq)
	return nil
}

// GetParamBatchOpenCardOpenReq ParamBatchOpenCardOpenReq Getter
func (r AlibabaAlscCrmCardBatchSellAPIRequest) GetParamBatchOpenCardOpenReq() *BatchOpenCardOpenReq {
	return r._paramBatchOpenCardOpenReq
}

var poolAlibabaAlscCrmCardBatchSellAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmCardBatchSellRequest()
	},
}

// GetAlibabaAlscCrmCardBatchSellRequest 从 sync.Pool 获取 AlibabaAlscCrmCardBatchSellAPIRequest
func GetAlibabaAlscCrmCardBatchSellAPIRequest() *AlibabaAlscCrmCardBatchSellAPIRequest {
	return poolAlibabaAlscCrmCardBatchSellAPIRequest.Get().(*AlibabaAlscCrmCardBatchSellAPIRequest)
}

// ReleaseAlibabaAlscCrmCardBatchSellAPIRequest 将 AlibabaAlscCrmCardBatchSellAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmCardBatchSellAPIRequest(v *AlibabaAlscCrmCardBatchSellAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmCardBatchSellAPIRequest.Put(v)
}
