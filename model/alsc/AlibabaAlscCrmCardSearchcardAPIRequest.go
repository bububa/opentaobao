package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmcardsearchcardAPIRequest 搜索卡实例列表(支持号段查询) API请求
// alibaba.alsc.crm.card.searchcard
//
// 搜索卡实例列表(支持号段查询)
type AlibabaalsccrmcardsearchcardAPIRequest struct {
	model.Params
	// 请求对象
	_paramSearchCardOpenReq *SearchCardOpenReq
}

// NewAlibabaalsccrmcardsearchcardRequest 初始化AlibabaalsccrmcardsearchcardAPIRequest对象
func NewAlibabaalsccrmcardsearchcardRequest() *AlibabaalsccrmcardsearchcardAPIRequest {
	return &AlibabaalsccrmcardsearchcardAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmcardsearchcardAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.searchcard"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmcardsearchcardAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmcardsearchcardAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamSearchCardOpenReq is ParamSearchCardOpenReq Setter
// 请求对象
func (r *AlibabaalsccrmcardsearchcardAPIRequest) SetParamSearchCardOpenReq(_paramSearchCardOpenReq *SearchCardOpenReq) error {
	r._paramSearchCardOpenReq = _paramSearchCardOpenReq
	r.Set("param_search_card_open_req", _paramSearchCardOpenReq)
	return nil
}

// GetParamSearchCardOpenReq ParamSearchCardOpenReq Getter
func (r AlibabaalsccrmcardsearchcardAPIRequest) GetParamSearchCardOpenReq() *SearchCardOpenReq {
	return r._paramSearchCardOpenReq
}
