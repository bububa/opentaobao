package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentTopicQuerytopicAPIRequest 迎客松专题查询 API请求
// yunos.tvpubadmin.content.topic.querytopic
//
// 迎客松专题查询
type YunosTvpubadminContentTopicQuerytopicAPIRequest struct {
	model.Params
	// TopicAuditQueryBO
	_topicAuditQuery string
}

// NewYunosTvpubadminContentTopicQuerytopicRequest 初始化YunosTvpubadminContentTopicQuerytopicAPIRequest对象
func NewYunosTvpubadminContentTopicQuerytopicRequest() *YunosTvpubadminContentTopicQuerytopicAPIRequest {
	return &YunosTvpubadminContentTopicQuerytopicAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentTopicQuerytopicAPIRequest) Reset() {
	r._topicAuditQuery = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentTopicQuerytopicAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.topic.querytopic"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentTopicQuerytopicAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentTopicQuerytopicAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopicAuditQuery is TopicAuditQuery Setter
// TopicAuditQueryBO
func (r *YunosTvpubadminContentTopicQuerytopicAPIRequest) SetTopicAuditQuery(_topicAuditQuery string) error {
	r._topicAuditQuery = _topicAuditQuery
	r.Set("topic_audit_query", _topicAuditQuery)
	return nil
}

// GetTopicAuditQuery TopicAuditQuery Getter
func (r YunosTvpubadminContentTopicQuerytopicAPIRequest) GetTopicAuditQuery() string {
	return r._topicAuditQuery
}

var poolYunosTvpubadminContentTopicQuerytopicAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentTopicQuerytopicRequest()
	},
}

// GetYunosTvpubadminContentTopicQuerytopicRequest 从 sync.Pool 获取 YunosTvpubadminContentTopicQuerytopicAPIRequest
func GetYunosTvpubadminContentTopicQuerytopicAPIRequest() *YunosTvpubadminContentTopicQuerytopicAPIRequest {
	return poolYunosTvpubadminContentTopicQuerytopicAPIRequest.Get().(*YunosTvpubadminContentTopicQuerytopicAPIRequest)
}

// ReleaseYunosTvpubadminContentTopicQuerytopicAPIRequest 将 YunosTvpubadminContentTopicQuerytopicAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentTopicQuerytopicAPIRequest(v *YunosTvpubadminContentTopicQuerytopicAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentTopicQuerytopicAPIRequest.Put(v)
}
