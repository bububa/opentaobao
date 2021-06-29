package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松专题查询 API请求
yunos.tvpubadmin.content.topic.querytopic

迎客松专题查询
*/
type YunosTvpubadminContentTopicQuerytopicRequest struct {
    model.Params
    // TopicAuditQueryBO
    topicAuditQuery   string
}

// 初始化YunosTvpubadminContentTopicQuerytopicRequest对象
func NewYunosTvpubadminContentTopicQuerytopicRequest() *YunosTvpubadminContentTopicQuerytopicRequest{
    return &YunosTvpubadminContentTopicQuerytopicRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentTopicQuerytopicRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.topic.querytopic"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentTopicQuerytopicRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TopicAuditQuery Setter
// TopicAuditQueryBO
func (r *YunosTvpubadminContentTopicQuerytopicRequest) SetTopicAuditQuery(topicAuditQuery string) error {
    r.topicAuditQuery = topicAuditQuery
    r.Set("topic_audit_query", topicAuditQuery)
    return nil
}

// TopicAuditQuery Getter
func (r YunosTvpubadminContentTopicQuerytopicRequest) GetTopicAuditQuery() string {
    return r.topicAuditQuery
}
