package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松专题查询 APIRequest
yunos.tvpubadmin.content.topic.querytopic

迎客松专题查询
*/
type YunosTvpubadminContentTopicQuerytopicRequest struct {
    model.Params

    // TopicAuditQueryBO
    topicAuditQuery   string 

}

func NewYunosTvpubadminContentTopicQuerytopicRequest() *YunosTvpubadminContentTopicQuerytopicRequest{
    return &YunosTvpubadminContentTopicQuerytopicRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentTopicQuerytopicRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.topic.querytopic"
}

func (r YunosTvpubadminContentTopicQuerytopicRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentTopicQuerytopicRequest) SetTopicAuditQuery(topicAuditQuery string) error {
    r.topicAuditQuery = topicAuditQuery
    r.Set("topic_audit_query", topicAuditQuery)
    return nil
}

func (r YunosTvpubadminContentTopicQuerytopicRequest) GetTopicAuditQuery() string {
    return r.topicAuditQuery
}

