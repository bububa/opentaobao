package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmcardquerytemplateAPIRequest 查询卡模板详情 API请求
// alibaba.alsc.crm.card.query.template
//
// 查询卡模板详情
type AlibabaalsccrmcardquerytemplateAPIRequest struct {
	model.Params
	// 请求对象
	_paramQueryCardTemplateOpenReq *QueryCardTemplateOpenReq
}

// NewAlibabaalsccrmcardquerytemplateRequest 初始化AlibabaalsccrmcardquerytemplateAPIRequest对象
func NewAlibabaalsccrmcardquerytemplateRequest() *AlibabaalsccrmcardquerytemplateAPIRequest {
	return &AlibabaalsccrmcardquerytemplateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmcardquerytemplateAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.query.template"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmcardquerytemplateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmcardquerytemplateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQueryCardTemplateOpenReq is ParamQueryCardTemplateOpenReq Setter
// 请求对象
func (r *AlibabaalsccrmcardquerytemplateAPIRequest) SetParamQueryCardTemplateOpenReq(_paramQueryCardTemplateOpenReq *QueryCardTemplateOpenReq) error {
	r._paramQueryCardTemplateOpenReq = _paramQueryCardTemplateOpenReq
	r.Set("param_query_card_template_open_req", _paramQueryCardTemplateOpenReq)
	return nil
}

// GetParamQueryCardTemplateOpenReq ParamQueryCardTemplateOpenReq Getter
func (r AlibabaalsccrmcardquerytemplateAPIRequest) GetParamQueryCardTemplateOpenReq() *QueryCardTemplateOpenReq {
	return r._paramQueryCardTemplateOpenReq
}
