package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminManageTopicAddAPIRequest 新增专题 API请求
// yunos.tvpubadmin.manage.topic.add
//
// 专题新增
type YunosTvpubadminManageTopicAddAPIRequest struct {
	model.Params
	// 新增topic的json信息
	_topicJson string
}

// NewYunosTvpubadminManageTopicAddRequest 初始化YunosTvpubadminManageTopicAddAPIRequest对象
func NewYunosTvpubadminManageTopicAddRequest() *YunosTvpubadminManageTopicAddAPIRequest {
	return &YunosTvpubadminManageTopicAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminManageTopicAddAPIRequest) Reset() {
	r._topicJson = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageTopicAddAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.manage.topic.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageTopicAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminManageTopicAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopicJson is TopicJson Setter
// 新增topic的json信息
func (r *YunosTvpubadminManageTopicAddAPIRequest) SetTopicJson(_topicJson string) error {
	r._topicJson = _topicJson
	r.Set("topic_json", _topicJson)
	return nil
}

// GetTopicJson TopicJson Getter
func (r YunosTvpubadminManageTopicAddAPIRequest) GetTopicJson() string {
	return r._topicJson
}

var poolYunosTvpubadminManageTopicAddAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminManageTopicAddRequest()
	},
}

// GetYunosTvpubadminManageTopicAddRequest 从 sync.Pool 获取 YunosTvpubadminManageTopicAddAPIRequest
func GetYunosTvpubadminManageTopicAddAPIRequest() *YunosTvpubadminManageTopicAddAPIRequest {
	return poolYunosTvpubadminManageTopicAddAPIRequest.Get().(*YunosTvpubadminManageTopicAddAPIRequest)
}

// ReleaseYunosTvpubadminManageTopicAddAPIRequest 将 YunosTvpubadminManageTopicAddAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminManageTopicAddAPIRequest(v *YunosTvpubadminManageTopicAddAPIRequest) {
	v.Reset()
	poolYunosTvpubadminManageTopicAddAPIRequest.Put(v)
}
