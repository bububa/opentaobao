package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidletemplatequesonlineAPIRequest 预上线SPU问卷版本 API请求
// alibaba.idle.template.ques.online
//
// 获取SPU最新版本问卷
type AlibabaidletemplatequesonlineAPIRequest struct {
	model.Params
	// 参数
	_questionnaireSupportCmd *QuestionnaireSupportCmd
}

// NewAlibabaidletemplatequesonlineRequest 初始化AlibabaidletemplatequesonlineAPIRequest对象
func NewAlibabaidletemplatequesonlineRequest() *AlibabaidletemplatequesonlineAPIRequest {
	return &AlibabaidletemplatequesonlineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidletemplatequesonlineAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.template.ques.online"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidletemplatequesonlineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidletemplatequesonlineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuestionnaireSupportCmd is QuestionnaireSupportCmd Setter
// 参数
func (r *AlibabaidletemplatequesonlineAPIRequest) SetQuestionnaireSupportCmd(_questionnaireSupportCmd *QuestionnaireSupportCmd) error {
	r._questionnaireSupportCmd = _questionnaireSupportCmd
	r.Set("questionnaire_support_cmd", _questionnaireSupportCmd)
	return nil
}

// GetQuestionnaireSupportCmd QuestionnaireSupportCmd Getter
func (r AlibabaidletemplatequesonlineAPIRequest) GetQuestionnaireSupportCmd() *QuestionnaireSupportCmd {
	return r._questionnaireSupportCmd
}
