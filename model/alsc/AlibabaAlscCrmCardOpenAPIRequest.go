package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardOpenAPIRequest 标准开卡流程 API请求
// alibaba.alsc.crm.card.open
//
// 标准开卡流程
type AlibabaAlscCrmCardOpenAPIRequest struct {
	model.Params
	// 开卡参数
	_paramOpenCardStandardOpenReq *OpenCardStandardOpenReq
}

// NewAlibabaAlscCrmCardOpenRequest 初始化AlibabaAlscCrmCardOpenAPIRequest对象
func NewAlibabaAlscCrmCardOpenRequest() *AlibabaAlscCrmCardOpenAPIRequest {
	return &AlibabaAlscCrmCardOpenAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardOpenAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.open"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardOpenAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamOpenCardStandardOpenReq Setter
// 开卡参数
func (r *AlibabaAlscCrmCardOpenAPIRequest) SetParamOpenCardStandardOpenReq(_paramOpenCardStandardOpenReq *OpenCardStandardOpenReq) error {
	r._paramOpenCardStandardOpenReq = _paramOpenCardStandardOpenReq
	r.Set("param_open_card_standard_open_req", _paramOpenCardStandardOpenReq)
	return nil
}

// Get ParamOpenCardStandardOpenReq Getter
func (r AlibabaAlscCrmCardOpenAPIRequest) GetParamOpenCardStandardOpenReq() *OpenCardStandardOpenReq {
	return r._paramOpenCardStandardOpenReq
}
