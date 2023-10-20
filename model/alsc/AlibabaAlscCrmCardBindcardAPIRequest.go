package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmcardbindcardAPIRequest 绑定物理卡 API请求
// alibaba.alsc.crm.card.bindcard
//
// 将会员卡和实例物理卡绑定在一起
type AlibabaalsccrmcardbindcardAPIRequest struct {
	model.Params
	// 请求参数
	_paramBindPhysicalCardOpenReq *BindPhysicalCardOpenReq
}

// NewAlibabaalsccrmcardbindcardRequest 初始化AlibabaalsccrmcardbindcardAPIRequest对象
func NewAlibabaalsccrmcardbindcardRequest() *AlibabaalsccrmcardbindcardAPIRequest {
	return &AlibabaalsccrmcardbindcardAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmcardbindcardAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.bindcard"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmcardbindcardAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmcardbindcardAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBindPhysicalCardOpenReq is ParamBindPhysicalCardOpenReq Setter
// 请求参数
func (r *AlibabaalsccrmcardbindcardAPIRequest) SetParamBindPhysicalCardOpenReq(_paramBindPhysicalCardOpenReq *BindPhysicalCardOpenReq) error {
	r._paramBindPhysicalCardOpenReq = _paramBindPhysicalCardOpenReq
	r.Set("param_bind_physical_card_open_req", _paramBindPhysicalCardOpenReq)
	return nil
}

// GetParamBindPhysicalCardOpenReq ParamBindPhysicalCardOpenReq Getter
func (r AlibabaalsccrmcardbindcardAPIRequest) GetParamBindPhysicalCardOpenReq() *BindPhysicalCardOpenReq {
	return r._paramBindPhysicalCardOpenReq
}
