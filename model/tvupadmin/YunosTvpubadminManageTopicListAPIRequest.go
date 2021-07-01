package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminManageTopicListAPIRequest
专题内容操作列表 API请求
yunos.tvpubadmin.manage.topic.list

获取外部可操作编辑的专题列表 */
type YunosTvpubadminManageTopicListAPIRequest struct {
	model.Params
	// 查询条件
	_query string
}

// New
