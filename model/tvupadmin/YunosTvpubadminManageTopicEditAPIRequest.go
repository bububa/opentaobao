package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminManageTopicEditAPIRequest 编辑专题 API请求
// yunos.tvpubadmin.manage.topic.edit
//
// 编辑专题
type YunosTvpubadminManageTopicEditAPIRequest struct {
	model.Params
	// 待编辑专题
	_topicJson string
}

// NewYunosTvpubadminManageTopicEditRequest 初始化YunosTvpubadminManageTopicEditAPIRequest对象
func NewYunosTvpubadminManageTopicEditRequest() *YunosTvpubadminManageTopicEditAPIRequest {
	return &YunosTvpubadminManageTopicEditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageTopicEditAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.manage.topic.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageTopicEditAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTopicJson is TopicJson Setter
// 待编辑专题
func (r *YunosTvpubadminManageTopicEditAPIRequest) SetTopicJson(_topicJson string) error {
	r._topicJson = _topicJson
	r.Set("topic_json", _topicJson)
	return nil
}

// GetTopicJson TopicJson Getter
func (r YunosTvpubadminManageTopicEditAPIRequest) GetTopicJson() string {
	return r._topicJson
}
