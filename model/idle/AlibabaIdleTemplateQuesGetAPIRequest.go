package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidletemplatequesgetAPIRequest 获取SPU最新版本问卷 API请求
// alibaba.idle.template.ques.get
//
// 获取SPU最新版本问卷
type AlibabaidletemplatequesgetAPIRequest struct {
	model.Params
	// 请求参数
	_spuQuestionnaireTopQry *SpuQuestionnaireTopQry
}

// NewAlibabaidletemplatequesgetRequest 初始化AlibabaidletemplatequesgetAPIRequest对象
func NewAlibabaidletemplatequesgetRequest() *AlibabaidletemplatequesgetAPIRequest {
	return &AlibabaidletemplatequesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidletemplatequesgetAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.template.ques.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidletemplatequesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidletemplatequesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSpuQuestionnaireTopQry is SpuQuestionnaireTopQry Setter
// 请求参数
func (r *AlibabaidletemplatequesgetAPIRequest) SetSpuQuestionnaireTopQry(_spuQuestionnaireTopQry *SpuQuestionnaireTopQry) error {
	r._spuQuestionnaireTopQry = _spuQuestionnaireTopQry
	r.Set("spu_questionnaire_top_qry", _spuQuestionnaireTopQry)
	return nil
}

// GetSpuQuestionnaireTopQry SpuQuestionnaireTopQry Getter
func (r AlibabaidletemplatequesgetAPIRequest) GetSpuQuestionnaireTopQry() *SpuQuestionnaireTopQry {
	return r._spuQuestionnaireTopQry
}
