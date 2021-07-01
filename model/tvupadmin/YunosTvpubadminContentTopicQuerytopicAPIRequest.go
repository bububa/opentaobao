package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentTopicQuerytopicAPIRequest
迎客松专题查询 API请求
yunos.tvpubadmin.content.topic.querytopic

迎客松专题查询 */
type YunosTvpubadminContentTopicQuerytopicAPIRequest struct {
	model.Params
	// TopicAuditQueryBO
	_topicAuditQuery string
}

// New
