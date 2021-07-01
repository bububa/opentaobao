package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminManageTopicFindbyidAPIRequest
根据id获取专题信息 API请求
yunos.tvpubadmin.manage.topic.findbyid

根据id获取专题信息 */
type YunosTvpubadminManageTopicFindbyidAPIRequest struct {
	model.Params
	// 专题id
	_id int64
}

// New
