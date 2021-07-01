package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminManageTopicContentaddAPIRequest
专题新增内容 API请求
yunos.tvpubadmin.manage.topic.contentadd

专题新增内容 */
type YunosTvpubadminManageTopicContentaddAPIRequest struct {
	model.Params
	// 新增的专题内容
	_contentJson string
}

// New
