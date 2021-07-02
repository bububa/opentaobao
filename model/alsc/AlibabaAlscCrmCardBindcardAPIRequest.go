package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardBindcardAPIRequest 绑定物理卡 API请求
// alibaba.alsc.crm.card.bindcard
//
// 将会员卡和实例物理卡绑定在一起
type AlibabaAlscCrmCardBindcardAPIRequest struct {
	model.Params
	// 请求参数
	_paramBindPhysicalCardOpenReq *BindPhysicalCardOpenReq
}

// NewAlibabaAlscCrmCardBindcardRequest 初始化AlibabaAlscCrmCardBindcardAPIRequest对象
func NewAlibabaAlscCrmCardBindcardRequest() *AlibabaAlscCrmCardBindcardAPIRequest {
	return &AlibabaAlscCrmCardBindcardAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardBindcardAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.bindcard"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardBindcardAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamBindPhysicalCardOpenReq Setter
// 请求参数
func (r *AlibabaAlscCrmCardBindcardAPIRequest) SetParamBindPhysicalCardOpenReq(_paramBindPhysicalCardOpenReq *BindPhysicalCardOpenReq) error {
	r._paramBindPhysicalCardOpenReq = _paramBindPhysicalCardOpenReq
	r.Set("param_bind_physical_card_open_req", _paramBindPhysicalCardOpenReq)
	return nil
}

// Get ParamBindPhysicalCardOpenReq Getter
func (r AlibabaAlscCrmCardBindcardAPIRequest) GetParamBindPhysicalCardOpenReq() *BindPhysicalCardOpenReq {
	return r._paramBindPhysicalCardOpenReq
}
