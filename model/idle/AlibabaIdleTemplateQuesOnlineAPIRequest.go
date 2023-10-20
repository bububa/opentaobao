package idle

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleTemplateQuesOnlineAPIRequest) Reset() {
	r._questionnaireSupportCmd = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleTemplateQuesOnlineAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.template.ques.online"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleTemplateQuesOnlineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleTemplateQuesOnlineAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaIdleTemplateQuesOnlineAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleTemplateQuesOnlineRequest()
	},
}

// GetAlibabaIdleTemplateQuesOnlineRequest 从 sync.Pool 获取 AlibabaIdleTemplateQuesOnlineAPIRequest
func GetAlibabaIdleTemplateQuesOnlineAPIRequest() *AlibabaIdleTemplateQuesOnlineAPIRequest {
	return poolAlibabaIdleTemplateQuesOnlineAPIRequest.Get().(*AlibabaIdleTemplateQuesOnlineAPIRequest)
}

// ReleaseAlibabaIdleTemplateQuesOnlineAPIRequest 将 AlibabaIdleTemplateQuesOnlineAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleTemplateQuesOnlineAPIRequest(v *AlibabaIdleTemplateQuesOnlineAPIRequest) {
	v.Reset()
	poolAlibabaIdleTemplateQuesOnlineAPIRequest.Put(v)
}
