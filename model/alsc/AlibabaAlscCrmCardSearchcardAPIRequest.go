package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardSearchcardAPIRequest 搜索卡实例列表(支持号段查询) API请求
// alibaba.alsc.crm.card.searchcard
//
// 搜索卡实例列表(支持号段查询)
type AlibabaAlscCrmCardSearchcardAPIRequest struct {
	model.Params
	// 请求对象
	_paramSearchCardOpenReq *SearchCardOpenReq
}

// NewAlibabaAlscCrmCardSearchcardRequest 初始化AlibabaAlscCrmCardSearchcardAPIRequest对象
func NewAlibabaAlscCrmCardSearchcardRequest() *AlibabaAlscCrmCardSearchcardAPIRequest {
	return &AlibabaAlscCrmCardSearchcardAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardSearchcardAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.searchcard"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardSearchcardAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamSearchCardOpenReq is ParamSearchCardOpenReq Setter
// 请求对象
func (r *AlibabaAlscCrmCardSearchcardAPIRequest) SetParamSearchCardOpenReq(_paramSearchCardOpenReq *SearchCardOpenReq) error {
	r._paramSearchCardOpenReq = _paramSearchCardOpenReq
	r.Set("param_search_card_open_req", _paramSearchCardOpenReq)
	return nil
}

// GetParamSearchCardOpenReq ParamSearchCardOpenReq Getter
func (r AlibabaAlscCrmCardSearchcardAPIRequest) GetParamSearchCardOpenReq() *SearchCardOpenReq {
	return r._paramSearchCardOpenReq
}
