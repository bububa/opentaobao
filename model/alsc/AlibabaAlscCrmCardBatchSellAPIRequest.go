package alsc

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardBatchSellAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.batch.sell"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardBatchSellAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamBatchOpenCardOpenReq Setter
// 请求对象
func (r *AlibabaAlscCrmCardBatchSellAPIRequest) SetParamBatchOpenCardOpenReq(_paramBatchOpenCardOpenReq *BatchOpenCardOpenReq) error {
	r._paramBatchOpenCardOpenReq = _paramBatchOpenCardOpenReq
	r.Set("param_batch_open_card_open_req", _paramBatchOpenCardOpenReq)
	return nil
}

// Get ParamBatchOpenCardOpenReq Getter
func (r AlibabaAlscCrmCardBatchSellAPIRequest) GetParamBatchOpenCardOpenReq() *BatchOpenCardOpenReq {
	return r._paramBatchOpenCardOpenReq
}
