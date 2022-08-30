package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleTemplateQuesOnlineAPIRequest 预上线SPU问卷版本 API请求
// alibaba.idle.template.ques.online
//
// 获取SPU最新版本问卷
type AlibabaIdleTemplateQuesOnlineAPIRequest struct {
	model.Params
	// 参数
	_questionnaireSupportCmd *QuestionnaireSupportCmd
}

// NewAlibabaIdleTemplateQuesOnlineRequest 初始化AlibabaIdleTemplateQuesOnlineAPIRequest对象
func NewAlibabaIdleTemplateQuesOnlineRequest() *AlibabaIdleTemplateQuesOnlineAPIRequest {
	return &AlibabaIdleTemplateQuesOnlineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleTemplateQuesOnlineAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.template.ques.online"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleTemplateQuesOnlineAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetQuestionnaireSupportCmd is QuestionnaireSupportCmd Setter
// 参数
func (r *AlibabaIdleTemplateQuesOnlineAPIRequest) SetQuestionnaireSupportCmd(_questionnaireSupportCmd *QuestionnaireSupportCmd) error {
	r._questionnaireSupportCmd = _questionnaireSupportCmd
	r.Set("questionnaire_support_cmd", _questionnaireSupportCmd)
	return nil
}

// GetQuestionnaireSupportCmd QuestionnaireSupportCmd Getter
func (r AlibabaIdleTemplateQuesOnlineAPIRequest) GetQuestionnaireSupportCmd() *QuestionnaireSupportCmd {
	return r._questionnaireSupportCmd
}
