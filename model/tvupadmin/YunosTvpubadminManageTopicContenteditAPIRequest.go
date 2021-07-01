package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminManageTopicContenteditAPIRequest
专题关联内容编辑 API请求
yunos.tvpubadmin.manage.topic.contentedit

编辑专题关联的内容 */
type YunosTvpubadminManageTopicContenteditAPIRequest struct {
	model.Params
	// 入参，json格式
	_contentJson string
}

// New
