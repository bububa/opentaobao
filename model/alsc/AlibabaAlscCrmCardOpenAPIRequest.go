package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmcardopenAPIRequest 标准开卡流程 API请求
// alibaba.alsc.crm.card.open
//
// 标准开卡流程
type AlibabaalsccrmcardopenAPIRequest struct {
	model.Params
	// 开卡参数
	_paramOpenCardStandardOpenReq *OpenCardStandardOpenReq
}

// NewAlibabaalsccrmcardopenRequest 初始化AlibabaalsccrmcardopenAPIRequest对象
func NewAlibabaalsccrmcardopenRequest() *AlibabaalsccrmcardopenAPIRequest {
	return &AlibabaalsccrmcardopenAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmcardopenAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.open"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmcardopenAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmcardopenAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOpenCardStandardOpenReq is ParamOpenCardStandardOpenReq Setter
// 开卡参数
func (r *AlibabaalsccrmcardopenAPIRequest) SetParamOpenCardStandardOpenReq(_paramOpenCardStandardOpenReq *OpenCardStandardOpenReq) error {
	r._paramOpenCardStandardOpenReq = _paramOpenCardStandardOpenReq
	r.Set("param_open_card_standard_open_req", _paramOpenCardStandardOpenReq)
	return nil
}

// GetParamOpenCardStandardOpenReq ParamOpenCardStandardOpenReq Getter
func (r AlibabaalsccrmcardopenAPIRequest) GetParamOpenCardStandardOpenReq() *OpenCardStandardOpenReq {
	return r._paramOpenCardStandardOpenReq
}
