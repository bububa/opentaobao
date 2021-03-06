package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardActiveAPIRequest 标准激活卡 API请求
// alibaba.alsc.crm.card.active
//
// 激活卡
type AlibabaAlscCrmCardActiveAPIRequest struct {
	model.Params
	// 请求参数
	_paramActiveCardOpenReq *ActiveCardOpenReq
}

// NewAlibabaAlscCrmCardActiveRequest 初始化AlibabaAlscCrmCardActiveAPIRequest对象
func NewAlibabaAlscCrmCardActiveRequest() *AlibabaAlscCrmCardActiveAPIRequest {
	return &AlibabaAlscCrmCardActiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardActiveAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.active"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardActiveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamActiveCardOpenReq is ParamActiveCardOpenReq Setter
// 请求参数
func (r *AlibabaAlscCrmCardActiveAPIRequest) SetParamActiveCardOpenReq(_paramActiveCardOpenReq *ActiveCardOpenReq) error {
	r._paramActiveCardOpenReq = _paramActiveCardOpenReq
	r.Set("param_active_card_open_req", _paramActiveCardOpenReq)
	return nil
}

// GetParamActiveCardOpenReq ParamActiveCardOpenReq Getter
func (r AlibabaAlscCrmCardActiveAPIRequest) GetParamActiveCardOpenReq() *ActiveCardOpenReq {
	return r._paramActiveCardOpenReq
}
