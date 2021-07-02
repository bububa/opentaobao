package tvupadmin

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentTopicQuerytopicAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.topic.querytopic"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentTopicQuerytopicAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
