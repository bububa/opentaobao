package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmcardbatchsellAPIRequest 批量开卡（售卡） API请求
// alibaba.alsc.crm.card.batch.sell
//
// 批量开卡（售卡）
type AlibabaalsccrmcardbatchsellAPIRequest struct {
	model.Params
	// 请求对象
	_paramBatchOpenCardOpenReq *BatchOpenCardOpenReq
}

// NewAlibabaalsccrmcardbatchsellRequest 初始化AlibabaalsccrmcardbatchsellAPIRequest对象
func NewAlibabaalsccrmcardbatchsellRequest() *AlibabaalsccrmcardbatchsellAPIRequest {
	return &AlibabaalsccrmcardbatchsellAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmcardbatchsellAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.batch.sell"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmcardbatchsellAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmcardbatchsellAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBatchOpenCardOpenReq is ParamBatchOpenCardOpenReq Setter
// 请求对象
func (r *AlibabaalsccrmcardbatchsellAPIRequest) SetParamBatchOpenCardOpenReq(_paramBatchOpenCardOpenReq *BatchOpenCardOpenReq) error {
	r._paramBatchOpenCardOpenReq = _paramBatchOpenCardOpenReq
	r.Set("param_batch_open_card_open_req", _paramBatchOpenCardOpenReq)
	return nil
}

// GetParamBatchOpenCardOpenReq ParamBatchOpenCardOpenReq Getter
func (r AlibabaalsccrmcardbatchsellAPIRequest) GetParamBatchOpenCardOpenReq() *BatchOpenCardOpenReq {
	return r._paramBatchOpenCardOpenReq
}
