package idle

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleTemplateQuesGetAPIRequest) Reset() {
	r._spuQuestionnaireTopQry = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleTemplateQuesGetAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.template.ques.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleTemplateQuesGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleTemplateQuesGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaIdleTemplateQuesGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleTemplateQuesGetRequest()
	},
}

// GetAlibabaIdleTemplateQuesGetRequest 从 sync.Pool 获取 AlibabaIdleTemplateQuesGetAPIRequest
func GetAlibabaIdleTemplateQuesGetAPIRequest() *AlibabaIdleTemplateQuesGetAPIRequest {
	return poolAlibabaIdleTemplateQuesGetAPIRequest.Get().(*AlibabaIdleTemplateQuesGetAPIRequest)
}

// ReleaseAlibabaIdleTemplateQuesGetAPIRequest 将 AlibabaIdleTemplateQuesGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleTemplateQuesGetAPIRequest(v *AlibabaIdleTemplateQuesGetAPIRequest) {
	v.Reset()
	poolAlibabaIdleTemplateQuesGetAPIRequest.Put(v)
}
