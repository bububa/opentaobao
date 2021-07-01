package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminManageTopicAddAPIRequest
新增专题 API请求
yunos.tvpubadmin.manage.topic.add

专题新增 */
type YunosTvpubadminManageTopicAddAPIRequest struct {
	model.Params
	// 新增topic的json信息
	_topicJson string
}

// New
