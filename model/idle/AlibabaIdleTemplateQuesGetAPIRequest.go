package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleTemplateQuesGetAPIRequest 获取SPU最新版本问卷 API请求
// alibaba.idle.template.ques.get
//
// 获取SPU最新版本问卷
type AlibabaIdleTemplateQuesGetAPIRequest struct {
	model.Params
	// 请求参数
	_spuQuestionnaireTopQry *SpuQuestionnaireTopQry
}

// NewAlibabaIdleTemplateQuesGetRequest 初始化AlibabaIdleTemplateQuesGetAPIRequest对象
func NewAlibabaIdleTemplateQuesGetRequest() *AlibabaIdleTemplateQuesGetAPIRequest {
	return &AlibabaIdleTemplateQuesGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleTemplateQuesGetAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.template.ques.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleTemplateQuesGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSpuQuestionnaireTopQry is SpuQuestionnaireTopQry Setter
// 请求参数
func (r *AlibabaIdleTemplateQuesGetAPIRequest) SetSpuQuestionnaireTopQry(_spuQuestionnaireTopQry *SpuQuestionnaireTopQry) error {
	r._spuQuestionnaireTopQry = _spuQuestionnaireTopQry
	r.Set("spu_questionnaire_top_qry", _spuQuestionnaireTopQry)
	return nil
}

// GetSpuQuestionnaireTopQry SpuQuestionnaireTopQry Getter
func (r AlibabaIdleTemplateQuesGetAPIRequest) GetSpuQuestionnaireTopQry() *SpuQuestionnaireTopQry {
	return r._spuQuestionnaireTopQry
}
