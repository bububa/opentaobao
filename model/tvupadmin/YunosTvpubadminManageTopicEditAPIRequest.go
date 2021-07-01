package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminManageTopicEditAPIRequest
编辑专题 API请求
yunos.tvpubadmin.manage.topic.edit

编辑专题 */
type YunosTvpubadminManageTopicEditAPIRequest struct {
	model.Params
	// 待编辑专题
	_topicJson string
}

// New
