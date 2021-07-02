package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardQueryTemplateAPIRequest 查询卡模板详情 API请求
// alibaba.alsc.crm.card.query.template
//
// 查询卡模板详情
type AlibabaAlscCrmCardQueryTemplateAPIRequest struct {
	model.Params
	// 请求对象
	_paramQueryCardTemplateOpenReq *QueryCardTemplateOpenReq
}

// NewAlibabaAlscCrmCardQueryTemplateRequest 初始化AlibabaAlscCrmCardQueryTemplateAPIRequest对象
func NewAlibabaAlscCrmCardQueryTemplateRequest() *AlibabaAlscCrmCardQueryTemplateAPIRequest {
	return &AlibabaAlscCrmCardQueryTemplateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardQueryTemplateAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.query.template"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardQueryTemplateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamQueryCardTemplateOpenReq Setter
// 请求对象
func (r *AlibabaAlscCrmCardQueryTemplateAPIRequest) SetParamQueryCardTemplateOpenReq(_paramQueryCardTemplateOpenReq *QueryCardTemplateOpenReq) error {
	r._paramQueryCardTemplateOpenReq = _paramQueryCardTemplateOpenReq
	r.Set("param_query_card_template_open_req", _paramQueryCardTemplateOpenReq)
	return nil
}

// Get ParamQueryCardTemplateOpenReq Getter
func (r AlibabaAlscCrmCardQueryTemplateAPIRequest) GetParamQueryCardTemplateOpenReq() *QueryCardTemplateOpenReq {
	return r._paramQueryCardTemplateOpenReq
}
