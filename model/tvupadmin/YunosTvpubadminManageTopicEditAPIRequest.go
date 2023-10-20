package tvupadmin

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminManageTopicEditAPIRequest) Reset() {
	r._topicJson = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageTopicEditAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.manage.topic.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageTopicEditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminManageTopicEditAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolYunosTvpubadminManageTopicEditAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminManageTopicEditRequest()
	},
}

// GetYunosTvpubadminManageTopicEditRequest 从 sync.Pool 获取 YunosTvpubadminManageTopicEditAPIRequest
func GetYunosTvpubadminManageTopicEditAPIRequest() *YunosTvpubadminManageTopicEditAPIRequest {
	return poolYunosTvpubadminManageTopicEditAPIRequest.Get().(*YunosTvpubadminManageTopicEditAPIRequest)
}

// ReleaseYunosTvpubadminManageTopicEditAPIRequest 将 YunosTvpubadminManageTopicEditAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminManageTopicEditAPIRequest(v *YunosTvpubadminManageTopicEditAPIRequest) {
	v.Reset()
	poolYunosTvpubadminManageTopicEditAPIRequest.Put(v)
}
