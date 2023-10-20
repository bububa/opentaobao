package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmcardactiveAPIRequest 标准激活卡 API请求
// alibaba.alsc.crm.card.active
//
// 激活卡
type AlibabaalsccrmcardactiveAPIRequest struct {
	model.Params
	// 请求参数
	_paramActiveCardOpenReq *ActiveCardOpenReq
}

// NewAlibabaalsccrmcardactiveRequest 初始化AlibabaalsccrmcardactiveAPIRequest对象
func NewAlibabaalsccrmcardactiveRequest() *AlibabaalsccrmcardactiveAPIRequest {
	return &AlibabaalsccrmcardactiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmcardactiveAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.active"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmcardactiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmcardactiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamActiveCardOpenReq is ParamActiveCardOpenReq Setter
// 请求参数
func (r *AlibabaalsccrmcardactiveAPIRequest) SetParamActiveCardOpenReq(_paramActiveCardOpenReq *ActiveCardOpenReq) error {
	r._paramActiveCardOpenReq = _paramActiveCardOpenReq
	r.Set("param_active_card_open_req", _paramActiveCardOpenReq)
	return nil
}

// GetParamActiveCardOpenReq ParamActiveCardOpenReq Getter
func (r AlibabaalsccrmcardactiveAPIRequest) GetParamActiveCardOpenReq() *ActiveCardOpenReq {
	return r._paramActiveCardOpenReq
}
