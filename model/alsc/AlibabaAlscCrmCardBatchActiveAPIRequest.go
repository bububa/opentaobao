package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmcardbatchactiveAPIRequest 批量激活卡 API请求
// alibaba.alsc.crm.card.batch.active
//
// 批量激活卡
type AlibabaalsccrmcardbatchactiveAPIRequest struct {
	model.Params
	// 请求对象
	_paramBatchActiveCardOpenReq *BatchActiveCardOpenReq
}

// NewAlibabaalsccrmcardbatchactiveRequest 初始化AlibabaalsccrmcardbatchactiveAPIRequest对象
func NewAlibabaalsccrmcardbatchactiveRequest() *AlibabaalsccrmcardbatchactiveAPIRequest {
	return &AlibabaalsccrmcardbatchactiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmcardbatchactiveAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.batch.active"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmcardbatchactiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmcardbatchactiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBatchActiveCardOpenReq is ParamBatchActiveCardOpenReq Setter
// 请求对象
func (r *AlibabaalsccrmcardbatchactiveAPIRequest) SetParamBatchActiveCardOpenReq(_paramBatchActiveCardOpenReq *BatchActiveCardOpenReq) error {
	r._paramBatchActiveCardOpenReq = _paramBatchActiveCardOpenReq
	r.Set("param_batch_active_card_open_req", _paramBatchActiveCardOpenReq)
	return nil
}

// GetParamBatchActiveCardOpenReq ParamBatchActiveCardOpenReq Getter
func (r AlibabaalsccrmcardbatchactiveAPIRequest) GetParamBatchActiveCardOpenReq() *BatchActiveCardOpenReq {
	return r._paramBatchActiveCardOpenReq
}
