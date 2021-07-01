package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminManageTopicContentdeleteAPIRequest
删除专题下内容 API请求
yunos.tvpubadmin.manage.topic.contentdelete

删除专题下内容信息 */
type YunosTvpubadminManageTopicContentdeleteAPIRequest struct {
	model.Params
	// 节目id
	_id int64
}

// New
