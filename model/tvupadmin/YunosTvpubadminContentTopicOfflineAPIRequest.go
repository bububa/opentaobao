package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentTopicOfflineAPIRequest
迎客松专题下线 API请求
yunos.tvpubadmin.content.topic.offline

迎客松专题下线 */
type YunosTvpubadminContentTopicOfflineAPIRequest struct {
	model.Params
	// id
	_id int64
}

// New
