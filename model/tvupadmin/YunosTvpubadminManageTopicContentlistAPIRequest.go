package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminManageTopicContentlistAPIRequest
查看专题内容列表 API请求
yunos.tvpubadmin.manage.topic.contentlist

查看专题内容列表 */
type YunosTvpubadminManageTopicContentlistAPIRequest struct {
	model.Params
	// 节目查询参数
	_programQuery string
}

// New
