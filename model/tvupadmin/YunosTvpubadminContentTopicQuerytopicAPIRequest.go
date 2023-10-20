package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontenttopicquerytopicAPIRequest 迎客松专题查询 API请求
// yunos.tvpubadmin.content.topic.querytopic
//
// 迎客松专题查询
type YunostvpubadmincontenttopicquerytopicAPIRequest struct {
	model.Params
	// TopicAuditQueryBO
	_topicAuditQuery string
}

// NewYunostvpubadmincontenttopicquerytopicRequest 初始化YunostvpubadmincontenttopicquerytopicAPIRequest对象
func NewYunostvpubadmincontenttopicquerytopicRequest() *YunostvpubadmincontenttopicquerytopicAPIRequest {
	return &YunostvpubadmincontenttopicquerytopicAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontenttopicquerytopicAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.topic.querytopic"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontenttopicquerytopicAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontenttopicquerytopicAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopicAuditQuery is TopicAuditQuery Setter
// TopicAuditQueryBO
func (r *YunostvpubadmincontenttopicquerytopicAPIRequest) SetTopicAuditQuery(_topicAuditQuery string) error {
	r._topicAuditQuery = _topicAuditQuery
	r.Set("topic_audit_query", _topicAuditQuery)
	return nil
}

// GetTopicAuditQuery TopicAuditQuery Getter
func (r YunostvpubadmincontenttopicquerytopicAPIRequest) GetTopicAuditQuery() string {
	return r._topicAuditQuery
}
